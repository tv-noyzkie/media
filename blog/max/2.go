package max

import "net/http"

func (b bolt_token) initiate() (*http.Response, error) {
   req, err := http.NewRequest(
      "POST", prd_api + "/authentication/linkDevice/initiate", nil,
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(b.st)
   req.Header.Set("x-device-info", device_info)
   return http.DefaultClient.Do(req)
}
