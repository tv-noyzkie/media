package main

import (
   "41.neocities.org/media/cineMember"
   "41.neocities.org/media/internal"
   "errors"
   "flag"
   "fmt"
   "log"
   "os"
   "path"
   "path/filepath"
)

type flags struct {
   address  cineMember.Address
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
   flag.Var(&f.address, "a", "address")
   flag.StringVar(&f.e.ClientId, "c", f.e.ClientId, "client ID")
   flag.StringVar(&f.email, "e", "", "email")
   flag.StringVar(&f.dash, "i", "", "DASH ID")
   flag.StringVar(&f.e.PrivateKey, "k", f.e.PrivateKey, "private key")
   flag.StringVar(&f.password, "p", "", "password")
   flag.Parse()
   switch {
   case f.password != "":
      err := f.write_user()
      if err != nil {
         panic(err)
      }
   case f.address[0] != "":
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

func (f *flags) write_user() error {
   data, err := cineMember.NewUser(f.email, f.password)
   if err != nil {
      return err
   }
   return f.write_file(f.media + "/cineMember/User", data)
}

func (f *flags) download() error {
   if f.dash != "" {
      f.e.Client = title
      return f.e.Download(&represent)
   }
   data, err := os.ReadFile(f.media + "/cineMember/User")
   if err != nil {
      return err
   }
   var user cineMember.User
   err = user.Unmarshal(data)
   if err != nil {
      return err
   }
   article, err := f.address.Article()
   if err != nil {
      return err
   }
   asset, ok := article.Film()
   if !ok {
      return errors.New(".Film()")
   }
   play, err := user.Play(article, asset)
   if err != nil {
      return err
   }
   title, ok := play.Dash()
   if !ok {
      return errors.New(".Dash()")
   }
   resp, err := http.Get(title.Manifest)
   if err != nil {
      return err
   }
   return internal.Mpd(f.media + "/Mpd", resp)
}
