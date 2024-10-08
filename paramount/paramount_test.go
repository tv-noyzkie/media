package paramount

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

var tests = map[string]struct{
   content_id string
   key_id string
   url string
}{
   "fr": {
      content_id: "Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ",
      key_id: "06c3b7eea1ce45779faee2abc8d01a55",
      url: "paramountplus.com/movies/video/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ",
   },
   "us": {
      content_id: "esJvFlqdrcS_kFHnpxSuYp449E7tTexD",
      key_id: "1fde0154d72a4f45912b34f0ce0777eb",
      url: "paramountplus.com/shows/video/esJvFlqdrcS_kFHnpxSuYp449E7tTexD",
   },
}

func TestWidevine(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      var pssh widevine.Pssh
      pssh.ContentId = []byte(test.content_id)
      pssh.KeyId, err = hex.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Cdm
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      var app AppToken
      app.ComCbsApp()
      session, err := app.Session(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      key, err := module.Key(session, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}
