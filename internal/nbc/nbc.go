package main

import (
   "41.neocities.org/media/internal"
   "41.neocities.org/media/nbc"
   "fmt"
   "io"
   "net/url"
   "os"
)

func (f *flags) do_print() error {
   var metadata nbc.Metadata
   err := metadata.New(f.nbc)
   if err != nil {
      return err
   }
   vod, err := metadata.Vod()
   if err != nil {
      return err
   }
   //////////////////////////////////////////////////////////////////////////////
   resp, err := vod.Mpd()
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   body, err := io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   internal.WriteFile(f.home + "/mpd_body", body)
   file, err := internal.Create(f.home + "/mpd_url")
   if err != nil {
      return err
   }
   defer file.Close()
   fmt.Fprint(file, resp.Request.URL)
   return f.s.Download(resp.Request.URL, body, "")
   //////////////////////////////////////////////////////////////////////////////
}

func (f *flags) do_download() error {
   var client nbc.Client
   client.New()
   f.s.Client = &client
   //////////////////////////////////////////////////////////////////////////////
   data, err := os.ReadFile(f.home + "/mpd_url")
   if err != nil {
      return err
   }
   var base url.URL
   err = base.UnmarshalBinary(data)
   if err != nil {
      return err
   }
   body, err := os.ReadFile(f.home + "/mpd_body")
   if err != nil {
      return err
   }
   return f.s.Download(&base, body, f.representation)
   //////////////////////////////////////////////////////////////////////////////
}
