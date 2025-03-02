package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/tubi"
   "errors"
   "flag"
   "fmt"
   "log"
   "os"
   "path/filepath"
)

type flags struct {
   dash  string
   e     internal.License
   media string
   tubi  int
}

func (f *flags) New() error {
   var err error
   f.media, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   f.media = filepath.ToSlash(f.media) + "/media"
   f.e.ClientId = f.media + "/client_id.bin"
   f.e.PrivateKey = f.media + "/private_key.pem"
   return nil
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.IntVar(&f.tubi, "b", 0, "Tubi ID")
   flag.StringVar(&f.e.ClientId, "c", f.e.ClientId, "client ID")
   flag.StringVar(&f.dash, "i", "", "DASH ID")
   flag.StringVar(&f.e.PrivateKey, "p", f.e.PrivateKey, "private key")
   flag.Parse()
   switch {
   case f.tubi >= 1:
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

func (f *flags) download() error {
   if f.dash != "" {
      f.e.Client = resource
      return f.e.Download(&represent)
   }
   content := &tubi.Content{}
   err = content.New(f.tubi)
   if err != nil {
      return err
   }
   if content.Episode() {
      err = content.New(content.SeriesId)
      if err != nil {
         return err
      }
   }
   if content.Series() {
      var ok bool
      content, ok = content.Get(f.tubi)
      if !ok {
         return errors.New(".Get")
      }
   }
   resp, err := http.Get(content.VideoResources[0].Manifest.Url)
   if err != nil {
      return err
   }
   return internal.Mpd(f.media + "/Mpd", resp)
}
