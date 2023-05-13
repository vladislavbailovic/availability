package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

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
	// TODO: auth
	http.HandleFunc("/activate/", activate)
	http.HandleFunc("/deactivate/", deactivate)
	http.HandleFunc("/add/", addNew)
}

func activate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		log.Printf("unsupported request type: %v", r.Method)
		sendServerError(w)
		return
	}
	rawID := strings.Replace(r.URL.String(), "/activate/", "", 1)
	siteID, err := strconv.Atoi(rawID)
	if err != nil {
		log.Println(err)
		sendServerError(w)
		return
	}

	query := new(sql.SourceActivator)
	if err := query.Connect(); err != nil {
		log.Println(err)
		sendServerError(w)
		return
	}
	defer query.Disconnect()
	if err := collections.UpdateSource(query, siteID); err != nil {
		log.Println(err)
		sendServerError(w)
		return
	}

	sendAllGood(w)
}

func deactivate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		log.Printf("unsupported request type: %v", r.Method)
		sendServerError(w)
		return
	}
	rawID := strings.Replace(r.URL.String(), "/deactivate/", "", 1)
	siteID, err := strconv.Atoi(rawID)
	if err != nil {
		log.Println(err)
		sendServerError(w)
		return
	}

	query := new(sql.SourceDeactivator)
	if err := query.Connect(); err != nil {
		log.Println(err)
		sendServerError(w)
		return
	}
	defer query.Disconnect()
	if err := collections.UpdateSource(query, siteID); err != nil {
		log.Println(err)
		sendServerError(w)
		return
	}

	sendAllGood(w)
}

func addNew(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("unsupported request type: %v", r.Method)
		sendServerError(w)
		return
	}

	defer r.Body.Close()
	src := new(model.NewSource)
	if err := jsonpb.Unmarshal(r.Body, src); err != nil {
		log.Print(err)
		sendServerError(w)
		return
	}

	query := new(sql.SourceInserter)
	id, err := collections.CreateNewSource(query, src)
	if err != nil {
		log.Println(err)
		sendServerError(w)
		return
	}

	log.Printf("insert new source: %v, ID: %d", src, id)
	sendAllGood(w)
}

func sendAllGood(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func sendServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}
