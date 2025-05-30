package itv

import (
   "fmt"
   "path"
   "testing"
   "time"
)

var tests = []struct {
   url string
   id  string
}{
   {
      url: "itv.com/watch/goldeneye/18910",
      id:  "18910",
   },
   {
      url: "itv.com/watch/gone-girl/10a5503a0001B",
      id:  "10/5503/0001B",
   },
   {
      url: "itv.com/watch/grace/2a7610",
      id:  "2/7610",
   },
   {
      url: "itv.com/watch/joan/10a3918",
      id:  "10/3918",
   },
}

func TestTitle(t *testing.T) {
   for i, test := range tests {
      if i >= 1 {
         fmt.Println("---------------------------------------------------------")
      }
      titles, err := LegacyId{test.id}.Titles()
      if err != nil {
         t.Fatal(err)
      }
      for i, title1 := range titles {
         if i >= 1 {
            fmt.Println()
         }
         fmt.Println(&title1)
      }
      time.Sleep(time.Second)
   }
}

func TestLegacyId(t *testing.T) {
   for _, test := range tests {
      var id LegacyId
      id.Set(path.Base(test.url))
      fmt.Println(id)
   }
}
