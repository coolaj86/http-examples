HTTP Echo Server
===

  * <http://foobar3000.com>
  * <http://helloworld3000.com>

HTTP Examples
====

I've included some raw http transcriptions to help myself and others understand HTTP by observation in addition to scrutiny of the obfuscated RFCs.

Each example contains a README of its own.

Notes:

  * I interpret `CRLF` as `\r\n` which is represented as follows:
    * `^M` represents `\r` or `CR` (as seen in vim, tthe supreme text editor)
    * `newline`s are represented as new lines rather than `\n` or `LF`

HTTP format in a nutshell
====

The basic format of an HTTP request is `<REQUEST><HEADERS><BODY>`

http request
---

Example:

    POST / HTTP/1.1

Explanation:

  * Essentially `<REQUEST>` is `<HTTP-VERB> <RESOURCE> <PROTOCOL>/<VERSION>`
    * Clients should use a single space between tokens
    * Servers should discard extra whitespace
    * The `<HTTP-VERB>` is one of `HEAD`, `OPTIONS`, `GET`, `POST`, `PUT`, `DELETE`, or possibly a rarer or custom verb

  * `<REQUEST>` details which resource should be retrieved (in this case `/`, the document root)

  * `<REQUEST>` is unique in that it does **not** begin with `<CRLF>`
    * like most other directives, `<REQUEST>` does **not** end with `<CRLF>` 

http headers format
----

Example:

    
    ^M
    Host: localhost:3000^M
    Keep-Alive: 115^M
    Connection: keep-alive^M
    Content-Type: multipart/form-data; boundary=---------------------------114772229410704779042051621609^M
    Content-Length: 1751^M
    

Note: there is a trailing new line after the last `^M` that may not show up

Explanation:

  * This example shows a `multipart` form `POST`
    * `Content-Type: multipart/form-data; boundary=---------------------------114772229410704779042051621609`

  * `<CRLF>` **preceedes** each `<HTTP-HEADER>`

  * The entire `<HTTP-HEADERS>` block is ended by a trailing `<CRLF>`

http body format
---

Assuming that you consider your `<HEADERS>` to be followed by `<CRLF>`, the `<BODY>` consists only of `<FIELDS>` and possibly a `<FIELD-FOOTER>` (i.e. a boundary on a multipart message)
