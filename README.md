HTTP Examples
====

I've included some raw http transcriptions to help myself and others understand HTTP by observation in addition to scrutiny of the obfuscated RFCs.

Each example contains a README of its own.

Notes:

  * I interpret `CRLF` as `\r\n` which is represented as follows:
    * `^M` represents `\r` or `CR` (as seen in vim, tthe supreme text editor)
    * `newline`s are represented as new lines rather than `\n` or `LF`

The basic format of an HTTP request is `<REQUEST><HEADERS><BODY>`

HTTP format in a nutshell
====

http request
---

Example:

    POST / HTTP/1.1

Explanation:

  * Essentially `<REQUEST>` is `<HTTP-VERB> <RESOURCE> <PROTOCOL>/<VERSION>`
    * Clients should use a single space between tokens
    * Servers should discard extra whitespace

  * `<REQUEST>` details the which resource should be retrieved (in this case, the document root - '/')

  * `<REQUEST>` is unique in that it does **not** begin with `<CRLF>`
    * like all of the other headers, `<REQUEST>` does not end with `<CRLF>` 

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

  * This example shows a multipart form POST
    * `Content-Type: multipart/form-data; boundary=---------------------------114772229410704779042051621609`

  * Each `<HTTP-HEADER>` is **preceeded** by `<CRLF>`

  * The entire `<HTTP-HEADERS>` block is followed by a single `<CRLF>`

http body format
---

Assuming that you consider your `<HEADERS>` to be followed by `<CRLF>`, the `<BODY>` consists only of `<FIELDS>` and possibly a `<FIELD-FOOTER>` (i.e. a boundary on a multipart message)
