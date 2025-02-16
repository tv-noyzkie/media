package draken

import (
   "41.neocities.org/widevine"
   "bytes"
   "encoding/base64"
   "fmt"
   "os"
   "os/exec"
   "strings"
   "testing"
   "time"
)

var films = []struct {
   content_id string
   custom_id  string
   key_id     string
   url        string
}{
   {
      content_id: "MjNkM2MxYjYtZTA0ZC00ZjMyLWIwYTYtOTgxYzU2MTliNGI0",
      custom_id:  "moon",
      key_id:     "74/ZQoQJukeOkUjy76DE+Q==",
      url:        "drakenfilm.se/film/moon",
   },
   {
      content_id: "MTcxMzkzNTctZWQwYi00YTE2LThiZTYtNjllNDE4YzRiYTQw",
      key_id:     "ToV4wH2nlVZE8QYLmLywDg==",
      custom_id:  "the-card-counter",
      url:        "drakenfilm.se/film/the-card-counter",
   },
}

func Test(t *testing.T) {
   for _, film := range films {
      var movie FullMovie
      if err := movie.New(film.custom_id); err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", movie)
      time.Sleep(99 * time.Millisecond)
   }
}
