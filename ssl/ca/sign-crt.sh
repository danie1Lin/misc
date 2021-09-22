openssl ecparam -out server.key -name prime256v1 -genkey
openssl req -new -sha256 -key server.key -out server.csr
openssl x509 -req -in server.csr -CA  ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 -sha256

openssl x509 -in server.crt -text -noout
