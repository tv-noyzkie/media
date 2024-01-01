package youtube

import (
   "154.pages.dev/stream"
   "fmt"
   "testing"
   "time"
)

var ids = []string{
   "2ZcDwdXEVyI", // episode
   "7KLCti7tOXE", // video
   "R9lZ8i8El4I", // film
}

func Test_Watch(t *testing.T) {
   for _, id := range ids {
      c, err := make_contents(id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(stream.Name(c))
      time.Sleep(time.Second)
   }
}