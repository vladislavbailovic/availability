package main

import (
	"log"
	"time"

	"availability/pkg/data/collections"
	"availability/pkg/data/fakes"
	"availability/pkg/data/model"
)

func main() {
	query := &fakes.IncidentReportCollector{
		Reports: []fakes.Report{
			fakes.Report{
				ID:      1312,
				URL:     "wat",
				Started: "2013-12-01 16:10:00",
				Err:     model.HttpErr_HTTPERR_NOT_FOUND,
				Msg:     "GTFO",
				Ended:   "2013-12-01 16:11:10",
			},
			fakes.Report{
				ID:      1312,
				URL:     "wat",
				Started: "2013-12-02 16:10:00",
				Err:     model.HttpErr_HTTPERR_NOT_FOUND,
				Msg:     "GTFO",
				Ended:   "2013-12-03 16:10:00",
			},
		},
	}
	r, err := collections.GetIncidentReportsFor(query, 1312, time.Duration(7*24)*time.Hour)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)
}
