package main

import (
   "41.neocities.org/media/hulu"
   "41.neocities.org/media/internal"
   "flag"
   "fmt"
   "log"
   "os"
   "path/filepath"
)

type flags struct {
   e              internal.License
   email          string
   entity         hulu.EntityId
   media          string
   password       string
   representation string
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
   flag.Var(&f.entity, "a", "address")
   flag.StringVar(&f.e.ClientId, "c", f.e.ClientId, "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.e.PrivateKey, "k", f.e.PrivateKey, "private key")
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
   switch {
   case f.password != "":
      err := f.authenticate()
      if err != nil {
         panic(err)
      }
   case f.entity[0] != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

func (f *flags) write_file(name string, data []byte) error {
   log.Println("WriteFile", f.media + name)
   return os.WriteFile(f.media + name, data, os.ModePerm)
}

func (f *flags) authenticate() error {
   data, err := hulu.NewAuthenticate(f.email, f.password)
   if err != nil {
      return err
   }
   return f.write_file("/hulu/Authenticate", data)
}

func (f *flags) download() error {
   if f.representation != "" {
      f.e.Client = play
      return f.e.Download(&represent)
   }
   data, err := os.ReadFile(f.media + "/hulu/Authenticate")
   if err != nil {
      return err
   }
   var auth hulu.Authenticate
   err = auth.Unmarshal(data)
   if err != nil {
      return err
   }
   
   deep, err := auth.DeepLink(&f.entity)
   if err != nil {
      return err
   }
   play, err := auth.Playlist(deep)
   if err != nil {
      return err
   }
   represents, err := internal.Mpd(play)
}
