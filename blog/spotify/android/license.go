package android

import (
   "io"
   "net/http"
   "net/url"
   "strings"
   "bytes"
   "fmt"
)

func license() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "guc3-spclient.spotify.com"
   req.URL.Path = "/widevine-license/v1/audio/license"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Authorization"] = []string{"Bearer BQA9WA-ZkKdTitbqkrs9XWYCkfJwDsCR80eW1LIVD6vEnua0V2g60hLaWb1d-ycakaRskAEboHQ2kS4xsh00BoGk9P-t4Ji_EBwiiWzl2Q18_WAl-USydjKNQqQNC1jd87m2ZvmmbDjJbJWZ05HnnTRQVgzmCUJyHXeS4Slk1yx0K_p-LoWnRw4Stvio0xHHLYJTU_4wxsNmpjdg5B-hnYGRckhJh566YLS4Zc1uBWHWUgc8cRIHNeeILIlm9cE84CcZ37ZhLhLSHiO9GKu69H_3LOkufVUMt5-omiiDlrF1xr-u_Wpv5mDuM7MeI16Gpnxr8vWDM-cAhxEvXDgZ"}
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   var b bytes.Buffer
   res.Write(&b)
   fmt.Printf("%q\n", b.Bytes())
}

var body = strings.NewReader("\b\x01\x12\x94,\x12K\nI\n3\b\x01\x12\x109$\x82\xfe\x9b\xedsr\xd1e}~\"\xf3+y\x1a\aspotify\"\x149$\x82\xfe\x9b\xedsr\xd1e}~\"\xf3+y)\x02\xf3\xbd\x10\x01\x1a\x10Jz1\xb3\xebc\"\x8c\nw\xdf\xf1\xfb\xa9\x16\xc4\x18\x01 \xf2Τ\xaf\x060\x168\xfe\xbe\xe6\x94\fB\xa7+\n\vspotify.com\x12\x10O-'\xd9aYz|\xbd\x8a*44h\xebR\x1a\xf0(*C/\x01\xfb\x84\x1cR\x81n\xb8\xf4\xe2\x7ffJcf\x04\xf49\xf8\x9cm\x10\xb2\x8bN\xfd<\x1c\xc5\xf0\t\x02(\xcb\xd7\xd7\x1ai6\xb5\"p\xfa\xab\xef%\xa7\xfe\x96\x9fGT\x1c\xafn\x83\xc7';\xacڬ\xec\xc54\x90\x7f\"\x9f\xc6Ⱦ\a\x1e\x8c\x1d$?\xf3\x88\xa6\xec\x99\x15eM\xca\xd2iQ\xf6PC\xbc4\xb5\x9c_Q\x8f亘\xd8=d\x13t\xa6\x0f\xdcњ\x18\x0f\xaal8'\xfc\\\x8cN\xaa/a\x94y.8Vy3\x13>I1\xa0\xac!\xb7 Yb\xae4\xa3\xca\x036\xd1җ\x8c#\xdaD.\xddy'\xde8\r\xe1\xfe3\x90G~)\xcc\xe4amyp\x9b|\"\xfe\x8d\x8f\n\xcc\xfb\fj N\xd5S\xf1\a⬂pƚD\xc2\xdeU\xc1\xa5\xff\x8f\xd3\x00\xe7\x85L\x81\xc5/\xfc\xc3\xc4\xd2\xf5\xcb\x11\xe6\xb5\\\nv\xb3y\r\x1f1:\xdaS\xed\x1a\xfb\xdee\xb5\x85\x91\v\xff\xbe~\xa0EX\xe4C\xb7E\xcb\xd1\x1f\xa1\xf13\x05~\xe9\f\xb5\x8a\xd45\a\x82\x00\x9c\x9bT\x942\xf7Ҕ\xf0\xef:aen\x1bc\x06ׁ\xf0\x19,\x1aPR\xb6v\xc5\xeb\xe7-/\x03\n\xe4\x91\xe1ڻ\x19ڗ\x10\x15\x82\x19\xc3B\xd0v\xa6Q\x91T\xcd\xeb<\xa3\xcc\x13\xf8c[$:8\xd0\x1e\x80\xb2\xd51k?\x0e\xfb\xc5LT\x8a\x12\xc8\xff\x99\x06o\xfb\xb1\xb5\xbf\x8d_\xfa\x01\x10/\xfe\xc9d\x01=\xbe\xeaB߂\xf4\xb0\xf3\xa2\x831\xd5\x17\x97n\x13\xb5#\xfb\xe6\xe3\xd4\xddK<\xcd\x00\xe2--\xad\xe8\xc0\xc8R\x99aҟ\xb9;\xb4\xab4\x94\xb1rT\x1aq\xb9\xa2b\xc8vL}1\xe0Wb\xf0\xf4\xe3\xaaJ\x97l\x1d\xe1}\xf2\x96$\xb2w9]\xc4#\xc6\xe1\xfb\xe5v\xbb\x10s\xaeE\xf79\x1a\xa7D\xa46!\xb3\xbd\x88F\xa0\x8f\xe1,,Ee\x14\xff\x91\n\xea\xa0\xe7M\xcdyB\xb3\xe5\xf9,\xb1A۶\"\xd0\x14)\x1bC\x033m\xb6\x9177\x1fn^\xb10F6a\x94\x85\xf5\x86\xc36\x1e\x9d\x1fP\x1b\x01\xa6E\xd1\xc2{i\x93Ja\x01\x1b:\f}\xc7\xfe=\xc2u\xddCs\xd5\x02\xe2\xed\xcc\xc4\xc0\xfd\xa2\xb5ۚ#\x97vJ\x0f\xd8N\x95#\t.\x1ad\x1b\xdbg⍻0\xcd\xf5\xe4\v<\x1a\x0fMi\x1b\x1c\x05\xbf\x88+\b⡫ϗ\x9fù\xa7\x99\x98\xae\xb4\x18\xee\x17յv>\x1c\x9b\x0fv\x9c\xd6xo?<a\x00\\\xda\xf5\xb1\xf8\xae\xce\xcd\v\xb8\xfd\xff\x10!\xadYf\xe1\x9d\xcb@\xad\x98&\xb1\xdd\xc5\xfdO\xcb\xf4\x19\xe5Q\xe4Joö\x12\xbe\xc4Zro&H\xd8\xe0\x99ZH\xe8\xe7\xc9:Aj`cE\a\xc2Y\x82\xf5pW\x8f\xc6\xd9\xee\x10b\xfd\xd1-\xa9oմ\xe3\xe5gzά\x1f\xad\xaeb\xf7<\x01\a\xbd\xa9\xaf7\x9e\x00i\x92ȅ\xa0\x01\x17H\x1fo\xdbx\xe2\x16W\xa0^\xcf$E\xbc\xb6\x8boJB\xe6s\xa8\x7f\x95i\xcc\x10\xa5ב\x8c\xfd\xbdO\xf0\xa8\x03\xe6;\x8ak-:r\xb9\xd9\f\xff?\"M_\x95\x02?\xfcW\xa3\x17\x87\xf3\v\xedb\x8c\xf1Ttŏ\x82/(\xca@\xc1\vշ\t\xdddt\xf2Nh\xdfG\x14\xaa/\xc8oI\x81C\xe3\xeer2\xcbX\xa5\xdcC\xf0\xae[\x86\x02ib\f\x86\xe0\xa2A8P\x16]\xd5h\r\xab\xabC{\xf0>\\~'cTF9\"\x00\xab\x0e_\x9b\x89I\xaf\xf1\xe9\xe5\xb1D'z\x8d\xe0\t\xa0\xa9i\x1c\xf9\xad\xe5\xe5\xf4\xf2\xbb\fd\x10\x81y\xa5TO:\xed\xc6\\ؗ\xa8\xcav\x17\xb0\x87\"GKo\xb2Q*\x85,M\x02\x10\x19\xc9l\x87\x89\xb7ɸ\xa3\x90\xd4sc&\xce}\xa0\xa5\r\xd6b,\xfa\x18\xa6\U0006888eҋ \x9b\xb8mBk\xa6·@9E_\x87j\xaa=\x85=\xfa\xb38ȥ\xb9\x1a\\\xb79\xa1~\xeah\x90\xeb\xe1\x9bNr&\x81|\xbe\xe0xx@0\xc2\xf9\xbcZw\xf3\b2F\xe5x\x02\x19\xabl\xe1\x1fC\xc7z*\xbb\x98G3\x11\xdc\xf1\xdb\x10\xfa\x8e\xee4^e\xac4\xc3\x02\xfd\xfd湘\n\x91l\x14\xa4Ɛ\x12Q\x83\a\xc3\\\xc1\xda\xe4\xad\x01G\x87\xf0\x970u\x88\xdf\xfaW\xa1\xab&\xcd%`#!\xf0\x875|AvF\xb2\x9aOb\x1c\x85\xae\x1c\xe1\f\xe2\xc6\xdc\aEH\x7f\xc8\x00\xebc\n^\x91s\x06\xac\xab\xa9\x9dp\xc1\xec\x91\x1b\x95{\x8d_C\xbb\xa3Ň\xce#1\xce\xf5\x88\x9c\xe2\x1a?\rY\xa7C$\xaa\x039\x9f\x9c\x88ҡo\x852\x9ea\xf7\xf5\x02\xaf\x95\x80w\xdf\xd1)\x1boe\xf0\xcdǢҼ\xae\x97t\x8dߕ\xc2\xee|ȩ\xb8.\xb4B\at\x1b\xe9\xa2tU\xe5DDG\x15\xb6\xa2ʳE\x94\x94\x1e9]}\xed\xba\xd3\x154\xe9\xfe\xc4mc\xb1\xd9I|D\xe4\x05\xb5\xcb\t\x03\xd5\r\xb0\xfd/\x14\b~\xeb\xc1{\x14\x9e&\x96n]\xe3?,\xa6\xe5\x044+\t/\x05٘\xb1\xae\x80\xbc|4{k1+∜K\xef\x87\x1b\xc0\x90kba\x1f~\x1dj\xc5\xf3\x03[\xfd-\xc2r\x14\xfa\r\xc4{\xf0\x1a\xf7\xb3\x96\x9c(\xf9\x97j\xcf\xc3\x02\xfa\x1dH~\xcd\x1aSG\x94fҷyg\x00x\xb6\xba;z#\xc2\xdc[\f\xdf)r\x88\x1f\xaf\x9bN\xe4\x12\xaclO\x9av\x9d\xb3\xf7\xc3Ep\xc2\x0fy\x95\xc7\b\xc9n\xe4\x8b\xc6t\xfa\xc1f:$\xb2B\xa2`p(\x10\xa1\xf8@\xe9ٹ\xba\x86\f\x92\xe7n\x15\x17)\xd3x.\xcam\x19[0g\x14\x8d\xe3\xe5\x0eT\f\xdc\xe9\xde\xe5\xfbE\xcc6e\xb5\xbe\xb2\xfa6\x12\x86\xc0\r\xc5\xed\xa7\x14\x97N\xa7\xe3\x02\a\x00\x80I\x93M\xc8\xfe\xd5#\x9dj\x1c\xe1\u0604\nJp[\xd3\xf4\xb0\xb5\x9c\xc1\xbc\x9f6\x99m;x\x95\xe4\xe2\xd7g\x1b!|\xeb\xd3&\xadΏx\xba\xeflf5\x03\xd7\x03\x96Si\xa8\xa6q\xe5J\x1a\x94\xe0ÿ\r\xe3[,/\x1d\xe36\x9fE\xef\x05\xa2\x9b\xa7\x8e\x8f\xba\x11&\xb5\xaa\xc8\xe4\xbe\xef\xa2Ē\x14g\xd9?\xe5\xafKL뽢\xb5VB\xec\xd6v!\x03\x81\x18\xa4\xa2m\xf3O\xe0){J\x15\x8a\xea \xa7\xa4\x96\xaeF9\xe8\x86\xeemS\x02h\x84\xcc\xda\f\xb3\xc0T]Ճb\x85\x10Q/ֺgo\x0f\x9b\xd0\xc8\xe8\x12l\x85\xad\xb3~B\xfbޢdv<\xd5\x04ݺ\a\xccq\x89\xf3\u009b\xc21aʣ\xbd\xf0<\xba8[\x91V\x8bÌZ\x1f\foa\x967\x81\xbf$\x96\xa5\xa6,\x7fҜ\xd9N/j\xe12\xd3\x05pɁ\xa1[\x96\x99\xfbu~0?\xcc4A\x14\x84\x125\xb0ƭ]/9\x91\xf0EUr\x9a7\x9e\xbd\xfcx\x02\xc8n\xd2G\xed%\xfd:\xeef4\x8dte\xdd\xea+\x05_[\xf3vK\x8d\x1do\x1a\xffXA\x85\x01\x80Q\xeb콝~7\xed\xb5\xe0S̎\xe0\xf2&\xd6&-\x9b\xbbԘ\xb2\xb9\xdb\xc1\xe7M\xdf\xc5\xfe\xe7.Ղ:\x8aH\x00r\xb2\xe1\x9a\xeb N!\x81\xb7)\xe8\xd3\xef\xa9\xe9i\x85[\x83gT\xbc`\x99i\\\xb7\xa7\xb6\x86'\x1a\x14o\x87cd\xa6\x11\xd7l\x86+\xbaǔ[~~C\xe2\xb2\x05\x04=v&\xda}S)`\xc5\x03\xbcݭ\xf3Xk\xacE.\xe2)\x13\x94̭y5g\x9b\x9a\x8c\r\x13\xe6\xc8t\x8f\x83\xeb)H\xce\b\xa0/\xc6\xfe\x8fL\x85\xf6\xa0<.\x7fI\xa75ZPž\xaf\x1bH.\xa2\x87\xa8\xb2b\xaf\xfd\xd3:\x8d\xd5\xe0\xe1Ƈ\x7f\xf8\x01\xb0ʪ\x02\xf4u\xe6\r\xc0\U0008ae18\x19fB\xe8O\t\x88Ea\x88\xdf\x1d\xa1\xf9\xe7\x99\xe4\xc1\x94\x88\x85\xdcp\xbf\x18 S\x1e\xa5\xda\xef\xd5v\xab\x11\xce\x17\x97\x0e\xc9šp\r*\xfc\xb0\xbc\x1bA̞\x96\x91\xf2\x93\x11+\xb1\x11!\x1c\x1czP\xc9\n\xf8\xb5\xa4\x14\xdao\xc7j\x02\xf37J\xc0\xb3\x16\x8b5\xce\xe8Fż+\x10P\xbc\xa7a\x14\xc1jc\xb5X\f\xfab\xebu$\xe4\xa9\f\xf0\xff\x87\x15\xb5ͥ\xb2\xbd\xb5\xda\xe2\xdb\x1f\xd5Q\x9e\x8f@U[\x9f\x8a\xcd\x14\xd0\xdbY\x99\x05c\x98\xc0:\x8aW\xad\\\xa5\xdc/;n\x1f\xff\xa8\b֭\xbf\x92\xa2П\x17%gs\xb3|\xcayQ2\xae.\xe4p\xab\x02\x91\xf3\x88\xf6\xf0\xe3'&\xdbƊ\x12=\xb4\x95\x1aW\xf5\xd3\xc1\xe6\x87@#\xad\x10\x86\xb1\xfcGί\x9f\x16!S췴٘\xb7W\\X\x920%x\x97\xa2\xf4>\xfd\xf8wn0jH|\xb0\xa1>\xb1~\xd8Up\x87\x80\x81\xc9_u\x8f\xf5(B\x98\xbb\xf8\xe2\xa4Z\xfc\x15i\xd1_g#\xac\xe1]\x8fO(\xb9=\xbaI\x99\x96\xc0\x9bn\xbat\x06\xdbCý\x84P@\x91`\v\xac6\xbcV?\xe6Q\xb4\xcco\x89\x915\x9d\xc5\xf7x\x06%\xf0\xd6>J\xa1\x14\x89\x8e\x00\x1a\x1f'I\ue4a4H\x18S\x8c:\xd5\xd9\xf4\xa8QUwM\x1b(\x14\xb4*\x7f3\"A\xe4\x12\xc2\x01T \x11\x05\xa9W\x036J\xb9<\xa8\xbf+\xd69\x90\x9aְ\xd7s>\xd6\xd0IH\xb02ǁ\x80\xaa\xa9\xc0\xa6%\x13\xc0\xf4\xecXe\xa1\xf6\x8b\xd2\fAJ(M\x9e\b\r\x87\xcc^܋\x1dY;\x1e\x1f\vOcr\xf9{\x9a\xbc\xd97\x06\"\a\xce@\x9c8\xa8p\xfa<\x8b\xa3m\xf3~6t)\x03\x89\xdfX\x96\x16\xe8\vg!\xcd\xd1w'\xcc\xd4[\xc0=\x02^3\xf6B\xd8r\xa5\xe6\xb6.\xdaj\x0fc\xcfb\x90L\xe8\aP\v\x86\xa98\xc7k|y>\xdc\xd1o\x86^ߐ\xb6\n\xd2\xf3\xb7(1&\xb0gR\x95\xb9\x80O\xb0\xe4+\xc0\x0e$\xb8-y\x8a\x91\xe8\xb5Y\x04\xc1\xc1\x9c\x96\xc0\x82\xac\xc6?p\xa3\xa8\xd6\xe8y\x17\x8e|\x86\x9e%\x96\xe1J\xa7F4\x9fզ?g7\xc5\x13\x1d\\\x8ae\x9d\x8b\xe5)\xbc\b\xd4! =h\x84\x05I\x8fa\x89A\xe9\x14w\xcf*\xf4٧\x16{(\x97\x19\u008aػ\x0e\xff\x96R\xecbZR\xb5\xaf\x83\x15\xfb\"\"\xed\xabr\xcan\xa3M\xa4&\x9f\x1d<%\x8dA\xef\x85z\xff\x02\xa9\xa9\xf9,5\xcb\x16_\x9b\xffJ\x90@@;\x8a\x8ecdUw@/\x11⊽\xe6\xde\xdd\x10\x97\x954\\\xd6ʕ\xa9\x1a\xe7\xb3r|\nx\x01DЙi\x93\x12\x1f\x8a6\xce[\xb4\x01\x99H\x8a`\\\x86\x9c\xd0?GƬ\x85u\xadN-\xcf\xd8\xf2Ec\x81\xc75D3?/\x86iW'y5\x1b\xb3\xcaV\xa6uE\x8d\"\xe6\x13\x94\xab8\x8b\xb6\x177q-\xb0\xf36x\x7f\x02\xe0\xba>\xe9\xf7aR\xfd\x8c\x03G%:\t\xc2\xce\xec\xecL\x7f?(0>\x1b'\xe8\xd2\xefp>\x8cmi\x1d\x9b\xd6:^\v9M\xe4-\xbf\x18\xd8Wظ\xd6\xed\x0e\x9d\x8e\xc7;\x982\xdbV\xb1\xdf\f|w\xe7\xb4\xfc\xbe\x1c+\xcd-\xb9\fq\x1d\x1e\xb0\xda}i1?g\"\xbc\x05*\xb6\xaa\xa9ZE\xc0\xb4f\xae\xd9QՅ\a\x7f\x06\xe0\xb2\x1cw\x17\x0e\xa6Ѫ\xf8\x15x;=\x1c\x01|ː\xc94b\xeb\x96\xd4_\xfbl8\xd4R\xfa\"O\xa7\xf1\xa6\xd3\xe8v\bm\xbc\x98\xb47\xbd\xe1\a\x8fG\x89\xf8$[\x17\x9bA\xcfL\x97\x87\x12Xm\xa7\xb4A\x84\xe8\xe2}\xc6g\f\xb3\xcc^\u05ec\x86`c.\xc5\xfa\x911#\xc8L\x9b\x9d$$\xbc\xf2\xb2\xf42\x0e\x19RR\x99\x9b\x01\x80\n\xc7\xd5\x15\xc5\xe1\x9a\xfb\xff\xf9O\xf8`\xd7>\xd4\x18*\xb1\xd4p\x138 \xaf\xf4ƞ\x96\xdbʑ\"ZG\xea<\xee\"\xe6md7\xf1\xd5\x00\xbf젌K\x13\xe1\xc5\xfe\x82oͦ\x11=+\x9d\xdd#pf~Zk\xd76K\xc8`\x12\xbcof\x8fN\xd8KKI\xd4}\x14-)}\xf1O\xfc˴#4\xcdYO\xd9\a\x89nQ\x0e\xdd\x12\xd7\xc0\xb8\xc1c¾D`\xfd\xd9U\x94妸\xf2\x19\xc3\nl\x9cVP\xef\xb0\f\xa4\x9b\xc0\x80b\x9f\x90[}y!<\x13\xb7\x01%\x04h\x9d\x06F\x92\x82\xf8_\xaa\xa34ط;Nv\xc7\xfe\xf7\xdd ?\x80\xc6YM\\\x11\xd3p\xcc\x1e\xadTVz\x02\xbb\x15\xbd\n\"r\xc2\xd5nٝ\xbc\xe69F\x83w\x8a\xda5\xe3\\w>\xe6e\xb1\xee\x95\xf7\n\x91,G^\x82/\x7f\xac\xa3\xf4`B\xfa\"\xd8N\xeet%\\\xb1\xf5\xc8˾hrG\x0ey-\xe1\xe9\xe8)\x15\xc0\xdc\x1c\xb1[\xe2\x11$V\x8f\xa5ȝ\xd3\xd5-\r7\xf6Z\xf5\x9eL\xa0\xda\xd5\x0e\x86\\\xb4\x009\xff\v\xc41\x87\xd0\xdb\xe7\x98\xeet\xafH\xbe\xe2\x024\x81\x9c\xd6o\xd0g\xb3<>\x81*\x9bTTa\xc9!vlb\x17%\x1fC;\xbar\x93\xac&\x7f\x90\xe1\xce\\\x00\x03\xff^\xe4|eg)[8\xe2\xa9iC\xe9\xc21Z\xf7sYÁ\x80\xd72\x8e\x1dy\xacz\x0ev\x04\xed\x1cg\xc4\xf1Rh5v<9v\x87ݧ\xd7\xf8X\v\xee\x1e8\xe2;\x9dA<\x1e\x183(\x87\xae~Id\xe0\tlOYvL\xca\v\xf84\x01\x12/\x9f\xb7\x8a\xe6%3\xb6m\x16~\tk\xa3\xa5\xcf'\xaa\xd5đ\xe9\xea\xdcGHz\x0f\xc8\x17\x15i\x8b\xec\x16\xaa\x13\x83]\xa3å\x9d\xdc\xc7x\f*\xdf<\x9e\xf9Oo4<\x92\x86\v\x0f\xd9\xd8\xffS\x9b\xcf\xe9ь\v\x99\x8d\x17\x9eK\xfb\xa7\f\x0e\x18eԥ|'\x17\x8e\x1d\xf77x3\xe5\xe3\xaaj3F\xaa\xd2\x14Q\xcd;\x16\x91\xef^\xc0&]\uf303#\xc1S\xa5T\xaa\x17\x84\x85\xa6\x9e\x9alYh\x91yw7R\xeb\x95\x14\xc1\xc6U\x86/\x88'О\r\xf7:\xb4\xa5\xb5\xbe\xca\xc9\xf0\xe5\x17X\x02\x88f8\x95B8\x01\x87܍ \x96\xcf\xe6\xae\xfbw\xd2\x01\xcf\xe1_dÝ\xe6\xb4a\x94\x8d\xb8d\xce\xff-\xc1헀\xff_\x17\xe4!\xa7\x963iS)\x1c\b\xa2#y\xbf\x0e\x8eֳά\xd9v\xbd\xf5%JI\x17~\x90Ic\xe2W\xe9Ik\xf5\xb6\xeeF\x04\x88\x06\x1f\xf7ce\xb8\xfbYPM\x9e\x8c\x8d\x18Y\x8a\xbe\xec>\x04D\xbd\x95鍬氋\x9b\xb8\xb2\xa3_\x1e\u05ee\xed\x19\xc32\x80\x92;\xa1l\x98\x12\xc3\xe6\xfc!7\xbb{\xde\xde~\x0eP\xe0\xa3\xc7\x06\x91\xf1\x03\x1bmN\x8fe\xf9\xef\xfa\r\x15:\xe0\x7f9\xd55\xed\x1fvx\xb3\xfa\x85M\x7f`\x1f\xb8n\x1c.\x91\xff\xa6\xd5|\\*\x17\xf4\x84\xe7ڕ\xc41\xe6ң\x1a\xd9g\x12ŭ@(\x03\xf4\xa9\xfe\xa4\x11\xbf\xbb\n\"\f\x95\xaf\x93\xd3+\f\xacMW\xaf\"\x80\x8f\xbdl^\x18\x90\x7f\xa9\xb8\xea\x9b\xf0ʮ\x04\x0f+v@L:\xed\xafGvi\x8f\x04\x0f\x12\x89\xd5$\x84YD\x01\xf8Dׄ\xfcI\x86\xdd\xf4!q\x8e\x9a(X\xcf\x10\xe4\x1d\x12\xdf\xc6\xc9\x0e\xaf_\xad/\xf6c\xa4D\x88ֿB7\xc3\xeaD\xa3b\xb8Z\xb4\x9f\xa7\xfe\xe1\xfdo\xb0\xa6n&\xf7\xb3]\x94MK\xecJ#\xc9<\x8e%B:Qo\x00MA;\xfcu͠[I\x1c\x12\xaa\\v\x05A\a\x00\xa1\xab\xfb\xafW0j\x01;\xba\xd2p\uf569LW2~$\x11\x86\xf5\x0fy趖7\xdd+\xed\xff\xe2\"\xb0\xfc\xc8\x1bP\xd2\xce\x06\xd9\xf5\xb3\xb0-h\xae\x04\x02\x94&&\xe3F6\x1a\xb5\x1a`\x03\x86\xeb\xdd\x0e7k\aHė\xdd\x16\x9d~\x99\xa8\xb9A\x87mL\x00\u05f8P\xa7\x83\t \xa5\xc2>fF\x86\x93i\xc5y\xbfMx\x0ff\xb0]Y\x80e\xf1\xbc\x987j\x83\x84\xd2{\x86\xf8ZQmBo\xc3FUǞ\xdd?\fO\x16M\x1e\nKQ8\x11Ҕ\x8d\xef\x86\xd2I#N\v\x02\xb7\x98\x1f\xe3\xba܇h\xcc\xfa\xcfd@\x9f\x16?ObuV6\xab\x7fZ\xe7%\xb8X\x84\xf4\xfa*]\xd5t\x00)\x11\xc8\x1c\x1a\xf8~,ǀ\x18\xa3G\x02\xcb\x19F\x18\x13Nn\xa6):\xf7\xb0\x13\xf2\xa8[\x81+X\xed\xb67}\x82\x87G\x01\x86\x99+48\x1apH%\xddU\xa8n\xae:\xc0\x19\x17'\xd8s@}\a\x16iyx\xa6$\xf5\x9f\xee%\xe1ͧvO\x12\xab\x83n\x8b\x00tyNEn\xf5\x13\xb7;\x88\x82F\x9d\xachd\x8d\xe7s\xd3j\x01\xe0G\xf8\"\xbc\xbc\x9bT'\xcf\xf9\xf8\xdcSA\\ \xaf1\xf1W\xd7\x14\xd3X\xe0ۜ\x8f\xfb\xc3\x17*\xcbɣ\x96\xd4<MaZ&r\xa5P\x96\xf7\x8c:\x9e>\x8b\x840\xde)\x0f\x10CRe\t3E\xe7\xa28id\xfbp\xa1\x12\xda)r\xca\xcaT\xfc\x1c\r\x91\x03H\x7f2\x8dJ'\xe33:VK\x88\xc9\x04\x8d\xef\x99\xeaq/\x95\xb5\xa4\xcb\xe2\xd7(}\xbf6\xb3\x87\ngR^\xe79ϖ\xda\xda\xccӖ\x7f\xb2\xe18\xc6V\x92\xdd\xda\xf6#p>\xe5\x8cP\fS9\xf7r.\xf4\xd6\xee\xa5\xd3\xd9w\\j\xbd\x8f\xe4;{\xbek)\x89$]\xe0?\xc4^\xb2*\xb6\x11ٜ=\x88\x83\xf0U\xed^:\xdd\x1cc\xb5\xe7_\x8a\rS\x93&6\xf0\xf7\xdbr\\\xc33J\xad\x00\x9e@,P\xa55-\xa8-\x1c\x05\xbf\xe8\x85GF%\xce\x12\xd4//\xfe\x11\xd1\x19Z@7\x18\xaa\x16W\xe0'I\xb8\xa6\xb2%\x03\xcc=\x8c\xf7\x02(S\xec;\x8f\xf1\x97p\xa9\x13\x88\xb2\xd4R\uf042\x87\x18\xe5\xf8A\xceڱOe>e{\xc5\x01\x91S\xc7\xc2c\x01'\xda\x06\xb4l\xf4\xb3\x01ۑPqaT\xe8\xc1G\xe7{\xe2\xedǥ\xc2V\x8f\xa4\x15\x1a>\xb6\xca~1\xcc\xf9\xcc\f:\xba\xa1|\xe5\f\xda\tA`\xe6\xeabϦ\xaeuH)\xa3\x97\xb1\x16\x97\xba>\x908\xa7\x95o8Pҥ\xc7\xcc\x18\xd2\xc7\x14\xd1\xfe\xb8\x02\x9b\xcc%L:\x173\xbc\x16\xad\x1f\x04\xd6K\xf6\xa8\x9e\xa6\xc35z\xee\xa3\x02\xea\xf3\x99p\xc1{0r\xca\x18\xad\x1d\x8a:\xe1\xeas\xfd\xa3E\xf13s)\tp\xf7\x8bu\xe3\x95\xc5l\xb6\xb0\xf7\xc7u\xac\x8eê\xf6\x1b>D\x1e\xf0d\xb4\xb6\xef\xb7\xd61\xcc\xf7*\xc4\xce\x7f\xc5h)v\xc7\v\x064OT>\xae\xc1@4,\x9b*\xf2\x97G@}Z\xaex\x81\x1c\x87}\xa6\xe0\xeb\x00aa\xee\xa1?\xc5bѣ_\xd5U\xa9N\xcb\xc1\x90\xeb\xe8h\xf04U\x01\fU\xca\xc6y3\xb7\xb2\x82F\xfed\xa3=\x18\xc5\x16J\x9fIzgB\x03\x94\xb6\x9f\xb4\xea\xec\xe9\xcb\xf9\xee\bHN\x1a \n`Bw\xa9\xb48l#\x7f\x94PЁ\xe6\x8e\xc2\xcc\xfcm\xa8Y4\xa0%U:#\xf6\xfa\x15\xbb=\x15*\x918^:\xc0\xab<B\xf2\xban\x9c\xb8\x10\xf6/\xf6,\xec+}\xb9\xa1\xffsS1\x89]\xb5\x13#\xedt\xba\b\xcdA\x9a\xc0\xfb\xe0N\xf5\xcc\t\x9f\x9b\x1cº\xb6\x85'\t8K#̱\xfe8\xf5\xd8X}\x9b\xedǹ@\xa1\xca\xf3\x7f\x99\x97:\x91bL|ZL\x92/ޟk61Q\x94m'\x8dx\x01\x1c\xa3h\xaf\xa4}S\x05\x7f\x9dza\xebU\x93\xa6\xed\xdbL;\x85\xee\xcc@qj\x8ef\x11\xd7\xd6\xe5\xd0\xec\xd7E\x94M\xa9\xe9\xa3ҥ\xe0ò\xe3\x8ao9\xe5\xe9R\xe8\x87U?D8\xa8\x94^\xfc\xabC\xc0\xe3\x14\x04\x98\x04\xd9\xfe_\xd4\x05\f\xee\xc1\xc2\xf0\xc1\x94W\xd6\x04r\xb38\xe3y\xd5\xf5v>\xf9\x92\v\xa0|\xdaL\xa4\x0e\xc7g\xbdO\x11*C\x86u\x99\xbe\xf1=\xd4\x01\xbaS\x83kku\x81}\x8f\x9b>\xb7ܨ\xa0\x89,\xf8rc\xb7\x7f̩\xbaZb\xf9\x8e\xa63VF\xa2\xfaݒ\xebA\xf0\x84\xcd\xc30\x99\x8b\x8a\x9a\xab\xb6\xac\xb0B\x12\xddz\xb1\rȴ\xd9\xfe\a\xa9\xb5-\xa8Lf\xa4R\x8b\xf5\x05\xe9\x91\tA(\f\xfa\x16\xec1h\xce\x12\xa7\xad@Q\xd4\xfd\xdb峡S~\f\xf90\xc7Y\xf3\xda[\x85\xfb\x9ewR\x01\xe2\x93\x18\xbf\br\x06\xeaN\xf6\x98\x92\xe2\x91&*Je|\xcb\x0fŊ\xd1\xcdoո\x8b\x822A\x80\xe6\xd5\x04\xe0#\xc9\xc8\xd29\x9c:\x8b\x9eu\x1a\x9c\x91E\x1bc8r\x0e?UAA\xb9o/n\x14,\xd3\xd6\xe1\xe0\xd6\xffs\x93\xd5ä\xa9\xa5p(\x85\xe2\xf5l\x83ހ(;\xa0t0\x015\x87|\x84\x99\x98\xefp\xa3\x19c\xbb\"\x10\xbezwj\u0e67\x9f\xd1gH\xc2\xdd\xca\t\xde*\x80\x02x3\x14fb\x8b\x1aH\x8ei\v\xdf\xe4h\xb8\x1a\xb6\x01ҕ\xa4\x8d\xb0\x1d\xa4\xcf\xf8\"\xc5\x03=\x05\x1c\xd01\x83\xcei\xa7\xc2\\\xe6-\xb8\x9e\x12\x80\x964S\xe41?\x18\x16\xad\xc7\xd6\u00ad\x92v\xd1I\x8b\xb1z\xdfT\x8d\xdbW\U000ae413]\xccs\x06\xaf\x19T\xe6k\x92J\xfbC\xe0\xd6\x00\x0fB\tj&\x10\v\xd5\xd7R.S\xba\x95/\x94\xa4Q\xd4'\x9c\x11s\x90^\x96!\x05$\x96\x0e\xed\x99\x0e^KMw8j$\xe1s\x19\xdem\xb7-\xf2\xd8߮\x8c\v\xeey\xca익\xc7/\"!ЦƯcܲ&\a^q\x10\ue69a\xfc\x88\xea\x03\xc5wN9\x9c\x01\xb3c\xf1\x8d\x02Bo\xeb-YИ\xb7{\xa3\xbee\xbc\x97;1g\xb1d=SĵT\xcbO\xb4\x8a\x83\xffFz\xcaiy\xce41ٱ:\xdd\xec\x9e2<\xb0\xb6v5B\xe0P\xef&)\xe4,\xc8<\x05\xcf1\x9c\x12É$\x98\x02J\v4.10.2710.0\x1a\x80\x01X\xb8\xdal\x86\x8dc\xa0\x81$\xfa\xc6A,\x96\x1a\xb0\xf2\x1b\xe8m\x98\x1e\xf9\xc6\xd5q=\x10\bK\xd0SQ\xea\u0605\x97Rv\x86\t\xbf&T\xedL\xc4-\t\x1be\xccW\xad\xb9\f\x98\xfe\xa9\xe91\x86$\xf5d8\xfc9U\x9ff\x88\xdb\x16\x1b\xff\xdbu\x87\xce2U\x9b1Za,G\xe1_\x88k\xe3\x02\x14\xba\x00s\xf32\xb4_(lވ\x9c\x01\x85\x02\xd4\xd2\xcd\xe3MH\xd4oz絊|m\x19\b-J\x14\x00\x00\x00\x01\x00\x00\x00\x14\x00\x05\x00\x10\u0099\x9f~\x1cRi\xff")
