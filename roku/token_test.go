package roku

import (
   "fmt"
   "os"
   "testing"
)

var tests = map[string]struct {
   key string
   key_id string
   url string
} {
   "episode": {
      key: "e258b67d75420066c8424bd142f84565",
      key_id: "bdfa4d6cdb39702e5b681f90617f9a7e",
      url: "therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76",
   },
   "movie": {
      key: "13d7c7cf295444944b627ef0ad2c1b3c",
      url: "therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f",
   },
}

func TestTokenWrite(t *testing.T) {
   var err error
   // AccountAuth
   var auth AccountAuth
   auth.Data, err = os.ReadFile("auth.txt")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   // AccountCode
   var code AccountCode
   code.Data, err = os.ReadFile("code.txt")
   if err != nil {
      t.Fatal(err)
   }
   code.Unmarshal()
   // AccountToken
   token, err := auth.Token(code)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.txt", token.Data, os.ModePerm)
}

func TestTokenRead(t *testing.T) {
   var err      error
   // AccountToken
   var token AccountToken
   token.Data, err = os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   token.Unmarshal()
   // AccountAuth
   var auth AccountAuth
   err = auth.New(&token)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", auth)
}