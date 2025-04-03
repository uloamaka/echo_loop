## Setup Instructions

1. **Generate Certificates:**  
   Use OpenSSL (or your preferred tool) to generate the CA, server, and client certificates. Place these files under the **certs/** directory.

# Create a CA (valid for 30 day)
```
openssl req -x509 -new -nodes -keyout ca.key -out ca.crt -days 30 -subj "/CN=my-ca"
```
# Generate Server Certificate (valid for 30 day)
```
openssl req -new -newkey rsa:4096 -nodes -keyout server.key -out server.csr -subj "/CN=localhost"
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 30 -extfile <(printf "subjectAltName=DNS:localhost")
```
# Generate Client Certificate (valid for 30 day)
```
openssl req -new -newkey rsa:4096 -nodes -keyout client.key -out client.csr -subj "/CN=client"
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 30
```
2. **Run the Server:**  
   Open a terminal, navigate to the **rpc_server/** directory, and run:
   ```bash
   go run main.go
   
3. **Run the Client:**  
   Open a terminal, navigate to the **rpc_client/** directory, and run:
   ```bash
   go run main.go