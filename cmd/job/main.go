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

func main() {
	rawSiteID := os.Getenv("AVBL_SITE_ID")
	rawSiteURL := os.Getenv("AVBL_SITE_URL")
	_ = os.Getenv("AVBL_PREVIOUSLY_DOWN")

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

	ctx := context.Background()
	for i := 0; i < maxJobRuns; i++ {
		err := run(ctx, siteID, siteURL)
		if err != nil {
			log.Printf("ERROR: %v", err)
		}
		time.Sleep(time.Duration(pingTimeoutSecs) * time.Second)
	}
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
	log.Printf("TODO: gonna save the outage info if applicable: %d", probeId)
	return nil
}
