package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/itv"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/http/cookiejar"
   "path"
)

func (f *flags) download() error {
   var id itv.LegacyId
   err := id.Set(path.Base(f.address))
   if err != nil {
      return err
   }
   discovery, err := id.Discovery()
   if err != nil {
      return err
   }
   play, err := discovery.Playlist()
   if err != nil {
      return err
   }
   file, ok := play.Resolution1080()
   if !ok {
      return errors.New("resolution 1080")
   }
   http.DefaultClient.Jar, err = cookiejar.New(nil)
   if err != nil {
      return err
   }
   resp, err := http.Get(file.Href.Data)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return errors.New(resp.Status)
   }
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   reps, err := dash.Unmarshal(data, resp.Request.URL)
   if err != nil {
      return err
   }
   for _, rep := range reps {
      switch f.representation {
      case "":
         fmt.Print(&rep, "\n\n")
      case rep.Id:
         f.s.Namer = itv.Namer{discovery}
         f.s.Wrapper = file
         return f.s.Download(rep)
      }
   }
   return nil
}
