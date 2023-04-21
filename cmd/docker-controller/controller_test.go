package main

import (
	"context"
	"errors"
	"testing"

	"github.com/docker/docker/api/types/container"
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

func Test_StopContainer(t *testing.T) {
	if err := stop(new(fakeStopperError), context.TODO(), 1, "test"); err == nil {
		t.Error("expected error")
	}

	s := new(fakeStopperSuccess)
	if err := stop(s, context.TODO(), 1, "test"); err != nil {
		t.Errorf("expected success, got %v", err)
	} else if *s.opts.Timeout != 0 {
		t.Errorf("invalid options: %v", s.opts)
	}
}

type fakeStopperError struct{}

func (x *fakeStopperError) ContainerStop(_ context.Context, _ string, _ container.StopOptions) error {
	return errors.New("fake error")
}

func (x *fakeStopperError) ContainerWait(_ context.Context, _ string, _ container.WaitCondition) (<-chan container.WaitResponse, <-chan error) {
	ch := make(chan container.WaitResponse)
	go func() {
		ch <- container.WaitResponse{}
	}()
	return ch, make(chan error)
}

type fakeStopperSuccess struct{ opts container.StopOptions }

func (x *fakeStopperSuccess) ContainerStop(_ context.Context, _ string, opts container.StopOptions) error {
	x.opts = opts
	return nil
}
func (x *fakeStopperSuccess) ContainerWait(_ context.Context, _ string, _ container.WaitCondition) (<-chan container.WaitResponse, <-chan error) {
	ch := make(chan container.WaitResponse)
	go func() {
		ch <- container.WaitResponse{}
	}()
	return ch, make(chan error)
}
