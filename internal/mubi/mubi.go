package main

import (
   "41.neocities.org/media/mubi"
   "41.neocities.org/net"
   "flag"
   "fmt"
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
   flag.Func("a", "address", func(data string) error {
      return f.slug.Parse(data)
   })
   flag.BoolVar(&f.auth, "auth", false, "authenticate")
   flag.StringVar(&f.e.ClientId, "c", f.e.ClientId, "client ID")
   flag.BoolVar(&f.code, "code", false, "link code")
   flag.StringVar(&f.dash, "d", "", "dash ID")
   flag.StringVar(&f.e.PrivateKey, "p", f.e.PrivateKey, "private key")
   flag.IntVar(&net.ThreadCount, "t", 1, "thread count")
   flag.BoolVar(&f.text, "text", false, "text track")
   flag.Parse()
   switch {
   case f.code:
      err = f.do_code()
   case f.auth:
      err = f.do_auth()
   case f.slug != "":
      err = f.do_slug()
   case f.dash != "":
      err = f.do_dash()
   default:
      flag.Usage()
   }
   if err != nil {
      panic(err)
   }
}

func (f *flags) do_dash() error {
   if f.text {
      data, err := os.ReadFile(f.media + "/mubi/SecureUrl")
      if err != nil {
         return err
      }
      var secure mubi.SecureUrl
      err = secure.Unmarshal(data)
      if err != nil {
         return err
      }
      for _, text := range secure.TextTrackUrls {
         err := get(&text)
         if err != nil {
            return err
         }
      }
   }
   data, err := os.ReadFile(f.media + "/mubi/Authenticate")
   if err != nil {
      return err
   }
   var auth mubi.Authenticate
   err = auth.Unmarshal(data)
   if err != nil {
      return err
   }
   f.e.Widevine = func(data []byte) ([]byte, error) {
      return auth.Widevine(data)
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

func write_file(name string, data []byte) error {
   log.Println("WriteFile", name)
   return os.WriteFile(name, data, os.ModePerm)
}

func (f *flags) do_code() error {
   data, err := mubi.NewLinkCode()
   if err != nil {
      return err
   }
   var code mubi.LinkCode
   err = code.Unmarshal(data)
   if err != nil {
      return err
   }
   fmt.Println(&code)
   return write_file(f.media+"/mubi/LinkCode", data)
}

func (f *flags) do_auth() error {
   data, err := os.ReadFile(f.media + "/mubi/LinkCode")
   if err != nil {
      return err
   }
   var code mubi.LinkCode
   err = code.Unmarshal(data)
   if err != nil {
      return err
   }
   data, err = code.Authenticate()
   if err != nil {
      return err
   }
   return write_file(f.media+"/mubi/Authenticate", data)
}

type flags struct {
   auth    bool
   code    bool
   dash    string
   e       net.License
   media   string
   slug    mubi.Slug
   text    bool
}

func (f *flags) do_slug() error {
   data, err := os.ReadFile(f.media + "/mubi/Authenticate")
   if err != nil {
      return err
   }
   var auth mubi.Authenticate
   err = auth.Unmarshal(data)
   if err != nil {
      return err
   }
   film, err := f.slug.Film()
   if err != nil {
      return err
   }
   err = auth.Viewing(film)
   if err != nil {
      return err
   }
   data, err = auth.SecureUrl(film)
   if err != nil {
      return err
   }
   err = write_file(f.media+"/mubi/SecureUrl", data)
   if err != nil {
      return err
   }
   var secure mubi.SecureUrl
   err = secure.Unmarshal(data)
   if err != nil {
      return err
   }
   resp, err := http.Get(secure.Url)
   if err != nil {
      return err
   }
   return net.Mpd(f.media+"/Mpd", resp)
}

func get(text *mubi.Text) error {
   resp, err := http.Get(text.Url)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   file, err := os.Create(text.Base())
   if err != nil {
      return err
   }
   defer file.Close()
   _, err = file.ReadFrom(resp.Body)
   if err != nil {
      return err
   }
   return nil
}
