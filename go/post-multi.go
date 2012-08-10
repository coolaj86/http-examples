// http://golang.org/pkg/net/http/
// http://golang.org/pkg/mime/multipart/#Writer.CreatePart
// http://golang.org/pkg/net/textproto/

package main

import (
  "fmt"
  "bytes"
  "net/http"
  "io"
  "io/ioutil"
  "net/textproto"
  "mime/multipart"
  "encoding/json"
)

type ReqData struct {
  Foo string `json:"foo"`
  Baz string `json:"baz"`
}

// https://gist.github.com/2504488
// https://groups.google.com/forum/?fromgroups#!msg/golang-nuts/Zjg5l4nKcQ0/qkGJWwBAmEUJ%5B1-25%5D
func main() {
  target := "http://localhost:3000"
  //target := "http://foobar3000.com/echo"


  /* Create a buffer to hold this multi-part form */
  body_buf := bytes.NewBufferString("")
  body_writer := multipart.NewWriter(body_buf)
  boundary := body_writer.Boundary()
  fmt.Println(boundary)
  content_type := body_writer.FormDataContentType()
  fmt.Println(content_type)


  /* Create a Form Field */
  field_first_name, err := body_writer.CreateFormField("firstName")
  field_first_name.Write([]byte("AJ"))


  /* Create a Form Field in a simpler way */
  body_writer.WriteField("lastName", "ONeal")


  /* Create a Form File */
  json_data := &ReqData {
    Foo: "bar",
    Baz: "qux",
  }
  json_bytes, err := json.Marshal(json_data)
  if nil != err {
    panic(err.Error())
  }
  filename := "foobar.json"
  file_writer, err := body_writer.CreateFormFile("foofile", filename)
  if nil != err {
    panic(err.Error())
  }
  io.Copy(file_writer, bytes.NewBuffer(json_bytes))


  /* Create a completely custom Form Part (or in this case, a file) */
  // http://golang.org/src/pkg/mime/multipart/writer.go?s=2274:2352#L86
  mh := make(textproto.MIMEHeader)
  mh.Set("Content-Type", "text/plain")
  mh.Set("Content-Disposition", "form-data; name=\"readme\"; filename=\"README.md\"")
  part_writer, err := body_writer.CreatePart(mh)
  if nil != err {
    panic(err.Error())
  }
  io.Copy(part_writer, bytes.NewBufferString("## Example\nThis is an example README.md\n### Coolness\nCoolness is a priority."))


  /* Close the body and send the request */
  body_writer.Close()
  resp, err := http.Post(target, content_type, body_buf)
  if nil != err {
    panic(err.Error())
  }


  /* Handle the response */
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  if nil != err {
    fmt.Println("errorination happened reading the body", err)
    return
  }

  fmt.Println(string(body[:]))
}
