// http://golang.org/pkg/net/http/
package main

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "net/url"
)

func main() {
  resp, err := http.PostForm("http://foobar3000.com/echo",
      url.Values{"foo": {"bar"}, "id": {"123"}})

  if nil != err {
    fmt.Println("errorination happened getting the response", err)
    return
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  if nil != err {
    fmt.Println("errorination happened reading the body", err)
    return
  }

  fmt.Println(string(body[:]))
}
