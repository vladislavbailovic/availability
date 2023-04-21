package main

import "testing"

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
