# Echo Loop: mTLS Authentication for RPC and HTTP Connections

`echo_loop` is a project demonstrating mutual TLS (mTLS) authentication across two distinct communication protocols: Remote Procedure Call (RPC) and Hypertext Transfer Protocol (HTTP). This repository serves as the main hub, containing two sub-repositories:

- **`rpc_echo`**: Implements an RPC-based server and client with mTLS authentication.
- **`http_echo`**: Implements an HTTP-based server and client with mTLS authentication.

Both sub-repositories showcase secure communication using certificates for mutual authentication, ensuring that only trusted clients and servers can interact.

## Purpose

The `echo_loop` project aims to:
- Demonstrate how mTLS authentication secures communication between clients and servers.
- Provide reusable examples for RPC and HTTP implementations using Go.
- Allow users to run these services locally or via Docker, with clear setup instructions.

## mTLS Authentication Flow

The following diagram illustrates the mTLS authentication process used in both `rpc_echo` and `http_echo`:

![mTLS Authentication Flow](./rpc_echo/images/mtls_flow.png)

This flow shows how the client and server exchange and verify certificates to establish a secure connection:
1. Client sends its certificate to the server.
2. Server sends its certificate to the client.
3. Both verify the received certificates against a trusted CA.
4. A secure TLS channel is established for communication.

## Getting Started

To explore and run the examples, navigate to the sub-repositories:

### RPC Connections with mTLS
- **Directory**: `rpc_echo/`
- **Description**: Implements an RPC server and client using mTLS for secure communication.
- **Instructions**: See [`rpc_echo/README.md`](./rpc_echo/README.md) for detailed setup steps, including certificate generation, running locally, and using Docker.

### HTTP Connections with mTLS
- **Directory**: `http_echo/`
- **Description**: Implements an HTTP server and client using mTLS for secure communication.
- **Instructions**: See [`http_echo/README.md`](./http_echo/README.md) for detailed setup steps, including certificate generation, running locally, and using Docker.

## Prerequisites

- **Go**: Required for running the applications locally (install via `sudo apt install golang` on Ubuntu or download from [golang.org](https://golang.org/dl/)).
- **Docker**: Required for running the applications in containers (install via [Docker Desktop](https://www.docker.com/products/docker-desktop) or your package manager).
- **OpenSSL**: Required for generating certificates (install via `sudo apt install openssl` on Ubuntu).

## Setup Overview

1. **Clone the Repository**:
   ```
   git clone https://github.com/uloamaka/echo_loop.git
   cd echo_loop

## Project Aim

This project is a contribution to pgwatch, aiming to implement gRPC support with mutual Transport Layer Security (mTLS) authentication. The goal is to enhance pgwatch's capabilities for secure and efficient communication between its components, proposed for the Google Summer of Code (gsoc) 2025.

CC pgwatch gRPC FOR 2025
