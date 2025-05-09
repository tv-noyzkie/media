package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/kanopy"
   "errors"
   "flag"
   "log"
   "os"
   "path/filepath"
)

type flags struct {
   dash     string
   e        internal.License
   email    string
   media    string
   password string
   video_id int
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
   flag.IntVar(&f.video_id, "b", 0, "video ID")
   flag.StringVar(&f.e.ClientId, "c", f.e.ClientId, "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.dash, "i", "", "DASH ID")
   flag.StringVar(&f.e.PrivateKey, "k", f.e.PrivateKey, "private key")
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
   switch {
   case f.password != "":
      err := f.authenticate()
      if err != nil {
         panic(err)
      }
   case f.video_id >= 1:
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}

func write_file(name string, data []byte) error {
   log.Println("WriteFile", name)
   return os.WriteFile(name, data, os.ModePerm)
}

func (f *flags) authenticate() error {
   data, err := kanopy.NewLogin(f.email, f.password)
   if err != nil {
      return err
   }
   return write_file(f.media + "/kanopy/Login", data)
}

func (f *flags) download() error {
   data, err := os.ReadFile(f.media + "/kanopy/Login")
   if err != nil {
      return err
   }
   var login kanopy.Login
   err = login.Unmarshal(data)
   if err != nil {
      return err
   }
   if f.dash != "" {
      data, err := os.ReadFile(f.media + "/kanopy/Plays")
      if err != nil {
         return err
      }
      var plays kanopy.Plays
      err = plays.Unmarshal(data)
      if err != nil {
         return err
      }
      manifest, _ := plays.Dash()
      f.e.Widevine = func(data []byte) ([]byte, error) {
         return login.Widevine(manifest, data)
      }
      return f.e.Download(f.media + "/Mpd", f.dash)
   }
   member, err := login.Membership()
   if err != nil {
      return err
   }
   data, err = login.Plays(member, f.video_id)
   if err != nil {
      return err
   }
   var plays kanopy.Plays
   err = plays.Unmarshal(data)
   if err != nil {
      return err
   }
   err = write_file(f.media + "/kanopy/Plays", data)
   if err != nil {
      return err
   }
   manifest, ok := plays.Dash()
   if !ok {
      return errors.New(".Dash()")
   }
   resp, err := manifest.Mpd()
   if err != nil {
      return err
   }
   return internal.Mpd(f.media + "/Mpd", resp)
}
