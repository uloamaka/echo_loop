package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	StartServer()
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request from:", r.RemoteAddr)
	fmt.Println("Client authenticated, responding...")
	io.WriteString(w, "Hello, Secure Client!\n")
}

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Error loading .env file: %v, using default port 8443", err)
	}

    port := os.Getenv("HTTP_PORT")
    if port == "" {
        port = "8443"
    }

	server := &http.Server{
		Addr:    ":"+port,
		Handler: mux,
		TLSConfig: LoadServerTLSConfig(),
	}

	fmt.Printf("🔐 mTLS Server running on https://localhost:%s", port)
	err = server.ListenAndServeTLS("../certs/server.crt", "../certs/server.key")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func LoadServerTLSConfig() *tls.Config {
	caCert, err := os.ReadFile("../certs/ca.crt")
	if err != nil {
		fmt.Println("Failed to read CA certificate:", err)
		return nil
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	return &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
		MinVersion: tls.VersionTLS13,
	}
}