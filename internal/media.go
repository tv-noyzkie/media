package internal

import (
   "41.neocities.org/dash"
   "41.neocities.org/sofia/pssh"
   "encoding/base64"
   "fmt"
   "io"
   "net/url"
   "os"
   "slices"
)

// try to get PSSH from DASH then MP4
func (a *alfa) dash_pssh(
   client_id, private_key string,
   d_client DashClient, w_client WidevineClient,
   home, id string,
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
      resp, err := d_client.Dash()
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
            return a.segment_base(
               client_id, private_key,
               ext,
               &represent,
               w_client,
            )
         }
         if represent.SegmentList != nil {
            return a.segment_list(
               client_id, private_key,
               ext,
               &represent,
               w_client,
            )
         }
         return a.segment_template(
            ext,
            &represent,
         )
      }
   }
   return nil
}
