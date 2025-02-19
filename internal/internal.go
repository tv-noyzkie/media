package internal

import (
   "41.neocities.org/dash"
   "41.neocities.org/sofia/file"
   "41.neocities.org/sofia/pssh"
   "41.neocities.org/sofia/sidx"
   "41.neocities.org/widevine"
   xhttp "41.neocities.org/x/http"
   "bytes"
   "encoding/base64"
   "fmt"
   "errors"
   "io"
   "log"
   "net/http"
   "net/url"
   "os"
   "slices"
   "strings"
)

func get_key(
   client_id string,
   key_id []byte,
   license WidevineLicense,
   private_key string,
   pssh1 []byte,
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
   if pssh1 == nil {
      var pssh2 widevine.Pssh
      pssh2.KeyIds = [][]byte{key_id}
      pssh1 = pssh2.Marshal()
   }
   log.Println("PSSH", base64.StdEncoding.EncodeToString(pssh1))
   var module widevine.Cdm
   err = module.New(private_key1, client_id1, pssh1)
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
