# Instructions

#### Generate  private key

    openssl genrsa -out private_key.pem 2048

#### Generate Certificate Singning Request

    openssl req -new -key  private_key.pem  -out csr.pem -config configuration.conf

#### Generate Self Singned Certificate

    openssl x509 -req -in csr.pem -signkey private_key.pem -out  certificate.pem -days 365 

#### View the Generated Certificate (Optional)

    openssl x509 -in certificate.pem -text -noout








