openssl ecparam -out ca.key -name prime256v1 -genkey
openssl req -new -sha256 -key ca.key -out ca.csr
openssl x509 -req -sha256 -days 365 -in ca.csr -signkey ca.key -out ca.crt

