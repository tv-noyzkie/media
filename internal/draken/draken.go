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

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.StringVar(&f.password, "p", "", "password")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.s.PrivateKey, "k", f.s.PrivateKey, "private key")
   flag.StringVar(&f.address, "a", "", "address")
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

func (f *flags) New() error {
   var err error
   f.home, err = os.UserHomeDir()
   if err != nil {
      return err
   }
   f.home = filepath.ToSlash(f.home)
   f.s.ClientId = f.home + "/widevine/client_id.bin"
   f.s.PrivateKey = f.home + "/widevine/private_key.pem"
   return nil
}

type flags struct {
   address        string
   email          string
   home           string
   password       string
   representation string
   s              internal.Stream
}
func (f *flags) download() error {
   data, err := os.ReadFile(f.home + "/draken.txt")
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
   title, err := login.Entitlement(&movie)
   if err != nil {
      return err
   }
   play, err := login.Playback(&movie, title)
   if err != nil {
      return err
   }
   represents, err := internal.Mpd(play)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Client = &draken.Client{&login, play}
         return f.s.Download(&represent)
      }
   }
   return nil
}

func (f *flags) authenticate() error {
   data, err := draken.Login{}.Marshal(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home+"/draken.txt", data, os.ModePerm)
}
