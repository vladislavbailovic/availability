package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"availability/pkg/data"
	"availability/pkg/data/collections"
	"availability/pkg/data/model"
	"availability/pkg/data/sql"

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
	http.HandleFunc("/activate/", handle(
		WithExpectedHeaders(hdr, WithExpectedMethod(http.MethodPut, activate))))
	http.HandleFunc("/deactivate/", handle(
		WithExpectedHeaders(hdr, WithExpectedMethod(http.MethodPut, deactivate))))
	http.HandleFunc("/add/", handle(
		WithExpectedHeaders(hdr, WithExpectedMethod(http.MethodPost, addNew))))
}

type handler func(http.ResponseWriter, *http.Request) error

func handle(f handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Printf("ERROR [%s %s %s]: %v",
				r.RemoteAddr, r.Method, r.URL.Path, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("DEBUG: [%s %s %s] OK",
			r.RemoteAddr, r.Method, r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}
}

func WithExpectedMethod(method string, f handler) handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		if r == nil {
			return errors.New("invalid request")
		}
		if r.Method != method {
			return fmt.Errorf("unsupported request type: %q, expected %q",
				r.Method, method)
		}
		return f(w, r)
	}
}

func WithExpectedHeaders(hdr http.Header, f handler) handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		if r == nil {
			return errors.New("invalid request")
		}
		for key, wants := range hdr {
			gots := r.Header.Values(key)
			if len(gots) == 0 {
				return fmt.Errorf("missing required header %q", key)
			}
			for idx, want := range wants {
				got := gots[idx]
				if want != got {
					return fmt.Errorf("invalid header %q(%d): %q", key, idx, got)
				}
			}
		}
		return f(w, r)
	}
}

func ExtractNumberFromPathAt(r *http.Request, at int) int {
	splits := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(splits) <= at {
		return 0
	}
	id, err := strconv.Atoi(splits[at])
	if err != nil {
		return 0
	}
	return id
}

func extractIDFromPath(r *http.Request) (data.DataID, error) {
	siteID := data.DataID(ExtractNumberFromPathAt(r, 1))
	if !siteID.IsValid() {
		return 0, errors.New("invalid site ID")
	}
	return siteID, nil
}

func activate(w http.ResponseWriter, r *http.Request) error {
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

func deactivate(w http.ResponseWriter, r *http.Request) error {
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

func addNew(w http.ResponseWriter, r *http.Request) error {
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
