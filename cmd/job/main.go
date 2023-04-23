package main

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"availability/pkg/data/collections"
	"availability/pkg/data/fakes"
	"availability/pkg/data/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	pingTimeoutSecs         int = 120
	maxResponseDurationSecs int = 10
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
	// TODO validate URL
	siteURL = rawSiteURL

	ctx := context.Background()
	for true {
		run(ctx, siteID, siteURL)
		time.Sleep(time.Duration(pingTimeoutSecs) * time.Second)
	}
}

func run(ctx context.Context, siteID int, siteURL string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tmout := time.Duration(maxResponseDurationSecs+1) * time.Second

	timer := time.AfterFunc(tmout, func() {
		cancel()
		log.Println("ERROR! timeout")
	})
	defer timer.Stop()

	query := new(fakes.ProbeInserter)
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
	return set.Persist(query)
}

func Probe(ctx context.Context, siteID int, siteURL string) *model.Probe {
	p := new(model.Probe)

	done := make(chan struct{})
	defer close(done)

	timer := time.AfterFunc(20*time.Second, func() {
		done <- struct{}{}
	})
	defer timer.Stop()

	select {
	case <-ctx.Done():
		log.Println("\t- Context DONE")
		return nil
	case <-done:
		log.Println("\t- Processing DONE")
		break
	}
	p.SiteID = int32(siteID)
	p.Err = model.HttpErr_HTTPERR_OK
	p.Recorded = timestamppb.New(time.Now())

	return p
}
