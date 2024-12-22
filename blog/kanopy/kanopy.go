package kanopy

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

type videos_response struct {
   Video *struct {
      ProductionYear int
      Title string
   }
}

func (w *web_token) videos(id int64) (*videos_response, error) {
   req, err := http.NewRequest("", "https://www.kanopy.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/kapi/videos/" + strconv.FormatInt(id, 10)
   req.Header = http.Header{
      "authorization": {"Bearer " + w.Jwt},
      "user-agent": {user_agent},
      "x-version": {x_version},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   videos := &videos_response{}
   err = json.NewDecoder(resp.Body).Decode(videos)
   if err != nil {
      return nil, err
   }
   return videos, nil
}
type video_manifest struct {
   DrmLicenseId string
   ManifestType string
   Url string
}

type video_plays struct {
   Manifests []video_manifest
}

func (w *web_token) plays(
   member *membership, video_id int64,
) (*video_plays, error) {
   data, err := json.Marshal(map[string]int64{
      "domainId": member.DomainId,
      "userId": w.UserId,
      "videoId": video_id,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.kanopy.com/kapi/plays", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "authorization": {"Bearer " + w.Jwt},
      "content-type": {"application/json"},
      "user-agent": {user_agent},
      "x-version": {x_version},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   play := &video_plays{}
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

func (w *web_token) membership() (*membership, error) {
   req, err := http.NewRequest("", "https://www.kanopy.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/kapi/memberships"
   req.URL.RawQuery = "userId=" + strconv.FormatInt(w.UserId, 10)
   req.Header = http.Header{
      "authorization": {"Bearer " + w.Jwt},
      "user-agent": {user_agent},
      "x-version": {x_version},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var member struct {
      List []membership
   }
   err = json.NewDecoder(resp.Body).Decode(&member)
   if err != nil {
      return nil, err
   }
   return &member.List[0], nil
}

type membership struct {
   DomainId int64
}

func (v *video_plays) dash() (*video_manifest, bool) {
   for _, manifest := range v.Manifests {
      if manifest.ManifestType == "dash" {
         return &manifest, true
      }
   }
   return nil, false
}

const x_version = "!/!/!/!"

const user_agent = "!"

// good for 10 years
type web_token struct {
   Jwt string
   UserId int64
}

func (web_token) marshal(email, password string) ([]byte, error) {
   value := map[string]any{
      "credentialType": "email",
      "emailUser": map[string]string{
         "email": email,
         "password": password,
      },
   }
   data, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.kanopy.com/kapi/login", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "content-type": {"application/json"},
      "user-agent": {user_agent},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

func (w *web_token) unmarshal(data []byte) error {
   return json.Unmarshal(data, w)
}

func (*poster) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (*poster) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (p *poster) RequestHeader() (http.Header, error) {
   h := http.Header{
      "authorization": {"Bearer " + p.token.Jwt},
      "user-agent": {user_agent},
      "x-version": {x_version},
   }
   return h, nil
}

type poster struct {
   manifest *video_manifest
   token *web_token
}

func (p *poster) RequestUrl() (string, bool) {
   var u url.URL
   u.Scheme = "https"
   u.Host = "www.kanopy.com"
   u.Path = "/kapi/licenses/widevine/" + p.manifest.DrmLicenseId
   return u.String(), true
}
