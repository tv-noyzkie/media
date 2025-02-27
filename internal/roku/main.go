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

type flags struct {
   code_write     bool
   e              internal.License
   home           string
   representation string
   roku           string
   token_read     bool
   token_write    bool
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
   flag.BoolVar(&f.code_write, "code", false, "1 write code")
   flag.BoolVar(&f.token_write, "token", false, "2 write token")
   flag.BoolVar(&f.token_read, "t", false, "3 read token")
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
