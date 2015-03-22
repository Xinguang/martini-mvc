package main

import (
	h "github.com/starboychina/martini-mvc/src/helpers"
	"log"
	"net/http"
)

func main() {
	m := h.Initialization()

	// HTTP
	go func() {
		if err := http.ListenAndServe(":80", m); err != nil {
			log.Fatal(err)
		}
	}()

	// HTTPS
	// To generate a development cert and key, run the following from your *nix terminal:
	// go run $GOROOT/src/crypto/tls/generate_cert.go --host="localhost"
	if err := http.ListenAndServeTLS(":443", "config/cert/cert.pem", "config/cert/key.pem", m); err != nil {
		log.Fatal(err)
	}
}
