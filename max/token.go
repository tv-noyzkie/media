package max

import (
   "bytes"
   "crypto/hmac"
   "crypto/sha256"
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
   "time"
)

func (d *DefaultToken) Unmarshal() error {
   d.Session.Value = SessionState{}
   d.Session.Value.Set(string(d.Session.Raw))
   var data struct {
      Data struct {
         Attributes struct {
            Token string
         }
      }
   }
   err := json.Unmarshal(d.Token.Raw, &data)
   if err != nil {
      return err
   }
   d.Token.Value = data.Data.Attributes.Token
   return nil
}

type DefaultToken struct {
   Session Value[SessionState]
   Token Value[string]
}

///

func (d *DefaultToken) Login(key PublicKey, login DefaultLogin) error {
   body, err := json.Marshal(login)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("POST", "/login", bytes.NewReader(body))
   if err != nil {
      return err
   }
   req.URL.Host = func() string {
      var b bytes.Buffer
      b.WriteString("https://default.any-")
      b.WriteString(home_market)
      b.WriteString(".prd.api.discomax.com")
      return b.String()
   }()
   req.Header.Set("authorization", "Bearer " + d.Token.Value)
   req.Header.Set("content-type", "application/json")
   req.Header.Set("x-disco-arkose-token", key.Token)
   req.Header.Set("x-disco-client-id", func() string {
      timestamp := time.Now().Unix()
      hash := hmac.New(sha256.New, default_key.Key)
      fmt.Fprintf(hash, "%v:POST:/login:%s", timestamp, body)
      signature := hash.Sum(nil)
      return fmt.Sprintf("%v:%v:%x", default_key.Id, timestamp, signature)
   }())
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b bytes.Buffer
      resp.Write(&b)
      return errors.New(b.String())
   }
   d.Token.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   d.Session.Raw = []byte(resp.Header.Get("x-wbd-session-state"))
   return nil
}

func (d *DefaultToken) New() error {
   req, err := http.NewRequest(
      "", "https://default.any-any.prd.api.discomax.com/token?realm=bolt", nil,
   )
   if err != nil {
      return err
   }
   // fuck you Max
   req.Header.Set("x-device-info", "!/!(!/!;!/!;!)")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b bytes.Buffer
      resp.Write(&b)
      return errors.New(b.String())
   }
   d.Token.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (d *DefaultToken) Playback(web Address) (*Playback, error) {
   body, err := func() ([]byte, error) {
      var p playback_request
      p.ConsumptionType = "streaming"
      p.EditId = web.EditId
      return json.Marshal(p)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://default.any-any.prd.api.discomax.com",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b bytes.Buffer
      b.WriteString("/playback-orchestrator/any/playback-orchestrator/v1")
      b.WriteString("/playbackInfo")
      return b.String()
   }()
   req.Header = http.Header{
      "authorization": {"Bearer " + d.Token.Value},
      "content-type": {"application/json"},
      "x-wbd-session-state": {d.Session.Value.String()},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b bytes.Buffer
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

func (d *DefaultToken) Routes(web Address) (*DefaultRoutes, error) {
   var req http.Request
   req.URL = &url.URL{}
   req.URL.Scheme = "https"
   req.URL.Host = func() string {
      var b strings.Builder
      b.WriteString("https://default.any-")
      b.WriteString(home_market)
      b.WriteString(".prd.api.discomax.com")
      return b.String()
   }()
   req.URL.Path = func() string {
      text, _ := web.MarshalText()
      var b strings.Builder
      b.WriteString("/cms/routes")
      b.Write(text)
      return b.String()
   }()
   req.URL.RawQuery = url.Values{
      "include": {"default"},
      // this is not required, but results in a smaller response
      "page[items.size]": {"1"},
   }.Encode()
   req.Header = http.Header{
      "authorization": {"Bearer " + d.Token.Value},
      "x-wbd-session-state": {d.Session.Value.String()},
   }
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   route := &DefaultRoutes{}
   err = json.NewDecoder(resp.Body).Decode(route)
   if err != nil {
      return nil, err
   }
   return route, nil
}

func (d *DefaultToken) decision() (*default_decision, error) {
   body, err := json.Marshal(map[string]string{
      "projectId": "d8665e86-8706-415d-8d84-d55ceddccfb5",
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://default.any-any.prd.api.discomax.com",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("authorization", "Bearer " + d.Token.Value)
   req.URL.Path = "/labs/api/v1/sessions/feature-flags/decisions"
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   decision := &default_decision{}
   err = json.NewDecoder(resp.Body).Decode(decision)
   if err != nil {
      return nil, err
   }
   return decision, nil
}

func (s SessionState) Set(text string) error {
   for text != "" {
      var key string
      key, text, _ = strings.Cut(text, ";")
      key, value, _ := strings.Cut(key, ":")
      s[key] = value
   }
   return nil
}

func (s SessionState) String() string {
   var (
      b strings.Builder
      sep bool
   )
   for key, value := range s {
      if sep {
         b.WriteByte(';')
      } else {
         sep = true
      }
      b.WriteString(key)
      b.WriteByte(':')
      b.WriteString(value)
   }
   return b.String()
}

type SessionState map[string]string

func (s SessionState) Delete() {
   for key := range s {
      switch key {
      case "device", "token", "user":
      default:
         delete(s, key)
      }
   }
}

type Value[T any] struct {
   Value T
   Raw []byte
}
