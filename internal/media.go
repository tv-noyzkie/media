package internal

import (
   "41.neocities.org/dash"
   "41.neocities.org/sofia/file"
   "41.neocities.org/sofia/pssh"
   "encoding/base64"
   "fmt"
   "io"
   "net/url"
   "os"
   "slices"
)

func init_protect(data []byte) ([]byte, error) {
   var file1 file.File
   err := file1.Read(data)
   if err != nil {
      return nil, err
   }
   if moov, ok := file1.GetMoov(); ok {
      for _, pssh1 := range moov.Pssh {
         if pssh1.SystemId.String() == widevine_system_id {
            a.pssh = pssh1.Data
         }
         copy(pssh1.BoxHeader.Type[:], "free") // Firefox
      }
      description := moov.Trak.Mdia.Minf.Stbl.Stsd
      if sinf, ok := description.Sinf(); ok {
         a.key_id = sinf.Schi.Tenc.S.DefaultKid[:]
         // Firefox
         copy(sinf.BoxHeader.Type[:], "free")
         if sample, ok := description.SampleEntry(); ok {
            // Firefox
            copy(sample.BoxHeader.Type[:], sinf.Frma.DataFormat[:])
         }
      }
   }
   return file1.Append(nil)
}

// try to get PSSH from DASH then MP4
func dash_pssh(
   client DashClient,
   home string,
   id string,
) error {
   base := &url.URL{}
   var data []byte
   if id != "" {
      var err error
      data, err = os.ReadFile(home + "/mpd_body")
      if err != nil {
         return err
      }
      data1, err := os.ReadFile(home + "/mpd_url")
      if err != nil {
         return err
      }
      err = base.UnmarshalBinary(data1)
      if err != nil {
         return err
      }
   } else {
      resp, err := client.Mpd()
      if err != nil {
         return err
      }
      defer resp.Body.Close()
      data, err = io.ReadAll(resp.Body)
      if err != nil {
         return err
      }
      err = write_file(home + "/mpd_body", data)
      if err != nil {
         return err
      }
      file1, err := create(home + "/mpd_url")
      if err != nil {
         return err
      }
      defer file1.Close()
      base = resp.Request.URL
      fmt.Fprint(file1, base)
   }
   var media dash.Mpd
   err := media.Unmarshal(data)
   if err != nil {
      return err
   }
   media.Set(base)
   represents := slices.SortedFunc(media.Representation(),
      func(a, b dash.Representation) int {
         return a.Bandwidth - b.Bandwidth
      },
   )
   for i, represent := range represents {
      switch id {
      case "":
         if i >= 1 {
            fmt.Println()
         }
         fmt.Println(&represent)
      case represent.Id:
         for _, protect := range represent.ContentProtection {
            if protect.SchemeIdUri != widevine_urn {
               continue
            }
            if protect.Pssh == "" {
               continue
            }
            data, err := base64.StdEncoding.DecodeString(protect.Pssh)
            if err != nil {
               return err
            }
            var box pssh.Box
            n, err := box.BoxHeader.Decode(data)
            if err != nil {
               return err
            }
            err = box.Read(data[n:])
            if err != nil {
               return err
            }
            a.pssh = box.Data
            break
         }
         ext, err := get_ext(&represent)
         if err != nil {
            return err
         }
         if represent.SegmentBase != nil {
            return a.segment_base(&represent, ext)
         }
         if represent.SegmentList != nil {
            return a.segment_list(&represent, ext)
         }
         return a.segment_template(&represent, ext)
      }
   }
   return nil
}
