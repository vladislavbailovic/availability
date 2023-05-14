package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"availability/pkg/data"
	"availability/pkg/data/collections"
	"availability/pkg/data/model"
	"availability/pkg/data/sql"
	"availability/pkg/env"
	"availability/pkg/server"

	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var sinceRangeMax time.Duration = 7 * 24 * time.Hour

func main() {
	apiPort := env.ApiPortData.WithFallback("3667")
	auth := env.ApiSecretData.Value()

	registerHandlers(server.GetAuthHeader(auth))
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

	report := new(model.PeriodicIncidentReport)
	report.Start = timestamppb.New(start)

	var err error
	limit := 100
	query := new(sql.IncidentReportPeriodCollector)
	report.Incidents, err = collections.GetIncidentReportsWithin(
		query, start, limit)
	if err != nil {
		return err
	}

	w.Header().Add("content-type", "application/json")
	enc := jsonpb.Marshaler{EnumsAsInts: true}
	if err := enc.Marshal(w, report); err != nil {
		return err
	}

	return nil
}

func daily(w *server.Response, r *http.Request) error {
	return handlePeriod(24*time.Hour)(w, r)
}

func weekly(w *server.Response, r *http.Request) error {
	return handlePeriod(7*24*time.Hour)(w, r)
}

func monthly(w *server.Response, r *http.Request) error {
	return handlePeriod(30*24*time.Hour)(w, r)
}

func handlePeriod(period time.Duration) server.Handler {
	return func(w *server.Response, r *http.Request) error {
		reports, err := sourcePeriodFromRequest(r, period)
		if err != nil {
			return err
		}

		w.Header().Add("content-type", "application/json")
		enc := jsonpb.Marshaler{EmitDefaults: true, EnumsAsInts: true}
		if err := enc.Marshal(w, reports); err != nil {
			return err
		}

		return nil
	}
}

func sourcePeriodFromRequest(r *http.Request, period time.Duration) (*model.PeriodicIncidentReport, error) {
	now := time.Now()
	report := new(model.PeriodicIncidentReport)
	report.Start = timestamppb.New(now.Add(-period))
	report.End = timestamppb.New(now)

	siteID := data.DataID(server.ExtractNumberFromPathAt(r, 1))
	if !siteID.IsValid() {
		return report, errors.New("invalid site ID")
	}

	var err error
	query := new(sql.IncidentReportCollector)
	report.Incidents, err = collections.GetIncidentReportsFor(
		query, siteID.ToNumericID(), period)
	if err != nil {
		return report, err
	}

	return report, nil
}
