package main

import (
	"availability/pkg/data/model"
	"strings"
	"testing"
)

func Test_getJobName(t *testing.T) {
	suite := map[string]string{
		"http://puppychowfoo.rocks":      "ping-161-puppychowfoo.rocks",
		"https://puppychowfoo.rocks":     "ping-161-puppychowfoo.rocks",
		"http://puppy.chowfoo.rocks":     "ping-161-puppy.chowfoo.rocks",
		"http://puppy.chow.foo.rocks":    "ping-161-puppy.chow.foo.rocks",
		"http://puppychowfoo.rocks:80":   "ping-161-puppychowfoo.rocks80",
		"http://puppychowfoo.rocks/test": "ping-161-puppychowfoo.rocks",
	}
	for test, want := range suite {
		t.Run(test, func(t *testing.T) {
			got := getJobName(161, test)
			if want != got {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}

func Test_getJobEnv(t *testing.T) {
	task := &model.Task{
		Source: &model.Source{
			SiteID: 161,
			URL:    "http://puppychowfoo.rocks",
		},
	}
	e := getJobEnv(task)

	if len(e) != int(_envNamesCount) {
		t.Errorf("unexpected env size: %d (wanted %d)", len(e), _envNamesCount)
	}

	if !strings.Contains(e[0], envSiteID.String()) {
		t.Errorf("missing site ID env var: %v", e[0])
	}
	if !strings.Contains(e[0], "161") {
		t.Errorf("invalid site ID env var: %v", e[0])
	}

	if !strings.Contains(e[1], envSiteURL.String()) {
		t.Errorf("missing site URL env var: %v", e[1])
	}
	if !strings.Contains(e[1], "http://puppychowfoo.rocks") {
		t.Errorf("invalid site URL env var: %v", e[1])
	}
}
