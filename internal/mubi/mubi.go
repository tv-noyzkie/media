package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/mubi"
   "41.neocities.org/platform/mullvad"
   "flag"
   "fmt"
   "log"
   "net/http"
   "os"
   "path/filepath"
)

func (f *flags) do_dash() error {
   if f.dash != "" {
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
            err = func() error {
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
            }()
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
   data, err := os.ReadFile(f.media + "/mubi/Authenticate")
   if err != nil {
      return err
   }
   var auth mubi.Authenticate
   err = auth.Unmarshal(data)
   if err != nil {
      return err
   }
   film, err := f.address.Film()
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
   err = write_file(f.media + "/mubi/SecureUrl", data)
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
   return internal.Mpd(f.media+"/Mpd", resp)
}

type flags struct {
   address mubi.Address
   auth    bool
   code    bool
   e       internal.License
   media   string
   dash    string
   text    bool
   mullvad bool
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.address, "a", "address")
   flag.BoolVar(&f.auth, "auth", false, "authenticate")
   flag.StringVar(&f.e.ClientId, "c", f.e.ClientId, "client ID")
   flag.BoolVar(&f.code, "code", false, "link code")
   flag.StringVar(&f.dash, "i", "", "dash ID")
   flag.BoolVar(&f.mullvad, "m", false, "Mullvad")
   flag.StringVar(&f.e.PrivateKey, "p", f.e.PrivateKey, "private key")
   flag.BoolVar(&f.text, "text", false, "text track")
   flag.Parse()
   if f.mullvad {
      http.DefaultClient.Transport = &mullvad.Transport{
         Protocols: &http.Protocols{}, // github.com/golang/go/issues/18639
         Proxy:     http.ProxyFromEnvironment,
      }
      defer mullvad.Disconnect()
   }
   switch {
   case f.code:
      err := f.do_code()
      if err != nil {
         panic(err)
      }
   case f.auth:
      err := f.do_auth()
      if err != nil {
         panic(err)
      }
   case f.address[0] != "":
      err := f.do_dash()
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
   return write_file(f.media + "/mubi/LinkCode", data)
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
   return write_file(f.media + "/mubi/Authenticate", data)
}
