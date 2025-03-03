package main

import (
   "41.neocities.org/media/draken"
   "41.neocities.org/media/internal"
   "flag"
   "fmt"
   "os"
   "path"
   "path/filepath"
)

type flags struct {
   address  string
   dash     string
   e        internal.License
   email    string
   media    string
   password string
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
   flag.StringVar(&f.address, "a", "", "address")
   flag.StringVar(&f.e.ClientId, "c", f.e.ClientId, "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.dash, "i", "", "dash ID")
   flag.StringVar(&f.e.PrivateKey, "k", f.e.PrivateKey, "private key")
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
   switch {
   case f.password != "":
      err := f.authenticate()
      if err != nil {
         panic(err)
      }
   case f.address != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

func (f *flags) authenticate() error {
   data, err := draken.NewLogin(f.email, f.password)
   if err != nil {
      return err
   }
   log.Println("WriteFile", f.media + "/draken/Login")
   return os.WriteFile(f.media+"/draken/Login", data, os.ModePerm)
}

func (f *flags) download() error {
   if f.dash != "" {
      f.e.Client = &draken.Client{&login, play}
      return f.e.Download(&represent)
   }
   data, err := os.ReadFile(f.media + "/draken/Login")
   if err != nil {
      return err
   }
   var login draken.Login
   err = login.Unmarshal(data)
   if err != nil {
      return err
   }
   var movie draken.Movie
   err = movie.New(path.Base(f.address))
   if err != nil {
      return err
   }
   title, err := login.Entitlement(movie)
   if err != nil {
      return err
   }
   play, err := login.Playback(&movie, title)
   if err != nil {
      return err
   }
   resp, err := http.Get(play.Playlist)
   if err != nil {
      return err
   }
   return internal.Mpd(f.media + "/Mpd", resp)
}
