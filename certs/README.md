## How to create a self-assigned ssl key / certificate pair

    openssl genrsa -des3 \
      -out fooserver-ssl.key.orig.pem 2048 # or 1024

    openssl req \
      -subj '/C=US/ST=Utah/L=Provo/CN=www.example.com' \
      -new -key fooserver-ssl.key.orig.pem \
      -out fooserver-ssl.csr

    openssl rsa \
      -in fooserver-ssl.key.orig.pem \
      -out fooserver-ssl.key.pem

    openssl x509 -req \
      -days 365 \
      -in fooserver-ssl.csr \
      -signkey fooserver-ssl.key.pem \
      -out fooserver-ssl.crt.pem

Resources:

  * http://www.akadia.com/services/ssh_test_certificate.html
  * https://devcenter.heroku.com/articles/ssl-certificate-self
  * http://stackoverflow.com/questions/9969250/create-pem-file-for-google-manage-domains?rq=1
  * https://developers.google.com/gdata/docs/auth/authsub#Registered
