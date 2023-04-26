package main

import (
	"context"
	"errors"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	"availability/pkg/data/collections"
	"availability/pkg/data/model"
	"availability/pkg/data/sql"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	pingTimeoutSecs         int = 120
	maxResponseDurationSecs int = 10
	maxJobRuns              int = 5
)

var incident *model.Incident

func main() {
	rawSiteID := os.Getenv("AVBL_SITE_ID")
	rawSiteURL := os.Getenv("AVBL_SITE_URL")
	rawIsDown := os.Getenv("AVBL_PREVIOUSLY_DOWN")

	var siteID int
	if x, err := strconv.Atoi(rawSiteID); err != nil {
		panic(err)
	} else {
		siteID = x
	}
	if siteID <= 0 {
		panic(errors.New("invalid site ID"))
	}

	var siteURL string
	lnk, err := url.Parse(rawSiteURL)
	if err != nil {
		panic(err)
	}
	siteURL = lnk.String()

	var downProbeID int
	if x, err := strconv.Atoi(rawIsDown); err == nil || x > 0 {
		downProbeID = x
	}
	log.Printf("Initiating probe job for %d: %s", siteID, siteURL)

	if downProbeID > 0 {
		log.Printf("Previous down probe: %d (%s)", downProbeID, rawIsDown)
		incident = getLatestIncident(siteID)
		if incident != nil {
			log.Printf("\t- Incident info: down probe %d, up probe %d", incident.DownProbeID, incident.UpProbeID)
		} else {
			log.Printf("\t- WARNING: we were supposed to load previous incident but that didn't happen")
		}
	} else {
		log.Println("Site was apparently up")
	}

	ctx := context.Background()
	for i := 0; i < maxJobRuns; i++ {
		log.Printf("Initiating probe cycle %d", i+1)
		err := run(ctx, siteID, siteURL)
		if err != nil {
			log.Printf("ERROR: %v", err)
		}
		time.Sleep(time.Duration(pingTimeoutSecs) * time.Second)
	}
	log.Println("Done probing, recycling")
}

func run(ctx context.Context, siteID int, siteURL string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tmout := time.Duration(maxResponseDurationSecs+1) * time.Second
	sleepTmout := time.Duration(maxResponseDurationSecs) * time.Second

	timer := time.AfterFunc(tmout, cancel)
	defer timer.Stop()

	query := new(sql.ProbeInserter)
	if err := query.Connect(); err != nil {
		log.Println("unable to connect:", err)
		return err
	}
	defer query.Disconnect()

	set := new(collections.ProbeSet)
	confirmation := false

	p := ping(ctx, siteID, siteURL)
	log.Printf("\t- Probe: %d, %dms",
		p.Err, p.ResponseTime.AsDuration().Milliseconds())
	if p != nil {
		set.Add(p)
		confirmation = p.IsDown()
	} else {
		p := new(model.Probe)
		p.SiteID = int32(siteID)
		p.Recorded = timestamppb.New(time.Now())
		p.Err = model.HttpErr_HTTPERR_INTERNAL
		p.Msg = "Timeout"
		set.Add(p)
		confirmation = true
	}

	if confirmation {
		// TODO better confirmation delay
		timer.Reset(tmout)
		time.Sleep(sleepTmout)
		timer.Reset(tmout)
		p := ping(ctx, siteID, siteURL)
		log.Printf("\t- Confirmation: %d, %dms",
			p.Err, p.ResponseTime.AsDuration().Milliseconds())
		if p != nil {
			set.Add(p)
		} else {
			p := new(model.Probe)
			p.SiteID = int32(siteID)
			p.Recorded = timestamppb.New(time.Now())
			p.Err = model.HttpErr_HTTPERR_INTERNAL
			p.Msg = "Timeout"
			set.Add(p)
		}
	}

	timer.Stop()

	probeId, err := set.Persist(query)
	if err != nil {
		return err
	}

	if incident != nil && !set.IsDown() && probeId != 0 {
		log.Println("We have ongoing incident and we're back up: closing off incident")
		incident.UpProbeID = int32(probeId)
		query := new(sql.IncidentUpdater)
		if err := collections.CloseOffIncident(query, incident); err != nil {
			return err
		}
		incident = nil
	} else if incident == nil && set.IsDown() && probeId != 0 {
		log.Println("No outgoing incident and we just went down: starting and persisting new incident")
		incident = new(model.Incident)
		incident.SiteID = int32(siteID)
		incident.DownProbeID = int32(probeId)
		query := new(sql.IncidentInserter)
		if id, err := collections.CreateNewIncident(query, incident); err != nil {
			return err
		} else {
			incident.IncidentID = int32(id)
		}
	}

	return nil
}

func getLatestIncident(siteID int) *model.Incident {
	query := new(sql.IncidentSelection)
	if err := query.Connect(); err != nil {
		log.Println("unable to connect:", err)
		return nil
	}
	defer query.Disconnect()

	incident, err := collections.GetSiteIncident(query, siteID)
	if err != nil {
		log.Printf("ERROR selecting last incident: %v", err)
		return nil
	}
	return incident
}
