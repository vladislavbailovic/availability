package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"availability/pkg/data"
	"availability/pkg/env"
	"availability/pkg/server"
)

var sinceRangeMax time.Duration = 7 * 24 * time.Hour

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

	http.HandleFunc("/daily/", server.Handle(server.WithExpectedHeaders(
		hdr, server.WithExpectedMethod(http.MethodGet, daily))))

	http.HandleFunc("/weekly/", server.Handle(server.WithExpectedHeaders(
		hdr, server.WithExpectedMethod(http.MethodGet, weekly))))

	http.HandleFunc("/monthly/", server.Handle(server.WithExpectedHeaders(
		hdr, server.WithExpectedMethod(http.MethodGet, monthly))))
}

func since(w *server.Response, r *http.Request) error {
	now := time.Now()
	start := time.Unix(
		int64(server.ExtractNumberFromPathAt(r, 1)), 0)
	if start.After(now) {
		return errors.New("in future")
	}
	threshold := now.Add(-sinceRangeMax)
	if start.Before(threshold) {
		return errors.New("too early")
	}

	log.Println(start)
	log.Println(start.Unix())

	return errors.New("TODO: implement since")
}

func daily(w *server.Response, r *http.Request) error {
	siteID, err := extractIDFromPath(r)
	if err != nil {
		return err
	}

	log.Println("daily", siteID)

	return errors.New("TODO: implement sourcePeriod")
}

func weekly(w *server.Response, r *http.Request) error {
	siteID, err := extractIDFromPath(r)
	if err != nil {
		return err
	}

	log.Println("weekly", siteID)

	return errors.New("TODO: implement sourcePeriod")
}

func monthly(w *server.Response, r *http.Request) error {
	siteID, err := extractIDFromPath(r)
	if err != nil {
		return err
	}

	log.Println("monthly", siteID)

	return errors.New("TODO: implement sourcePeriod")
}

func extractIDFromPath(r *http.Request) (data.DataID, error) {
	siteID := data.DataID(server.ExtractNumberFromPathAt(r, 1))
	if !siteID.IsValid() {
		return 0, errors.New("invalid site ID")
	}
	return siteID, nil
}
