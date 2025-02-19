package internal

import (
   "41.neocities.org/dash"
   "41.neocities.org/sofia/file"
   "41.neocities.org/sofia/pssh"
   xhttp "41.neocities.org/x/http"
   "encoding/base64"
   "fmt"
   "errors"
   "io"
   "log"
   "net/http"
   "net/url"
   "os"
   "slices"
)

func segment_template(represent *dash.Representation, ext string) error {
   file1, err := create(ext)
   if err != nil {
      return err
   }
   defer file1.Close()
   if initial := represent.SegmentTemplate.Initialization; initial != "" {
      address, err := initial.Url(represent)
      if err != nil {
         return err
      }
      data, err := get(address)
      if err != nil {
         return err
      }
      data, err = init_protect(data)
      if err != nil {
         return err
      }
      _, err = file1.Write(data)
      if err != nil {
         return err
      }
   }
   key, err := get_key()
   if err != nil {
      return err
   }
   http.DefaultClient.Transport = nil
   var segments []int
   for r := range represent.Representation() {
      segments = slices.AppendSeq(segments, r.Segment())
   }
   var progress xhttp.ProgressParts
   progress.Set(len(segments))
   for _, segment := range segments {
      media, err := represent.SegmentTemplate.Media.Url(represent, segment)
      if err != nil {
         return err
      }
      data, err := get(media)
      if err != nil {
         return err
      }
      progress.Next()
      data, err = write_segment(data, key)
      if err != nil {
         return err
      }
      _, err = file1.Write(data)
      if err != nil {
         return err
      }
   }
   return nil
}

func write_file(name string, data []byte) error {
   log.Println("WriteFile", name)
   return os.WriteFile(name, data, os.ModePerm)
}

func create(name string) (*os.File, error) {
   log.Println("Create", name)
   return os.Create(name)
}

type DashClient interface {
   Mpd() (*http.Response, error)
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

func init() {
   log.SetFlags(log.Ltime)
   xhttp.Transport{}.DefaultClient()
}

const (
   widevine_system_id = "edef8ba979d64acea3c827dcd51d21ed"
   widevine_urn       = "urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"
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
// must return byte slice to cover unwrapping
type WidevineLicense interface {
   License([]byte) ([]byte, error)
}

func get_ext(represent *dash.Representation) (string, error) {
   switch *represent.MimeType {
   case "audio/mp4":
      return ".m4a", nil
   case "text/vtt":
      return ".vtt", nil
   case "video/mp4":
      return ".m4v", nil
   }
   return "", errors.New(*represent.MimeType)
}
