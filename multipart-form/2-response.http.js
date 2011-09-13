(function () {
  "use strict";

  var rawBody
    , headers
    , msg
    ;

  rawBody = 'Hello World';

  headers = '' + 
    'HTTP/1.1 200 OK\r\n' +
    'X-Powered-By: Express\r\n' +
    'Content-Type: text/html; charset=utf-8\r\n' +
    'Content-Length: ' + rawBody.length + '\r\n' +
    'Date: Wed, 12 Jan 2011 19:22:51 GMT\r\n' +
    'X-Response-Time: 0ms\r\n' +
    'Connection: keep-alive\r\n' +
    '\r\n'
    ;

  msg = encodeURIComponent(headers + rawBody);
  console.log(msg);

}());
