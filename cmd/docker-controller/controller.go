package main

import (
	"fmt"
	"net/url"
	"strings"
)

type envName uint

const (
	envSiteID envName = iota
	envSiteURL

	_envNamesCount
)

func (x envName) String() string {
	switch x {
	case envSiteID:
		return "AVBL_SITE_ID"
	case envSiteURL:
		return "AVBL_SITE_URL"
	default:
		panic("unknown env var")
	}
}

func getJobName(siteID int, siteURL string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "ping-%d", siteID)

	lnk, err := url.Parse(siteURL)
	if err == nil {
		b.WriteString("-")
		b.WriteString(strings.Replace(lnk.Host, ":", "", -1))
	}

	return b.String()
}

func getJobEnv(siteID int, siteURL string) []string {
	return []string{
		fmt.Sprintf("%s=%d", envSiteID.String(), siteID),
		fmt.Sprintf("%s=%s", envSiteURL.String(), siteURL),
	}
}
