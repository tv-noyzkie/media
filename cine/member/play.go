package member

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
)

// geo block, not x-forwarded-for
func (o *OperationUser) Play(asset *ArticleAsset) (*OperationPlay, error) {
   var body struct {
      Query     string `json:"query"`
      Variables struct {
         ArticleId int `json:"article_id"`
         AssetId   int `json:"asset_id"`
      } `json:"variables"`
   }
   body.Query = query_play
   body.Variables.AssetId = asset.Id
   body.Variables.ArticleId = asset.article.Id
   raw, err := json.Marshal(body)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://api.audienceplayer.com/graphql/2/user",
      bytes.NewReader(raw),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "authorization": {"Bearer " + o.Data.UserAuthenticate.AccessToken},
      "content-type":  {"application/json"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var play OperationPlay
   play.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   return &play, nil
}

type OperationPlay struct {
   Body struct {
      ArticleAssetPlay struct {
         Entitlements []struct {
            Manifest string
            Protocol string
         }
      }
   }
   Raw []byte
}

func (o *OperationPlay) Unmarshal() error {
   return json.Unmarshal(o.Raw, &o.Body)
}
const query_play = `
mutation($article_id: Int, $asset_id: Int) {
   ArticleAssetPlay(article_id: $article_id asset_id: $asset_id) {
      entitlements {
         ... on ArticleAssetPlayEntitlement {
            manifest
            protocol
         }
      }
   }
}
`

func (o *OperationPlay) Dash() (string, bool) {
   for _, title := range o.Body.ArticleAssetPlay.Entitlements {
      if title.Protocol == "dash" {
         return title.Manifest, true
      }
   }
   return "", false
}
