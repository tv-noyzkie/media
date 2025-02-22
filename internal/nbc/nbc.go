package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/nbc"
   "flag"
   "os"
   "path/filepath"
)

func (f *flags) download() error {
   if f.representation != "" {
      f.e.Widevine = nbc.Widevine
      return f.e.Download(f.home, f.representation)
   }
   var metadata nbc.Metadata
   err := metadata.New(f.nbc)
   if err != nil {
      return err
   }
   vod, err := metadata.Vod()
   if err != nil {
      return err
   }
   resp, err := vod.Mpd()
   if err != nil {
      return err
   }
   return internal.Mpd(resp, f.home)
}


type flags struct {
   e              internal.License
   home           string
   nbc            int
   representation string
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
   flag.StringVar(&f.e.ClientId, "c", f.e.ClientId, "client ID")
   flag.StringVar(&f.e.PrivateKey, "p", f.e.PrivateKey, "private key")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.IntVar(&f.nbc, "b", 0, "NBC ID")
   flag.Parse()
   if f.nbc >= 1 {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
