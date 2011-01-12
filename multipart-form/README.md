multipart-form
====

This example shows the raw conversation between a client and server as well as each part as an individual message.

Please look at the raw files to get a gist of the idea before reading the explanations in this README.

  * 0-full-transcript.http.raw
    * 1-request.http.raw
    * 2-response.http.raw
    * 3-post-multiple-inputs.http.raw

post-multiple-inputs
====

This shows how the client `POST`ed the form.

  * The formt is essentially `<REQUEST><HTTP-HEADERS><FIELDS><FORM-BOUNDARY-END>`

Note: `^M` represents `\r` or `CR` of `\r\n` or `CRLF`.

Note: newlines are represented as new lines rather than `\n` or `LF`

http headers format
----

Example:

    POST / HTTP/1.1^M
    Host: localhost:3000^M
    Keep-Alive: 115^M
    Connection: keep-alive^M
    Content-Type: multipart/form-data; boundary=---------------------------114772229410704779042051621609^M
    Content-Length: 1751^M
    

Note: there is a trailing new line after the last `^M` that may not show up

Explanation:

  * The `<REQUEST>` is unique in that it is the only HTTP directive not preceeded by `<CRLF>`

  * This example shows a multipart form POST
    * `Content-Type: multipart/form-data; boundary=---------------------------114772229410704779042051621609`

  * Each `<HTTP-HEADER>` is **preceeded** `<CRLF>`

  * Either of the following (but not both) can be considered true
    * EITHER `<HTTP-HEADERS>` is followed by `<CRLF>`
    * OR `<CRLF>` preceeds the `<HTTP-BODY>`

field format
----

Note: The first directive in the `<HTTP-BODY>` will be preceeded by `<CRLF>`

Explanation:

  * The format for each `<FIELD>` is essentially `<FULL-BOUNDARY><HEADERS><FIELD-BODY>`
    * Each HTTP line is preceeded by `<CRLF>` (`\r\n` which is show as `^M` and a newline)
  
  * The `<FULL-BOUNDARY>` is `--<boundary>` (two hyphens plus a (probably random) string) which delimits each field.

  * `<HEADERS>` also has a trailing `<CRLF>`, which is not part of `<FIELD-BODY>`


text-field format
----

Example:

    ^M
    -----------------------------114772229410704779042051621609^M
    Content-Disposition: form-data; name="name"^M
    ^M
    AJ ONeal

Explanation:

  * The `<FULL-BOUNDARY>` is `--<boundary>` (two hyphens plus a random string) which delimits each field.
    * in this example `<boundary>` is 57 characters (the maximum is 70)
      * `---------------------------114772229410704779042051621609`
    * in total, this `--<boundary>` is 59 characters
      * (the preceeding `<CRLF>` is required to exist, but is not actually part of the boundary, hence those characters are not counted)
      * `-----------------------------114772229410704779042051621609`

  * The next line is `Content-Disposition: form-data; name="<FIELD-NAME>"`.
    * `Content-Disposition: form-data; name="name"`

  * There is an extra `<CRLF>` between the last (and in this case only) `<HEADER>` between `<HEADERS>` and the `<FIELD-BODY>`
    * `<CRLF><FIELD-BODY>`
    * `\r\nAJ ONeal`

file-field format
----

Example:

    ^M
    -----------------------------114772229410704779042051621609^M
    Content-Disposition: form-data; name="attachments[]"; filename="file2.txt"^M
    Content-Type: text/plain^M
    ^M
    This is file 2
    It has three lines
    And was created on OS X

Explanation:

  * The format is essentially the same as the above, with a notable exception: the `<HEADERS>` includes a `Content-Type` field in the case of a file.

multipart footer
----

Example:

    ^M
    -----------------------------114772229410704779042051621609--^M

Explanation:

  * The only line that ends with `<CRLF>` is the footer
