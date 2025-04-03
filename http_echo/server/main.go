package main

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"fmt"
	"io"
	"net/http"
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

	server := &http.Server{
		Addr:    ":8443",
		Handler: mux,
		TLSConfig: LoadServerTLSConfig(),
	}

	fmt.Println("🔐 mTLS Server running on https://localhost:8443")
	err := server.ListenAndServeTLS("../certs/server.crt", "../certs/server.key")
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