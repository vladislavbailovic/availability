package main

import (
	"log"
	"os"
	"time"

	"availability/pkg/data"
	"availability/pkg/data/collections"
	"availability/pkg/data/fakes"
	"availability/pkg/data/model"
	"availability/pkg/graph"
)

func main() {
	dailyResponseTimesPlot("tmp/test.svg")
}

func dailyResponseTimesPlot(outfile string) {
	query := &fakes.ProbeCollector{
		Probes: []fakes.Probe{
			fakes.Probe{
				SiteID:       1312,
				Recorded:     "2013-12-01 16:10:00",
				ResponseTime: 161,
				Err:          200,
				Msg:          "",
			},
			fakes.Probe{
				SiteID:       1312,
				Recorded:     "2013-12-01 16:12:00",
				ResponseTime: 13,
				Err:          200,
				Msg:          "",
			},
			fakes.Probe{
				SiteID:       1312,
				Recorded:     "2013-12-01 16:14:00",
				ResponseTime: 12,
				Err:          200,
				Msg:          "",
			},
			fakes.Probe{
				SiteID:       1312,
				Recorded:     "2013-12-01 16:16:00",
				ResponseTime: 161,
				Err:          200,
				Msg:          "",
			},
			fakes.Probe{
				SiteID:       1312,
				Recorded:     "2013-12-01 16:18:00",
				ResponseTime: 16,
				Err:          200,
				Msg:          "",
			},
			fakes.Probe{
				SiteID:       1312,
				Recorded:     "2013-12-01 16:20:00",
				ResponseTime: 161,
				Err:          200,
				Msg:          "",
			},
			fakes.Probe{
				SiteID:       1312,
				Recorded:     "2013-12-01 16:22:00",
				ResponseTime: 80,
				Err:          200,
				Msg:          "",
			},
		},
	}
	now := data.TimestampFromDatetime("2013-12-01 16:00:00").AsTime()
	r, err := collections.GetProbesForWithin(query, 1312, 24*time.Hour)
	if err != nil {
		log.Fatal(err)
	}

	maker := responseTimesPlotMaker{
		Meta: graph.Meta{
			Start:      now,
			End:        now.Add(time.Hour),
			Resolution: time.Duration(4) * time.Minute,
		},
		probes: r,
	}
	image := maker.Make().Render()
	os.WriteFile(outfile, []byte(image), 0600)

}

func weeklyIncidentsGraph(outfile string) {
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
		Meta: graph.Meta{
			Start:      now.AddDate(0, 0, -3),
			End:        now.AddDate(0, 0, 4),
			Resolution: time.Hour * 24,
		},
		reports: r,
	}

	image := maker.Make().Render()
	os.WriteFile(outfile, []byte(image), 0600)
}
