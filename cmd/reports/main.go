package main

import (
	"log"
	"os"
	"time"

	"availability/cmd/reports/incidents"
	"availability/cmd/reports/probes"
	"availability/pkg/data/collections"
	"availability/pkg/data/sql"
	"availability/pkg/graph"
)

func main() {
	dailyResponseTimesPlot("tmp/probes.svg")
	weeklyIncidentsGraph("tmp/incidents.svg")
}

func dailyResponseTimesPlot(outfile string) {
	query := new(sql.ProbeCollector)
	now := time.Now().Truncate(time.Hour)
	r, err := collections.GetProbesForWithin(query, 161, 24*time.Hour)
	if err != nil {
		log.Fatal(err)
	}

	maker := probes.ResponseTimesPlot{
		Meta: graph.Meta{
			Start:      now.Add(-1 * time.Hour),
			End:        now,
			Resolution: time.Duration(4) * time.Minute,
		},
		Probes: r,
	}
	image := maker.Make().Render()
	os.WriteFile(outfile, []byte(image), 0666)

}

func weeklyIncidentsGraph(outfile string) {
	weekAgo := time.Duration(7*24) * time.Hour
	now := time.Now()
	query := new(sql.IncidentReportCollector)
	r, err := collections.GetIncidentReportsFor(query, 1312, weekAgo)
	if err != nil {
		log.Fatal(err)
	}

	maker := incidents.ReportGraph{
		Meta: graph.Meta{
			Start:      now.AddDate(0, 0, -7),
			End:        now,
			Resolution: time.Hour * 24,
		},
		Reports: r,
	}

	image := maker.Make().Render()
	os.WriteFile(outfile, []byte(image), 0666)
}
