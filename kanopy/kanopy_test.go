package kanopy

import (
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "net/http"
   "os"
   "os/exec"
   "strings"
   "testing"
   "time"
)

var tests = []struct {
   key_id   string
   url      string
   video_id int
}{
   {
      key_id:   "DUCS1DH4TB6Po1oEkG9xUA==",
      url:      "kanopy.com/en/product/13808102",
      video_id: 13808102,
   },
   {
      key_id:   "sYcEuBtnTH6Bqn65yIE0Ww==",
      url:      "kanopy.com/en/product/14881167",
      video_id: 14881167,
   },
}

type transport struct{}

func (transport) RoundTrip(req *http.Request) (*http.Response, error) {
   fmt.Println(req.URL)
   return http.DefaultTransport.RoundTrip(req)
}

func Test(t *testing.T) {
   http.DefaultClient.Transport = transport{}
   var token WebToken
   t.Run("Marshal", func(t *testing.T) {
      data, err := exec.Command("password", "kanopy.com").Output()
      if err != nil {
         t.Fatal(err)
      }
      email, password, _ := strings.Cut(string(data), ":")
      data, err = token.Marshal(email, password)
      if err != nil {
         t.Fatal(err)
      }
      os.WriteFile("ignore/token.txt", data, os.ModePerm)
   })
   t.Run("Unmarshal", func(t *testing.T) {
      data, err := os.ReadFile("ignore/token.txt")
      if err != nil {
         t.Fatal(err)
      }
      err = token.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
   })
   var member *Membership
   t.Run("Membership", func(t *testing.T) {
      var err error
      member, err = token.Membership()
      if err != nil {
         t.Fatal(err)
      }
   })
   t.Run("Plays", func(t *testing.T) {
      for _, test := range tests {
         _, err := token.Plays(member, test.video_id)
         if err != nil {
            t.Fatal(err)
         }
         time.Sleep(time.Second)
      }
   })
}
