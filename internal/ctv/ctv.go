package main

import (
   "41.neocities.org/media/ctv"
   "41.neocities.org/media/internal"
   "flag"
   "fmt"
   "os"
   "path/filepath"
)

type flags struct {
   address        ctv.Address
   manifest       bool
   representation string
   s              internal.Stream
}

func main() {
   var f flags
   err := f.New()
   if err != nil {
      panic(err)
   }
   flag.Var(&f.address, "a", "address")
   flag.StringVar(&f.s.ClientId, "c", f.s.ClientId, "client ID")
   flag.StringVar(&f.representation, "i", "", "representation")
   flag.BoolVar(&f.manifest, "w", false, "manifest")
   flag.StringVar(&f.s.PrivateKey, "p", f.s.PrivateKey, "private key")
   flag.Parse()
   switch {
   case f.manifest:
      err := f.get_manifest()
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

func (f *flags) New() error {
   home, err := os.UserHomeDir()
   if err != nil {
      return err
   }
   home = filepath.ToSlash(home)
   f.s.ClientId = home + "/widevine/client_id.bin"
   f.s.PrivateKey = home + "/widevine/private_key.pem"
   return nil
}

func (f *flags) get_manifest() error {
   resolve, err := f.address.Resolve()
   if err != nil {
      return err
   }
   axis, err := resolve.Axis()
   if err != nil {
      return err
   }
   content, err := axis.Content()
   if err != nil {
      return err
   }
   data, err := ctv.Manifest{}.Marshal(axis, content)
   if err != nil {
      return err
   }
   return os.WriteFile("manifest.txt", data, os.ModePerm)
}

func (f *flags) download() error {
   data, err := os.ReadFile("manifest.txt")
   if err != nil {
      return err
   }
   var manifest ctv.Manifest
   err = manifest.Unmarshal(data)
   if err != nil {
      return err
   }
   represents, err := internal.Mpd(manifest)
   if err != nil {
      return err
   }
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Client = ctv.Client{}
         return f.s.Download(&represent)
      }
   }
   return nil
}
