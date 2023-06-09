package main

import (
	"context"
	"errors"
	"log"
	"net/url"
	"strconv"
	"time"

	"availability/pkg/data/collections"
	"availability/pkg/data/model"
	"availability/pkg/data/sql"
	"availability/pkg/env"
)

const (
	pingTimeoutSecs         int = 120
	maxResponseDurationSecs int = 10
	maxJobRuns              int = 5
)

var activeIncident *model.Incident

func main() {
	rawSiteID := env.SiteID.Expect()
	rawSiteURL := env.SiteURL.Expect()
	rawIsDown := env.PreviouslyDown.Expect()

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
		activeIncident = getLatestIncident(siteID)
		if activeIncident != nil {
			log.Printf("\t- Incident info: down probe %d, up probe %d",
				activeIncident.DownProbeID, activeIncident.UpProbeID)
		} else {
			log.Printf("\t- WARNING: did not load previous incident; we should have")
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
		if i < maxJobRuns-1 {
			// Don't pause if we're about to break,
			// so we don't skip a beat.
			time.Sleep(time.Duration(pingTimeoutSecs) * time.Second)
		} else {
			// Well, actually...
			// *Do* pause, but only by a fraction.
			// This is in order to minimize the changes of skipping
			// a beat when scheduler picks up the job again.
			time.Sleep(time.Duration(pingTimeoutSecs/4) * time.Second)
		}
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
		set.Add(model.NewTimeoutProbe(siteID))
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
			set.Add(model.NewTimeoutProbe(siteID))
		}
	}

	timer.Stop()

	probeId, err := set.Persist(query)
	if err != nil {
		return err
	}

	if activeIncident != nil && !set.IsDown() && probeId != 0 {
		log.Println("We have ongoing incident and we're back up: closing off incident")
		activeIncident.Close(probeId.ToNumericID())
		query := new(sql.IncidentUpdater)
		defer query.Disconnect()
		if err := collections.CloseOffIncident(query, activeIncident); err != nil {
			return err
		}
		activeIncident = nil
	} else if activeIncident == nil && set.IsDown() && probeId != 0 {
		log.Println("No ongoing incident and site went down: " +
			"starting and persisting new incident")
		activeIncident = model.NewIncident(siteID, probeId.ToNumericID())
		query := new(sql.IncidentInserter)
		defer query.Disconnect()
		if id, err := collections.CreateNewIncident(query, activeIncident); err != nil {
			return err
		} else {
			activeIncident.IncidentID = id.ToItemID()
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
