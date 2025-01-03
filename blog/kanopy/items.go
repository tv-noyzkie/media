package kanopy

import (
   "encoding/json"
   "net/http"
   "strconv"
)

func (v video_items) item(id int64) (int, *video_item) {
   for key, value := range v.List {
      if value.Video != nil {
         if value.Video.VideoId == id {
            return key, &value
         }
      } else {
         if value.Playlist.VideoId == id {
            return key, &value
         }
      }
   }
   return -1, nil
}

type video_items struct {
   List []video_item
}

type video_item struct {
   Playlist *struct {
      Title string
      VideoId int64
   }
   Video *struct {
      Title string
      VideoId int64
   }
}

func (v *video_item) String() string {
   var b []byte
   if v.Video != nil {
      b = append(b, "title = "...)
      b = append(b, v.Video.Title...)
      b = append(b, "\nvideo id = "...)
      b = strconv.AppendInt(b, v.Video.VideoId, 10)
   } else {
      b = append(b, "title = "...)
      b = append(b, v.Playlist.Title...)
      b = append(b, "\nvideo id = "...)
      b = strconv.AppendInt(b, v.Playlist.VideoId, 10)
   }
   return string(b)
}

func (w *web_token) items(id int64) (*video_items, error) {
   req, err := http.NewRequest("", "https://www.kanopy.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      b := []byte("/kapi/videos/")
      b = strconv.AppendInt(b, id, 10)
      b = append(b, "/items"...)
      return string(b)
   }()
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
   items := &video_items{}
   err = json.NewDecoder(resp.Body).Decode(items)
   if err != nil {
      return nil, err
   }
   return items, nil
}
