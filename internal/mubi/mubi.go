package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/mubi"
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
   flag.Var(&f.address, "a", "address")
   flag.BoolVar(&f.auth, "auth", false, "authenticate")
   flag.BoolVar(&f.code, "code", false, "link code")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.BoolVar(&f.secure, "w", false, "secure URL")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.Parse()
   switch {
   case f.auth:
      err := f.write_auth()
      if err != nil {
         panic(err)
      }
   case f.code:
      err := write_code()
      if err != nil {
         panic(err)
      }
   case f.secure:
      err := f.write_secure()
      if err != nil {
         panic(err)
      }
   case f.address.String() != "":
      err := f.download()
      if err != nil {
         panic(err)
      }
   default:
      flag.Usage()
   }
}
type flags struct {
   address        mubi.Address
   auth           bool
   code           bool
   home           string
   representation string
   s              internal.Stream
   secure         bool
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
func (f *flags) download() error {
   data, err := os.ReadFile(f.address.String() + ".txt")
   if err != nil {
      return err
   }
   var secure mubi.SecureUrl
   err = secure.Unmarshal(data)
   if err != nil {
      return err
   }
   for _, text := range secure.TextTrackUrls {
      switch f.representation {
      case "":
         fmt.Print(&text, "\n\n")
      case text.Id:
         return f.timed_text(text.Url)
      }
   }
   // github.com/golang/go/issues/18639
   // we dont need this until later, but you have to call before the first
   // request in the program
   os.Setenv("GODEBUG", "http2client=0")
   represents, err := internal.Mpd(&secure)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         data, err = os.ReadFile(f.home + "/mubi.txt")
         if err != nil {
            return err
         }
         var auth mubi.Authenticate
         err = auth.Unmarshal(data)
         if err != nil {
            return err
         }
         f.s.Client = &auth
         return f.s.Download(&represent)
      }
   }
   return nil
}

func (f *flags) write_secure() error {
   data, err := os.ReadFile(f.home + "/mubi.txt")
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
   data, err = mubi.SecureUrl{}.Marshal(&auth, film)
   if err != nil {
      return err
   }
   return os.WriteFile(f.address.String()+".txt", data, os.ModePerm)
}

func (f *flags) write_auth() error {
   data, err := os.ReadFile("code.txt")
   if err != nil {
      return err
   }
   var code mubi.LinkCode
   err = code.Unmarshal(data)
   if err != nil {
      return err
   }
   data, err = mubi.Authenticate{}.Marshal(&code)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home+"/mubi.txt", data, os.ModePerm)
}

func write_code() error {
   var code mubi.LinkCode
   data, err := code.Marshal()
   if err != nil {
      return err
   }
   err = os.WriteFile("code.txt", data, os.ModePerm)
   if err != nil {
      return err
   }
   err = code.Unmarshal(data)
   if err != nil {
      return err
   }
   fmt.Println(&code)
   return nil
}

func (f *flags) timed_text(url string) error {
   resp, err := http.Get(url)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   file, err := os.Create(".vtt")
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
