package main

import (
   "41.neocities.org/media/internal"
   "path/filepath"
   "flag"
   "os"
)

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.IntVar(&f.nbc, "b", 0, "NBC ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.Parse()
   if f.representation != "" {
      err := f.do_download()
      if err != nil {
         panic(err)
      }
   } else if f.nbc >= 1 {
      err := f.do_print()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

type flags struct {
   home           string
   nbc            int
   representation string
   s              internal.Stream
}

func (f *flags) New() error {
   var err error
   f.home, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   f.home = filepath.ToSlash(f.home) + "/media"
   f.s.ClientId = f.home + "/client_id.bin"
   f.s.PrivateKey = f.home + "/private_key.pem"
   return nil
}
