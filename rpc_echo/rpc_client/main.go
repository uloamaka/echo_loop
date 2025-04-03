package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type RPCRequest struct {
	Method string	`json:"method"`
	Params interface{} `json:"params, omitempty"`
}

type RPCResponse struct {
	Result interface{} `json:"result,omitempty"`
	Error string	`json:"error,omitempty"`
}

func main() {
	tlsConfig := LoadClientTLSConfig()

	// Send an RPS request every 30 seconds.
	for {
		conn, err := tls.Dial("tcp", "localhost:8443", tlsConfig)
		if err != nil {
			log.Printf("Dial error: %v", err)
		} else {
			encoder := json.NewEncoder(conn)
			decoder := json.NewDecoder(conn)

			req := RPCRequest{
				Method: "Hello",
			}

			if err := encoder.Encode(&req); err != nil {
				log.Printf("Dial error: %v", err)
			}

			var res RPCResponse
			if err := decoder.Decode(&res); err != nil {
				log.Printf("DEcode error: %v", err)
			} else {
				fmt.Printf("Rpc Response: %+v\n", res)
			}
			conn.Close()
		}
		time.Sleep(30 * time.Second)
	}
}

func LoadClientTLSConfig() *tls.Config {
	caCert, err := os.ReadFile("../certs/ca.crt")
	if err != nil {
		log.Fatalf("Failed to load CA certificate: %v", err)
	}

	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatalf("Failed to append CA certificate")
	}

	cert, err := tls.LoadX509KeyPair("../certs/client.crt", "../certs/client.key")
	if err != nil {
		log.Fatalf("Failed to load client certificate: %v", err)
	}

	return &tls.Config{
		RootCAs: caCertPool,
		Certificates: []tls.Certificate{cert},
	}
}