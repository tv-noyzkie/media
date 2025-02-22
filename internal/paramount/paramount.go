package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/paramount"
   "41.neocities.org/x/http"
   "flag"
   "fmt"
   "log"
   "os"
   "path/filepath"
)

func main() {
   http.Transport{}.DefaultClient()
   log.SetFlags(log.Ltime)
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.StringVar(&f.content_id, "b", "", "content ID")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.BoolVar(&f.intl, "n", false, "intl")
   flag.Parse()
   switch {
   case f.content_id != "":
      err := f.do_read()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

func (f *flags) New() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   home = filepath.ToSlash(home)
   f.s.ClientId = home + "/widevine/client_id.bin"
   f.s.PrivateKey = home + "/widevine/private_key.pem"
   return nil
}

type flags struct {
   content_id     string
   intl           bool
   representation string
   s              internal.Stream
}
func (f *flags) do_read() error {
   // item
   var token paramount.AppToken
   if f.intl {
      token = paramount.ComCbsCa
   } else {
      token = paramount.ComCbsApp
   }
   var item paramount.Item
   data, err := item.Marshal(&token, f.content_id)
   if err != nil {
      return err
   }
   err = item.Unmarshal(data)
   if err != nil {
      return err
   }
   // mpd
   represents, err := internal.Mpd(&item)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         // INTL does NOT allow anonymous key request, so if you are INTL you
         // will need to use US VPN until someone codes the INTL login
         f.s.Client, err = paramount.ComCbsApp.Session(f.content_id)
         if err != nil {
            return err
         }
         return f.s.Download(&represent)
      }
   }
   return nil
}
