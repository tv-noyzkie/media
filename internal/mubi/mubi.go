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

type flags struct {
   address        mubi.Address
   e              internal.License
   media          string
   representation string
   auth           bool
   code           bool
}

func timed_text(url string) error {
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
   flag.StringVar(&f.e.ClientId, "c", f.e.ClientId, "client ID")
   flag.StringVar(&f.e.PrivateKey, "p", f.e.PrivateKey, "private key")
   flag.Parse()
   switch {
   case f.code:
      err := f.write_code()
      if err != nil {
         panic(err)
      }
   case f.auth:
      err := f.write_auth()
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

func (f *flags) write_code() error {
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
   return f.write_file("/mubi/LinkCode", data)
}

///

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
         return timed_text(text.Url)
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
         data, err = os.ReadFile(f.media + "/mubi.txt")
         if err != nil {
            return err
         }
         var auth mubi.Authenticate
         err = auth.Unmarshal(data)
         if err != nil {
            return err
         }
         f.e.Client = &auth
         return f.e.Download(&represent)
      }
   }
   return nil
}

func (f *flags) write_secure() error {
   data, err := os.ReadFile(f.media + "/mubi.txt")
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
   return os.WriteFile(f.media+"/mubi.txt", data, os.ModePerm)
}
