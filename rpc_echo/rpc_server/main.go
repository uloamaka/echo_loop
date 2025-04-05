package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

type RPCRequest struct {
	Method string `json:"method"`
	Params interface{} `json:"params,omitempty"`
}

type RPCResponse struct {
	Result interface{} `json:"result,omitempty"`
	Error string `json:"error,omitempty"`
}

func main() {
	// Load TLS configuration
	tlsConfig := LoadServerTLSConfig()

    port := os.Getenv("RPC_PORT")
    if port == "" {
        port = "8443"
    }
	address := fmt.Sprintf(":%s", port)

	ln, err := tls.Listen("tcp", address, tlsConfig)
	if err != nil {
		log.Fatalf("Falled to listen: %v", err)
	}
	log.Printf("🔐 mTLS Server running on port:%s (mTLS enabled)", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	var req RPCRequest
	if err := decoder.Decode(&req); err != nil {
		log.Printf("Decode error: %v", err)
		return
	}

	log.Printf("Received RPC request: %+v", req)

	var res RPCResponse
	switch req.Method {
	case "Hello":
		res.Result = "Hello, Secure RPC Client!"
	default:
		res.Error = "Unknown method"
	}

	if err := encoder.Encode(res); err != nil {
		log.Printf("Encode error: %v", err)
	}
}

func LoadServerTLSConfig() *tls.Config {
	// Load CA certificate to verify client certificates.
	caCert, err := os.ReadFile("../certs/ca.crt")
	if err != nil {
		log.Fatalf("Failed to load CA certificate: %v", err)
	}

	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatalf("Failed to append CA certificate")
	}

	// Load server certificate and private key.
	cert, err := tls.LoadX509KeyPair("../certs/server.crt", "../certs/server.key")
	if err != nil {
		log.Fatalf("Failed to load server certificate: %v", err)
	}

	return &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
		Certificates: []tls.Certificate{cert},
		MinVersion: tls.VersionTLS13,
	}
}
