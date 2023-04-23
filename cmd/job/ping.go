package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"availability/pkg/data/model"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ping(ctx context.Context, siteID int, siteURL string) (status *model.Probe) {
	status = new(model.Probe)
	status.SiteID = int32(siteID)

	defer func() {
		if err := recover(); err != nil {
			status.Err = model.HttpErr_HTTPERR_INTERNAL
			status.Msg = fmt.Sprintf("panic: %v", err)
		}
	}()

	req, err := http.NewRequest(
		http.MethodHead,
		siteURL,
		nil)
	if err != nil {
		panic(err)
	}
	req.Header = http.Header{
		"User-Agent": {"Wat"},
	}

	client := &http.Client{
		Timeout: time.Duration(maxResponseDurationSecs) * time.Second,
	}
	start := time.Now()
	status.Recorded = timestamppb.New(start)
	resp, err := client.Do(req)
	status.ResponseTime = durationpb.New(time.Now().Sub(start))
	if err != nil {
		panic(err)
	}

	status.Err = model.HttpErr(resp.StatusCode)

	return status
}
