package internal

import (
   "41.neocities.org/dash"
   "41.neocities.org/widevine"
   xhttp "41.neocities.org/x/http"
   "bytes"
   "encoding/base64"
   "errors"
   "io"
   "log"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func get_key(
   client_id, private_key string, key_id, pssh []byte,
   license WidevineLicense,
) ([]byte, error) {
   if key_id == nil {
      return nil, nil
   }
   private_key1, err := os.ReadFile(private_key)
   if err != nil {
      return nil, err
   }
   client_id1, err := os.ReadFile(client_id)
   if err != nil {
      return nil, err
   }
   if pssh == nil {
      var pssh1 widevine.Pssh
      pssh1.KeyIds = [][]byte{key_id}
      pssh = pssh1.Marshal()
   }
   log.Println("PSSH", base64.StdEncoding.EncodeToString(pssh))
   var module widevine.Cdm
   err = module.New(private_key1, client_id1, pssh)
   if err != nil {
      return nil, err
   }
   data, err := module.RequestBody()
   if err != nil {
      return nil, err
   }
   data, err = license.License(data)
   if err != nil {
      return nil, err
   }
   var body widevine.ResponseBody
   err = body.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   block, err := module.Block(body)
   if err != nil {
      return nil, err
   }
   containers := body.Container()
   for {
      container, ok := containers()
      if !ok {
         return nil, errors.New("ResponseBody.Container")
      }
      if bytes.Equal(container.Id(), key_id) {
         key := container.Key(block)
         log.Println("key", base64.StdEncoding.EncodeToString(key))
         return key, nil
      }
   }
}

func get(address *url.URL) ([]byte, error) {
   resp, err := http.Get(address.String())
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var data strings.Builder
      resp.Write(&data)
      return nil, errors.New(data.String())
   }
   return io.ReadAll(resp.Body)
}

func segment_base(
   client_id, private_key string, key_id, pssh []byte,
   ext string,
   license WidevineLicense,
   represent *dash.Representation,
) error {
   file1, err := create(ext)
   if err != nil {
      return err
   }
   defer file1.Close()
   base := represent.SegmentBase
   var req http.Request
   req.Header = http.Header{}
   // need to use Set for lower case
   req.Header.Set("range", "bytes="+base.Initialization.Range.String())
   req.URL = represent.BaseUrl[0]
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusPartialContent {
      return errors.New(resp.Status)
   }
   data, err := io.ReadAll(resp.Body)
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
   key, err := get_key(client_id, private_key, key_id, pssh, license)
   if err != nil {
      return err
   }
   references, err := write_sidx(&req, base.IndexRange)
   if err != nil {
      return err
   }
   http.DefaultClient.Transport = nil
   var progress xhttp.ProgressParts
   progress.Set(len(references))
   for _, reference := range references {
      base.IndexRange[0] = base.IndexRange[1] + 1
      base.IndexRange[1] += uint64(reference.Size())
      data, err = func() ([]byte, error) {
         req.Header.Set("range", "bytes="+base.IndexRange.String())
         resp, err := http.DefaultClient.Do(&req)
         if err != nil {
            return nil, err
         }
         defer resp.Body.Close()
         if resp.StatusCode != http.StatusPartialContent {
            return nil, errors.New(resp.Status)
         }
         return io.ReadAll(resp.Body)
      }()
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
