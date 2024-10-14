## Context

The goal of this project is to generate **Self Signed SSL Certificate** that refers to an "unpublished" `Domain Name`, and then use the generated cerfiticate in a tiny Golang REST webservice and proceed with *curl* requests without ignoring `TLS verification` . 


## Requiremtns

* Golang 1.21+
* openssl


## Running it 

Spin up the REST webservice contained in `app1` folder, running:

    go run main.go    

### Testing with cURL

Before testing, add to your `/etc/hosts` the following entry: `127.0.0.1       clerio.com` and then test it referenciating your certificate file with the option `--cacert`

    curl  --cacert certificates/certificate.pem https://clerio.com:8443/ -v

We use the `-v` option in the curl command to turn it **verbosed** in order to  have a detailed output for analysis 

## Understating the `output`

### 1. Connection Attempt:

    *   Trying 127.0.0.1:8443...
    * Connected to clerio.com (127.0.0.1) port 8443 (#0)

* **Trying 127.0.0.1:8443...**: `curl` attempts to connect to the IP address `127.0.0.1`(localhost) on port `8443` (commonly used for HTTPS).
* **Connected to clerio.com (127.0.0.1)**: The connection was successful. The server at `127.0.0.1` is associated with the domain `clerio.com`

### 2. ALPN Protocol Negotiation:
bash

    * ALPN, offering h2
    * ALPN, offering http/1.1

*  **ALPN (Application-Layer Protocol Negotiation)**: This is a mechanism that allows the client and server to agree on which protocol to use. Here, `curl` offers two protocols: `h2` (HTTP/2) and `http/1.1` (HTTP/1.1).

### 3. Certificate and SSL/TLS Handshake:
bash

    *  CAfile: certificate.pem
    *  CApath: /etc/ssl/certs
    * TLSv1.0 (OUT), TLS header, Certificate Status (22):
    * TLSv1.3 (OUT), TLS handshake, Client hello (1):
    * TLSv1.2 (IN), TLS header, Certificate Status (22):
    * TLSv1.3 (IN), TLS handshake, Server hello (2):

*  **CAfile and CApath**: This shows that `curl` is trying to verify the server’s certificate using the provided CA file (`certificate.pem`) and the system’s CA bundle located in `/etc/ssl/certs`.
*  **TLS Handshake**:
   *  The client starts the TLS handshake by sending a "Client hello" to negotiate security parameters (supported versions, ciphers, etc.).
   *  The server responds with a "Server hello" to agree on the TLS version and cipher suite. This is the start of establishing a secure session.

`bash`

    TLSv1.3 (IN), TLS handshake, Encrypted Extensions (8):
    TLSv1.3 (IN), TLS handshake, Certificate (11):
    TLSv1.3 (IN), TLS handshake, CERT verify (15):
    TLSv1.3 (IN), TLS handshake, Finished (20):
    
* **Encrypted Extensions**: Server sends additional information as part of the TLS handshake (e.g., protocol settings).
* **Certificate**: The server sends its certificate to the client.
* **CERT verify**: The server verifies its own certificate and confirms it to the client.
* **Finished**: This concludes the TLS handshake, confirming the established secure connection.
### 4. TLS Version and Cipher Suite:

    * SSL connection using TLSv1.3 / TLS_AES_128_GCM_SHA256

* **TLSv1.3**: The connection uses **TLS 1.3**, which is a secure version of the TLS protocol.
* **TLS_AES_128_GCM_SHA256**: The cipher suite used for encryption, which provides confidentiality, integrity, and authentication.
### 5. ALPN and HTTP/2 Negotiation:


    
    * ALPN, server accepted to use h2

The server accepts **HTTP/2** (`h2`) as the protocol to use for this session, as negotiated via ALPN.
### 6. Server Certificate Information:

bash

    * Server certificate:
    *  subject: C=MZ; ST=Maputo; L=Maputo; O=Mozambique Clerio Software; OU=IT Department; CN=clerio.com
    *  start date: Oct  2 09:26:41 2024 GMT
    *  expire date: Oct  2 09:26:41 2025 GMT
    *  common name: clerio.com (matched)
    *  issuer: C=MZ; ST=Maputo; L=Maputo; O=Mozambique Clerio Software; OU=IT Department; CN=clerio.com
    *  SSL certificate verify ok.

* **Server certificate**: This shows the details of the  
    * **server's SSL certificate**.
    * **Subject**: The certificate is issued to `clerio.com`, with organizational information (Country, State, Location, etc.).
    * **Validity**: The certificate is valid from Oct 2, 2024, to Oct 2, 2025.
    * **Common Name (CN)**: The Common Name `clerio.com` matches the domain you’re accessing.
    * **Issuer**: The certificate is self-signed because the issuer information is the same as the subject.
    SSL certificate verify ok: The certificate is successfully verified.

### 7. HTTP/2 and Multiplexing Support:
bash

    * Using HTTP2, server supports multiplexing

* **Using HTTP/2**: The client and server agreed to use HTTP/2, which supports **multiplexing** (multiple streams over a single connection).

### 8. Sending the HTTP GET Request:
bash

    > GET / HTTP/2
    > Host: clerio.com:8443
    > user-agent: curl/7.81.0
    > accept: */*

* **GET / HTTP/2**: The client sends an HTTP `GET` request to the server over HTTP/2 to the root (`/`) of the server.
* **Host**: The request is made to `clerio.com` on port `8443`.
* **user-agent**: The client identifies itself as `curl/7.81.0`.
* **accept**: The client accepts any content type (*/*).

### 9. Server Response:
bash

    < HTTP/2 200 
    < content-type: application/json
    < content-length: 17
    < date: Wed, 02 Oct 2024 10:17:45 GMT

* **HTTP/2 200**: The server responds with HTTP status `200`, indicating a successful request.
* **content-type**: application/json: The server returns JSON data.
* **content-length**: 17: The response body is 17 bytes long.
* **date**: The server timestamp of the response.
### 10. Response Body:
bash

    "{status:alive}"

The server responds with a small JSON message indicating `{status:alive}`.
### 11. Connection Closure:
bash

    * Connection #0 to host clerio.com left intact

The connection remains open (`left intact`) but will eventually be closed when needed.

