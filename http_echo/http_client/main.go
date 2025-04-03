package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	StartClient()
}

// StartClient makes a periodic HTTPs GET request to the server.
func StartClient() {
	tlsConfig := LoadClientTLSConfig()

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Printf("Error loading .env file: %v, using default port 8443", err)
	}

    port := os.Getenv("HTTP_PORT")
    if port == "" {
        port = "8443"
    }
	host := os.Getenv("HTTP_HOST")
    if port == "" {
        host = "localhost"
    }

	client := &http.Client{Transport: transport}
	url := fmt.Sprintf("https://%s:%s", host, port)
	
	for {
		res, err := client.Get(url)
		if err != nil {
			log.Println("Request failed:", err)
		} else {
			body, err  := io.ReadAll(res.Body)
			if err != nil {
				log.Println("Failed to read response body:", err)
			} else {
				fmt.Println("Response from server:", string(body))
			}
			res.Body.Close()
		}
		time.Sleep(30 * time.Second) 
	}
	
}

// LoadClientTLSConfig loads the client certificate and CA certificate.
func LoadClientTLSConfig() *tls.Config {
	caCert, err := os.ReadFile("../certs/ca.crt")
	if err != nil {
		fmt.Println("Failed to load CA certificate:", err)
		return nil
	}

	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		fmt.Println("Failed to append CA certificate")
		return nil
	}

	clientCert, err := tls.LoadX509KeyPair("../certs/client.crt", "../certs/client.key")
	if err != nil {
		fmt.Println("Failed to load client certificate:", err)
		return nil
	}

	return &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{clientCert},
	}
}