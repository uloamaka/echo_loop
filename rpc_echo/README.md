# Setup Instructions for rpc_server and rpc_client
These instructions allow you to run the `rpc_server` and `rpc_client` either directly on your local machine (localhost) or using Docker (locally built or pulled from Docker Hub).

## Project Structure
  /mnt/c/Users/godsg/GSoC_2025/echo_loop/rpc_echo/
  ├── certs/
  │   ├── ca.crt
  │   ├── ca.key
  │   ├── server.crt
  │   ├── server.key
  │   ├── client.crt
  │   └── client.key
  ├── rpc_server/
  │   └── main.go
  ├── rpc_client/
  │   └── main.go
  ├── .env
  ├── go.mod
  └── go.sum


## 1. Generate Certificates

Use OpenSSL to generate the CA, server, and client certificates, and place them in the `certs/` directory 

```bash
mkdir -p certs
cd certs
```

# Create a CA for Localhost (valid for 30 days)
openssl req -x509 -new -nodes -keyout ca.key -out ca.crt -days 30 -subj "/CN=my-ca"

# Generate Server Certificate (valid for 30 days with SAN for localhost)
openssl req -new -newkey rsa:4096 -nodes -keyout server.key -out server.csr -subj "/CN=localhost"
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 30 -extfile <(printf "subjectAltName=DNS:localhost")
rm server.csr 

# Generate Client Certificate (valid for 30 days)
openssl req -new -newkey rsa:4096 -nodes -keyout client.key -out client.csr -subj "/CN=client"
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 30
rm client.csr  # Clean up CSR

## Run the server and client directly on your machine without Docker.

**Build and run Server:**
Navigate to the `rpc_server/` directory:
```
go build -o rpc_server main.go
./rpc_server
```
Alternatively:
```
go run main.go
```
**Build and run Client:**
Navigate to the `rpc_client/` directory:

```
go build -o rpc_client main.go
./rpc_client
```
Alternatively:
```
go run main.go
```
## Running with Docker:
**create a network:**
```
docker network create rpc_network
```
**generate certicates:**
# Create a CA for Localhost (valid for 30 days)
openssl req -x509 -new -nodes -keyout ca.key -out ca.crt -days 30 -subj "/CN=my-ca"

# Generate Server Certificate (valid for 30 days with SAN for rpc_network)
openssl req -new -newkey rsa:4096 -nodes -keyout server.key -out server.csr -subj "/CN=localhost"
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 30 -extfile <(printf "subjectAltName=DNS:localhost")
rm server.csr 

# Generate Client Certificate (valid for 30 days)
openssl req -new -newkey rsa:4096 -nodes -keyout client.key -out client.csr -subj "/CN=client"
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 30
rm client.csr  # Clean up CSR

**Pull images from Docker hub:**
```
docker pull your_username/rpc_server:latest
docker pull your_username/rpc_client:latest
```
**Run the Server:**
```
docker run --name rpc_server --network rpc_network \
  --env-file $(pwd)/.env \
  -v $(pwd)/../certs:/certs \
  -e DOCKER_ENV=true \
  your_username/rpc_server:latest
```
**Run the Client:**
docker run --name rpc_client --network rpc_network \
  --env-file $(pwd)/.env \
  -v $(pwd)/../certs:/certs \
  -e DOCKER_ENV=true \
  your_username/rpc_client:latest