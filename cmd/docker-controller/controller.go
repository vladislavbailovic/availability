package main

import (
	"fmt"
	"net/url"
	"strings"

	"availability/pkg/data/model"
)

type envName uint

const (
	envSiteID envName = iota
	envSiteURL
	envPreviouslyDown

	_envNamesCount
)

func (x envName) String() string {
	switch x {
	case envSiteID:
		return "AVBL_SITE_ID"
	case envSiteURL:
		return "AVBL_SITE_URL"
	case envPreviouslyDown:
		return "AVBL_PREVIOUSLY_DOWN"
	default:
		panic("unknown env var")
	}
}

func getJobName(siteID int32, siteURL string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "ping-%d", siteID)

	lnk, err := url.Parse(siteURL)
	if err == nil {
		b.WriteString("-")
		b.WriteString(strings.Replace(lnk.Host, ":", "", -1))
	}

	return b.String()
}

func getJobEnv(task *model.Task) []string {
	down := 0
	if task.WasPreviouslyDown() {
		down = 1
	}
	return []string{
		fmt.Sprintf("%s=%d", envSiteID.String(), task.Source.SiteID),
		fmt.Sprintf("%s=%s", envSiteURL.String(), task.Source.URL),
		fmt.Sprintf("%s=%d", envPreviouslyDown.String(), down),
	}
}
