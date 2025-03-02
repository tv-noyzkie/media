package tubi

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "net/url"
   "strconv"
)

// geo block
func (c *Content) New(id int) error {
   req, _ := http.NewRequest("", "https://uapi.adrise.tv/cms/content", nil)
   req.URL.RawQuery = url.Values{
      "content_id": {strconv.Itoa(id)},
      "deviceId":   {"!"},
      "platform":   {"android"},
      "video_resources[]": {
         "dash",
         "dash_widevine",
      },
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(c)
}

type Content struct {
   Children       []*Content
   DetailedType   string `json:"detailed_type"`
   Id             int    `json:",string"`
   SeriesId       int    `json:"series_id,string"`
   // these should already be in reverse order by resolution
   VideoResources []VideoResource `json:"video_resources"`
}

func (c *Content) Series() bool {
   return c.DetailedType == "series"
}

func (c *Content) Episode() bool {
   return c.DetailedType == "episode"
}

type VideoResource struct {
   LicenseServer *struct {
      Url string
   } `json:"license_server"`
   Manifest struct {
      Url string // MPD
   }
   Type       string
}

func (c *Content) Get(id int) (*Content, bool) {
   if c.Id == id {
      return c, true
   }
   for _, child := range c.Children {
      content1, ok := child.Get(id)
      if ok {
         return content1, true
      }
   }
   return nil, false
}

type Byte[T any] []byte

func (v *VideoResource) License(data []byte) ([]byte, error) {
   resp, err := http.Post(
      v.LicenseServer.Url, "application/x-protobuf", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}
