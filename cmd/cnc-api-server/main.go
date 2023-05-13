package main

import (
	"errors"
	"log"
	"net/http"

	"availability/pkg/data"
	"availability/pkg/data/collections"
	"availability/pkg/data/model"
	"availability/pkg/data/sql"
	"availability/pkg/server"

	"github.com/gogo/protobuf/jsonpb"
)

func main() {
	registerHandlers()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func registerHandlers() {
	// TODO: unstub auth header
	hdr := http.Header{
		"x-avbl-auth": []string{"test"},
	}
	http.HandleFunc("/activate/", server.Handle(server.WithExpectedHeaders(
		hdr, server.WithExpectedMethod(http.MethodPut, activate))))
	http.HandleFunc("/deactivate/", server.Handle(server.WithExpectedHeaders(
		hdr, server.WithExpectedMethod(http.MethodPut, deactivate))))
	http.HandleFunc("/add/", server.Handle(server.WithExpectedHeaders(
		hdr, server.WithExpectedMethod(http.MethodPost, addNew))))
}
func extractIDFromPath(r *http.Request) (data.DataID, error) {
	siteID := data.DataID(server.ExtractNumberFromPathAt(r, 1))
	if !siteID.IsValid() {
		return 0, errors.New("invalid site ID")
	}
	return siteID, nil
}

func activate(w *server.Response, r *http.Request) error {
	siteID, err := extractIDFromPath(r)
	if err != nil {
		return err
	}

	query := new(sql.SourceActivator)
	if err := collections.UpdateSource(query, siteID); err != nil {
		return err
	}

	return nil
}

func deactivate(w *server.Response, r *http.Request) error {
	siteID, err := extractIDFromPath(r)
	if err != nil {
		return err
	}

	query := new(sql.SourceDeactivator)
	if err := collections.UpdateSource(query, siteID); err != nil {
		return err
	}

	return nil
}

func addNew(w *server.Response, r *http.Request) error {
	defer r.Body.Close()
	src := new(model.NewSource)
	if err := jsonpb.Unmarshal(r.Body, src); err != nil {
		return err
	}

	query := new(sql.SourceInserter)
	id, err := collections.CreateNewSource(query, src)
	if err != nil {
		return err
	}

	log.Printf("insert new source: %v, ID: %d", src, id)
	return nil
}
