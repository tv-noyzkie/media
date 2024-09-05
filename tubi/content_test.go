package tubi

import (
   "154.pages.dev/text"
   "fmt"
   "testing"
   "time"
)

var tests = []struct{
   content_id int
   key_id     string
   url        string
}{
   {
      content_id: 590133,
      url:        "tubitv.com/movies/590133",
   },
   {
      content_id: 200042567,
      url:        "tubitv.com/tv-shows/200042567",
   },
}

func TestContent(t *testing.T) {
   for _, test := range tests {
      content := &VideoContent{}
      err := content.New(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      if content.Episode() {
         err := content.New(content.SeriesId)
         if err != nil {
            t.Fatal(err)
         }
         time.Sleep(time.Second)
         var ok bool
         content, ok = content.Get(test.content_id)
         if !ok {
            t.Fatal("get")
         }
      }
      name, err := text.Name(Namer{content})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
   }
}
