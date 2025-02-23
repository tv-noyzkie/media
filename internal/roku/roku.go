package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/roku"
   "41.neocities.org/x/http"
   "flag"
   "fmt"
   "log"
   "os"
   "path/filepath"
)

func (f *flags) download() error {
   if f.representation != "" {
      f.e.Widevine = play.Widevine()
      return f.e.Download(f.home, f.representation)
   }
   var code *roku.Code
   if f.token_read {
      data, err := os.ReadFile(f.home + "/roku.txt")
      if err != nil {
         return err
      }
      code = &roku.Code{}
      err = code.Unmarshal(data)
      if err != nil {
         return err
      }
   }
   var token roku.Token
   data, err := token.Marshal(code)
   if err != nil {
      return err
   }
   err = token.Unmarshal(data)
   if err != nil {
      return err
   }
   play, err := token.Playback(f.roku)
   if err != nil {
      return err
   }
   resp, err := play.Mpd()
   if err != nil {
      return err
   }
   return internal.Mpd(resp, f.home)
}

type flags struct {
   code_write     bool
   home           string
   representation string
   roku           string
   token_read     bool
   token_write    bool
   e internal.License
}

func (f *flags) New() error {
   var err error
   f.home, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   f.home = filepath.ToSlash(f.home) + "/media"
   f.e.ClientId = f.home + "/client_id.bin"
   f.e.PrivateKey = f.home + "/private_key.pem"
   return nil
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.StringVar(&f.roku, "b", "", "Roku ID")
   flag.StringVar(&f.e.ClientId, "c", f.e.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.e.PrivateKey, "k", f.e.PrivateKey, "private key")
   flag.BoolVar(&f.code_write, "code", false, "write code")
   flag.BoolVar(&f.token_write, "token", false, "write token")
   flag.BoolVar(&f.token_read, "t", false, "read token")
   flag.Parse()
   switch {
   case f.code_write:
      err := write_code()
      if err != nil {
         panic(err)
      }
   case f.token_write:
      err := f.write_token()
      if err != nil {
         panic(err)
      }
   case f.roku != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

func (f *flags) write_token() error {
   data, err := os.ReadFile("activation.txt")
   if err != nil {
      return err
   }
   var activation roku.Activation
   err = activation.Unmarshal(data)
   if err != nil {
      return err
   }
   data, err = os.ReadFile("token.txt")
   if err != nil {
      return err
   }
   var token roku.Token
   err = token.Unmarshal(data)
   if err != nil {
      return err
   }
   data, err = roku.Code{}.Marshal(&activation, &token)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home+"/roku.txt", data, os.ModePerm)
}

func write_code() error {
   var token roku.Token
   data, err := token.Marshal(nil)
   if err != nil {
      return err
   }
   err = os.WriteFile("token.txt", data, os.ModePerm)
   if err != nil {
      return err
   }
   err = token.Unmarshal(data)
   if err != nil {
      return err
   }
   var activation roku.Activation
   data, err = activation.Marshal(&token)
   if err != nil {
      return err
   }
   err = os.WriteFile("activation.txt", data, os.ModePerm)
   if err != nil {
      return err
   }
   err = activation.Unmarshal(data)
   if err != nil {
      return err
   }
   fmt.Println(activation)
   return nil
}
