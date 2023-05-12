package main

import (
	"fmt"
	"net/url"
	"strings"

	"availability/pkg/data/model"
	"availability/pkg/env"
)

func getJobName(siteID int32, siteURL string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "avbl-ping-%d", siteID)

	lnk, err := url.Parse(siteURL)
	if err == nil {
		b.WriteString("-")
		b.WriteString(strings.Replace(lnk.Host, ":", "", -1))
	}

	return b.String()
}

func getJobEnv(task *model.Task) []string {
	var down int32 = 0
	if task.WasPreviouslyDown() {
		down = task.Previous.ProbeID
	}
	return []string{
		fmt.Sprintf("%s=%d", env.SiteID.String(), task.Source.SiteID),
		fmt.Sprintf("%s=%s", env.SiteURL.String(), task.Source.URL),
		fmt.Sprintf("%s=%d", env.PreviouslyDown.String(), down),
	}
}
