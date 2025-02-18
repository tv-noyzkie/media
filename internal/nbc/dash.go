package main

import (
   "41.neocities.org/dash"
   "41.neocities.org/media/nbc"
   "fmt"
   "net/url"
   "slices"
)

func alfa(data []byte) error {
   var media dash.Mpd
   err := media.Unmarshal(data)
   if err != nil {
      return err
   }
   represents := slices.SortedFunc(media.Representation(),
      func(a, b dash.Representation) int {
         return a.Bandwidth - b.Bandwidth
      },
   )
   var line bool
   for _, represent := range represents {
      if line {
         fmt.Println()
      } else {
         line = true
      }
      fmt.Println(&represent)
   }
   return nil
}

func (f *flags) bravo(raw_url, raw_body []byte) error {
   var base url.URL
   err := base.UnmarshalBinary(raw_url)
   if err != nil {
      return err
   }
   var media dash.Mpd
   err = media.Unmarshal(raw_body)
   if err != nil {
      return err
   }
   media.Set(&base)
   for represent := range media.Representation() {
      if represent.Id == f.representation {
         var client nbc.Client
         client.New()
         f.s.Client = &client
         return f.s.Download(&represent)
      }
   }
   return nil
}
