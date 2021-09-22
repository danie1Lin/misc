ALGORITHM=prime256v1

make_all: server_signed.crt
%_signed.crt: $(CA).crt %.csr
	openssl x509 -req -sha256 -days 365 -in $*.csr -signkey $*.key -out $*_signed.crt -CA $(CA).crt -CAkey $(CA).key -CAcreateserial

.PRECIOUS: %.crt
%.crt: %.csr
	openssl x509 -req -sha256 -days 365 -in $*.csr -signkey $*.key -out $*.crt

.PRECIOUS: %.csr
%.csr: %.key
	openssl req -new -sha256 -key $*.key -out $*.csr

.PRECIOUS: %.key
%.key:
	openssl ecparam -out $*.key -name $(ALGORITHM) -genkey

.PHONY: info_%
info_%:
	openssl x509 -in $*.crt -text -noout

verify:
	

clean_%:
	rm $*.*
