package mubi

import (
   "bytes"
   "encoding/base64"
   "encoding/json"
   "io"
   "net/http"
   "strings"
)

func (a *Authenticate) Widevine() func([]byte) ([]byte, error) {
   return func(data []byte) ([]byte, error) {
      return a.license(data)
   }
}

func (a *Authenticate) license(data []byte) ([]byte, error) {
   // final slash is needed
   req, err := http.NewRequest(
      "POST", "https://lic.drmtoday.com/license-proxy-widevine/cenc/",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   data, err = json.Marshal(map[string]any{
      "merchant":  "mubi",
      "sessionId": a.Token,
      "userId":    a.User.Id,
   })
   if err != nil {
      return nil, err
   }
   req.Header.Set("dt-custom-data", base64.StdEncoding.EncodeToString(data))
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   if strings.Contains(string(data), forbidden[0]) {
      return nil, forbidden
   }
   var value struct {
      License []byte
   }
   err = json.Unmarshal(data, &value)
   if err != nil {
      return nil, err
   }
   return value.License, nil
}
