package main

import (
   "41.neocities.org/media/amc"
   "41.neocities.org/media/internal"
   "errors"
   "flag"
   "log"
   "net/http"
   "os"
   "path/filepath"
)

func (f *flags) download() error {
   data, err := os.ReadFile(f.media + "/amc/Auth")
   if err != nil {
      return err
   }
   var auth amc.Auth
   err = auth.Unmarshal(data)
   if err != nil {
      return err
   }
   data, err = auth.Refresh()
   if err != nil {
      return err
   }
   err = f.write_file("/amc.txt", data)
   if err != nil {
      return err
   }
   err = auth.Unmarshal(data)
   if err != nil {
      return err
   }
   play, err := auth.Playback(f.address)
   if err != nil {
      return err
   }
   source, ok := play.Dash()
   if !ok {
      return errors.New(".Dash()")
   }
   if f.representation != "" {
      f.e.Widevine = play.Widevine(source)
      return f.e.Download(f.media + "/Mpd", f.representation)
   }
   resp, err := http.Get(source.Src)
   if err != nil {
      return err
   }
   return internal.Mpd(f.media + "/Mpd", resp)
}

type flags struct {
   address        amc.Address
   e              internal.License
   email          string
   media           string
   password       string
   representation string
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.address, "a", "address")
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.password, "p", "", "password")
   flag.StringVar(&f.e.ClientId, "c", f.e.ClientId, "client ID")
   flag.StringVar(&f.e.PrivateKey, "k", f.e.PrivateKey, "private key")
   flag.Parse()
   switch {
   case f.email != "":
      err := f.login()
      if err != nil {
         panic(err)
      }
   case f.address[1] != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
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

func (f *flags) write_file(name string, data []byte) error {
   log.Println("WriteFile", f.media + name)
   return os.WriteFile(f.media + name, data, os.ModePerm)
}

func (f *flags) login() error {
   var auth amc.Auth
   err := auth.Unauth()
   if err != nil {
      return err
   }
   data, err := auth.Login(f.email, f.password)
   if err != nil {
      return err
   }
   return f.write_file("/amc/Auth", data)
}
