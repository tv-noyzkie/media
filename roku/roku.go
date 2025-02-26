package roku

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

type Playback struct {
   Drm struct {
      Widevine struct {
         LicenseServer string
      }
   }
   Url string
}

type Token struct {
   AuthToken string
}

type Activation struct {
   Code string
}

type Code struct {
   Token string
}

const user_agent = "trc-googletv; production; 0"

func (a *Activation) String() string {
   var b strings.Builder
   b.WriteString("1 Visit the URL\n")
   b.WriteString("  therokuchannel.com/link\n")
   b.WriteString("\n")
   b.WriteString("2 Enter the activation code\n")
   b.WriteString("  ")
   b.WriteString(a.Code)
   return b.String()
}

// code can be nil
func RequestToken(c *Code) (*http.Response, error) {
   req, _ := http.NewRequest("", "https://googletv.web.roku.com", nil)
   req.URL.Path = "/api/v1/account/token"
   req.Header.Set("user-agent", user_agent)
   if c != nil {
      req.Header.Set("x-roku-content-token", c.Token)
   }
   return http.DefaultClient.Do(req)
}

func (t *Token) Read(resp *http.Response) error {
   return json.NewDecoder(resp.Body).Decode(t)
}

func (t *Token) Activation() (*http.Response, error) {
   data, err := json.Marshal(map[string]string{"platform": "googletv"})
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://googletv.web.roku.com/api/v1/account/activation",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "content-type":         {"application/json"},
      "user-agent":           {user_agent},
      "x-roku-content-token": {t.AuthToken},
   }
   return http.DefaultClient.Do(req)
}

func (a *Activation) Read(resp *http.Response) error {
   return json.NewDecoder(resp.Body).Decode(a)
}

func (t *Token) Code(activate *Activation) (*http.Response, error) {
   req, _ := http.NewRequest("", "https://googletv.web.roku.com", nil)
   req.URL.Path = "/api/v1/account/activation/" + activate.Code
   req.Header = http.Header{
      "user-agent":           {user_agent},
      "x-roku-content-token": {t.AuthToken},
   }
   return http.DefaultClient.Do(req)
}

func (c *Code) Read(resp *http.Response) error {
   return json.NewDecoder(resp.Body).Decode(c)
}

func (t *Token) Playback(roku_id string) (*Playback, error) {
   data, err := json.Marshal(map[string]string{
      "mediaFormat": "DASH",
      "providerId":  "rokuavod",
      "rokuId":      roku_id,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://googletv.web.roku.com/api/v3/playback",
      bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "content-type":         {"application/json"},
      "user-agent":           {user_agent},
      "x-roku-content-token": {t.AuthToken},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   play := &Playback{}
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

func (p *Playback) Mpd() (*http.Response, error) {
   return http.Get(p.Url)
}

func (p *Playback) Widevine() func([]byte) ([]byte, error) {
   return func(data []byte) ([]byte, error) {
      resp, err := http.Post(
         p.Drm.Widevine.LicenseServer, "application/x-protobuf",
         bytes.NewReader(data),
      )
      if err != nil {
         return nil, err
      }
      defer resp.Body.Close()
      return io.ReadAll(resp.Body)
   }
}
