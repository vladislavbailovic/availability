package main

import (
	"log"
	"os"
	"time"

	"availability/pkg/data"
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
	weekAgo := time.Duration(7*24) * time.Hour
	now := data.TimestampFromDatetime("2013-12-01 00:00:00").AsTime()
	r, err := collections.GetIncidentReportsFor(query, 1312, weekAgo)
	if err != nil {
		log.Fatal(err)
	}

	maker := incidentReportGraphMaker{
		start:      now.AddDate(0, 0, -3),
		end:        now.AddDate(0, 0, 4),
		resolution: time.Hour,
		reports:    r,
	}

	os.WriteFile("tmp/test.svg", []byte(maker.Make().Render()), 0600)
}
