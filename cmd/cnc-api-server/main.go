package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"availability/pkg/data/collections"
	"availability/pkg/data/sql"
)

func main() {
	registerHandlers()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func registerHandlers() {
	http.HandleFunc("/activate/", activate)
	http.HandleFunc("/deactivate/", deactivate)
}

func activate(w http.ResponseWriter, r *http.Request) {
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

func sendAllGood(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func sendServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}