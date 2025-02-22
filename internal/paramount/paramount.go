package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/paramount"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   content_id     string
   intl           bool
   representation string
   e              internal.License
   home           string
}

func (f *flags) New() error {
   var err error
   f.home, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   f.home = filepath.ToSlash(f.home) + "/media"
   f.e.ClientId = f.home + "/client_id.bin"
   f.e.PrivateKey = f.home + "/widevine/private_key.pem"
   return nil
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.StringVar(&f.content_id, "b", "", "content ID")
   flag.StringVar(&f.e.ClientId, "c", f.e.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.e.PrivateKey, "p", f.e.PrivateKey, "private key")
   flag.BoolVar(&f.intl, "n", false, "intl")
   flag.Parse()
   switch {
   case f.content_id != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

// INTL does NOT allow anonymous key request, so if you are INTL you
// will need to use US VPN until someone codes the INTL login
func (f *flags) download() error {
   if f.representation != "" {
      session, err := paramount.ComCbsApp.Session(f.content_id)
      if err != nil {
         return err
      }
      f.e.Widevine = session.Widevine()
      return f.e.Download(f.home, f.representation)
   }
   var token paramount.AppToken
   if f.intl {
      token = paramount.ComCbsCa
   } else {
      token = paramount.ComCbsApp
   }
   item, err := token.Item(f.content_id)
   if err != nil {
      return err
   }
   resp, err := item.Mpd()
   if err != nil {
      return err
   }
   return internal.Mpd(resp, f.home)
}
