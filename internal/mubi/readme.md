# mubi

starting with Go 1.24 (Feb 2025), you can now do this:

~~~go
var pro http.Protocols
pro.SetHTTP1(true)
http.DefaultClient.Transport = &http.Transport{Protocols: &pro}
~~~

https://github.com/golang/go/issues/18639

## mexico

~~~
"x-forwarded-for": {"149.88.22.158"},
~~~
