package cineMember

import (
   "fmt"
   "os"
   "os/exec"
   "strings"
   "testing"
)

const test_url = "cinemember.nl/films/american-hustle"

func Test(t *testing.T) {
   data, err := os.ReadFile("authenticate.txt")
   if err != nil {
      t.Fatal(err)
   }
   var user Authenticate
   err = user.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   var web Address
   err = web.Set(test_url)
   if err != nil {
      t.Fatal(err)
   }
   _, err = web.Article()
   if err != nil {
      t.Fatal(err)
   }
}
