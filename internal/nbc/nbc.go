package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/nbc"
   "flag"
   "io"
   "log"
   "os"
   "path/filepath"
)

func (f *flags) do_download() error {
   raw_url, err := os.ReadFile(f.home + "/mpd_url")
   if err != nil {
      return err
   }
   raw_body, err := os.ReadFile(f.home + "/mpd_body")
   if err != nil {
      return err
   }
   var client nbc.Client
   client.New()
   f.s.Client = &client
   return f.s.Bravo(f.representation, raw_body, raw_url)
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
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   err = write_file(f.home + "/mpd_body", data)
   if err != nil {
      return err
   }
   err = write_file(
      f.home + "/mpd_url", []byte(resp.Request.URL.String()),
   )
   if err != nil {
      return err
   }
   return internal.Alfa(data)
}

func write_file(name string, data []byte) error {
   log.Println("WriteFile", name)
   return os.WriteFile(name, data, os.ModePerm)
}

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
