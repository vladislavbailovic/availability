package main

import (
	"context"
	"errors"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/errdefs"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
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
	if err := Stop(context.TODO(), new(fakeStopperError), 1, "test"); err == nil {
		t.Error("expected error")
	}

	s := new(fakeStopperSuccess)
	if err := Stop(context.TODO(), s, 1, "test"); err != nil {
		t.Errorf("expected success, got %v", err)
	} else if *s.opts.Timeout != 0 {
		t.Errorf("invalid options: %v", s.opts)
	}
}

func Test_RunContainer_ConflictResolved(t *testing.T) {
	s := new(fakeCRSOneConflictError)
	if err := Run(context.TODO(), s, 1312, "test"); err != nil {
		t.Errorf("expected conflict to resolve")
	}
}

func Test_RunContainer_ConflictRepeated(t *testing.T) {
	s := new(fakeCRSRepeatedConflict)
	if err := Run(context.TODO(), s, 1312, "test"); err == nil {
		t.Error("repeated conflict should error out")
	} else if !errdefs.IsConflict(err) {
		t.Errorf("unexpected error: %T %v", err, err)
	}
}

func Test_RunContainer_HappyPath(t *testing.T) {
	s := new(fakeCRSHappyPath)
	if err := Run(context.TODO(), s, 1312, "test"); err != nil {
		t.Errorf("unexpected error: %T %v", err, err)
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

type fakeCRSOneConflictError struct{ counter int }

func (x *fakeCRSOneConflictError) ContainerStop(_ context.Context, _ string, opts container.StopOptions) error {
	return nil
}
func (x *fakeCRSOneConflictError) ContainerWait(_ context.Context, _ string, _ container.WaitCondition) (<-chan container.WaitResponse, <-chan error) {
	ch := make(chan container.WaitResponse)
	go func() {
		ch <- container.WaitResponse{}
	}()
	return ch, make(chan error)
}
func (x *fakeCRSOneConflictError) ContainerCreate(context.Context, *container.Config, *container.HostConfig, *network.NetworkingConfig, *specs.Platform, string) (container.CreateResponse, error) {
	if x.counter > 0 {
		return container.CreateResponse{}, nil
	}
	x.counter += 1
	return container.CreateResponse{}, errdefs.Conflict(errors.New("test"))
}
func (x *fakeCRSOneConflictError) ContainerStart(context.Context, string, types.ContainerStartOptions) error {
	return nil
}

type fakeCRSRepeatedConflict struct{}

func (x *fakeCRSRepeatedConflict) ContainerStop(_ context.Context, _ string, opts container.StopOptions) error {
	return nil
}
func (x *fakeCRSRepeatedConflict) ContainerWait(_ context.Context, _ string, _ container.WaitCondition) (<-chan container.WaitResponse, <-chan error) {
	ch := make(chan container.WaitResponse)
	go func() {
		ch <- container.WaitResponse{}
	}()
	return ch, make(chan error)
}
func (x *fakeCRSRepeatedConflict) ContainerCreate(context.Context, *container.Config, *container.HostConfig, *network.NetworkingConfig, *specs.Platform, string) (container.CreateResponse, error) {
	return container.CreateResponse{}, errdefs.Conflict(errors.New("test"))
}
func (x *fakeCRSRepeatedConflict) ContainerStart(context.Context, string, types.ContainerStartOptions) error {
	return nil
}

type fakeCRSHappyPath struct{}

func (x *fakeCRSHappyPath) ContainerStop(_ context.Context, _ string, opts container.StopOptions) error {
	return nil
}
func (x *fakeCRSHappyPath) ContainerWait(_ context.Context, _ string, _ container.WaitCondition) (<-chan container.WaitResponse, <-chan error) {
	ch := make(chan container.WaitResponse)
	go func() {
		ch <- container.WaitResponse{}
	}()
	return ch, make(chan error)
}
func (x *fakeCRSHappyPath) ContainerCreate(context.Context, *container.Config, *container.HostConfig, *network.NetworkingConfig, *specs.Platform, string) (container.CreateResponse, error) {
	return container.CreateResponse{}, nil
}
func (x *fakeCRSHappyPath) ContainerStart(context.Context, string, types.ContainerStartOptions) error {
	return nil
}
