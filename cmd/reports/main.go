package main

import (
	"log"

	"availability/pkg/data/collections"
	"availability/pkg/data/fakes"
	"availability/pkg/data/model"
)

func main() {
	query := &fakes.IncidentReportScanner{
		Report: fakes.Report{
			ID:      1312,
			URL:     "wat",
			Started: "2013-12-01 16:10:00",
			Err:     model.HttpErr_HTTPERR_NOT_FOUND,
			Msg:     "GTFO",
			Ended:   "2013-12-01 16:11:10",
		},
	}
	r, err := collections.GetIncidentReportFor(query, 1312)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)
}
