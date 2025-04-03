## Setup Instructions

1. **Generate Certificates:**  
   Use OpenSSL (or your preferred tool) to generate the CA, server, and client certificates. Place these files under the **certs/** directory.

# Create a CA
```
openssl req -x509 -new -nodes -keyout ca.key -out ca.crt -days 365 -subj "/CN=my-ca"
```
# Generate Server Certificate
```
openssl req -new -newkey rsa:4096 -nodes -keyout server.key -out server.csr -subj "/CN=localhost"
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365
```
# Generate Client Certificate
```
openssl req -new -newkey rsa:4096 -nodes -keyout client.key -out client.csr -subj "/CN=client"
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365
```