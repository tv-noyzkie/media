package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/tubi"
   "flag"
   "log"
   "net/http"
   "os"
   "path/filepath"
)

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
   flag.IntVar(&internal.ThreadCount, "t", 1, "thread count")
   flag.Parse()
   switch {
   case f.tubi >= 1:
      err := f.do_tubi()
      if err != nil {
         panic(err)
      }
   case f.dash != "":
      err := f.do_dash()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

type flags struct {
   e       internal.License
   media   string
   tubi    int
   dash    string
}

func write_file(name string, data []byte) error {
   log.Println("WriteFile", name)
   return os.WriteFile(name, data, os.ModePerm)
}

func (f *flags) do_tubi() error {
   data, err := tubi.NewContent(f.tubi)
   if err != nil {
      return err
   }
   var content tubi.Content
   err = content.Unmarshal(data)
   if err != nil {
      return err
   }
   err = write_file(f.media+"/tubi/Content", data)
   if err != nil {
      return err
   }
   resp, err := http.Get(content.VideoResources[0].Manifest.Url)
   if err != nil {
      return err
   }
   return internal.Mpd(f.media+"/Mpd", resp)
}

func (f *flags) do_dash() error {
   data, err := os.ReadFile(f.media + "/tubi/Content")
   if err != nil {
      return err
   }
   var content tubi.Content
   err = content.Unmarshal(data)
   if err != nil {
      return err
   }
   f.e.Widevine = func(data []byte) ([]byte, error) {
      return content.VideoResources[0].Widevine(data)
   }
   return f.e.Download(f.media+"/Mpd", f.dash)
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
