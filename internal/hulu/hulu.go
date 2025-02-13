package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/hulu"
   "fmt"
   "io"
   "net/http"
   "os"
   "slices"
)

func (f *flags) download() error {
   data, err := os.ReadFile(f.home + "/hulu.txt")
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
   resp, err := http.Get(play.StreamUrl)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   data, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   var mpd dash.Mpd
   err = mpd.Unmarshal(data)
   if err != nil {
      return err
   }
   represents := slices.SortedFunc(mpd.Representation(),
      func(a, b dash.Representation) int {
         return a.Bandwidth - b.Bandwidth
      },
   )
   for _, represent := range represents {
      switch f.representation {
      case "":
         fmt.Print(&represent, "\n\n")
      case represent.Id:
         f.s.Wrapper = play
         return f.s.Download(&represent)
      }
   }
   return nil
}

func (f *flags) authenticate() error {
   data, err := hulu.Authenticate{}.Marshal(f.email, f.password)
   if err != nil {
      return err
   }
   return os.WriteFile(f.home+"/hulu.txt", data, os.ModePerm)
}
