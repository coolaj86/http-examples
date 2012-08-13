/*jshint strict:true node:true es5:true onevar:true laxcomma:true laxbreak:true eqeqeq:true immed:true latedef:true*/
(function () {
  "use strict";

  // See: http://nodejs.org/api/https.html
  var https = require('https')
    , crypto = require('crypto')
    , fs = require('fs')
    , path = require('path')
    , certPath = path.join(__dirname, '..', 'certs')
    // needs passphrase
    , privateKey = fs.readFileSync(path.join(certPath, 'fooserver-ssl.key.orig.pem').toString())
    // doesn't need passphrase
    //, privateKey = fs.readFileSync(path.join(certPath, 'fooserver-ssl.key.pem').toString())
    , certificate = fs.readFileSync(path.join(certPath, 'fooserver-ssl.crt.pem').toString())
    , credentials = {
          key: privateKey
        , cert: certificate
        // Only needed with the key that needs a passphrase
        , passphrase: "HurpDurpHuzzah"
      }
    , server
    ;

  server = https.createServer(credentials, function (req, res) {
      var data = ''
        ;

      console.log(req.connection.remoteAddress + ':' + req.socket.remotePort);
      console.log(req.method + ' ' + req.url);
      req.on('data', function (chunk) {
        data += chunk.toString('utf8');
      });
      req.on('end', function () {
        console.log(data.toString('utf8'));
        res.writeHead(200, {'Content-Type': 'text/plain'});
        res.end('Hello World\n');
      });
  });
  server.listen(8443, function () {
    console.log('Listening on https://'
      + server.address().address
      + ':' + server.address().port
    );
  });
}());
