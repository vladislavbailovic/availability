package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"availability/pkg/env"
	"availability/pkg/server"
)

func main() {
	apiPort := env.ApiPortData.WithFallback("3667")
	auth := env.ApiSecretData.Value()

	hdr := http.Header{}
	if auth != "" {
		hdr.Add("x-avbl-auth", auth)
	}
	registerHandlers(hdr)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", apiPort), nil))
}

func registerHandlers(hdr http.Header) {
	http.HandleFunc("/since/", server.Handle(server.WithExpectedHeaders(
		hdr, server.WithExpectedMethod(http.MethodGet, since))))
	http.HandleFunc("/", server.Handle(server.WithExpectedHeaders(
		hdr, server.WithExpectedMethod(http.MethodGet, sourcePeriod))))
}

func since(w *server.Response, r *http.Request) error {
	return errors.New("since")
}

func sourcePeriod(w *server.Response, r *http.Request) error {
	return errors.New("sourcePeriod")
}
