#### Testing with cURL

Before testing add to your `/etc/hosts` the following line: `127.0.0.1       clerio.com` and then test it referenciating your certificate file with the option `--cacert`

    curl  --cacert certificates/certificate.pem https://clerio.com:8443/ -v

