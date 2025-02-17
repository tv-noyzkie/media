package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/internal"
   "41.neocities.org/media/nbc"
   "flag"
   "fmt"
   "io"
   "log"
   "net/url"
   "path/filepath"
   "slices"
   "41.neocities.org/x/http"
   "os"
)

func write_file(name string, data []byte) error {
   log.Println("WriteFile", name)
   return os.WriteFile(name, data, os.ModePerm)
}

func main() {
   http.Transport{}.DefaultClient()
   log.SetFlags(log.Ltime)
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
   nbc            int
   representation string
   s              internal.Stream
   home string
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

func (f *flags) do_print() error {
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
   defer resp.Body.Close()
   // Body
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   err = write_file(f.home + "/mpd_body", data)
   if err != nil {
      return err
   }
   // Request.URL
   err = write_file(
      f.home + "/mpd_url", []byte(resp.Request.URL.String()),
   )
   if err != nil {
      return err
   }
   var media dash.Mpd
   err = media.Unmarshal(data)
   if err != nil {
      return err
   }
   represents := slices.SortedFunc(media.Representation(),
      func(a, b dash.Representation) int {
         return a.Bandwidth - b.Bandwidth
      },
   )
   var line bool
   for _, represent := range represents {
      if line {
         fmt.Println()
      } else {
         line = true
      }
      fmt.Println(&represent)
   }
   return nil
}

func (f *flags) do_download() error {
   // Body
   data, err := os.ReadFile(f.home + "/mpd_body")
   if err != nil {
      return err
   }
   var media dash.Mpd
   err = media.Unmarshal(data)
   if err != nil {
      return err
   }
   // Request.URL
   data, err = os.ReadFile(f.home + "/mpd_url")
   if err != nil {
      return err
   }
   var base url.URL
   err = base.UnmarshalBinary(data)
   if err != nil {
      return err
   }
   media.Set(&base)
   for represent := range media.Representation() {
      if represent.Id == f.representation {
         var client nbc.Client
         client.New()
         f.s.Client = &client
         return f.s.Download(&represent)
      }
   }
   return nil
}
