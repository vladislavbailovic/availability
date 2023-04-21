package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func Run(ctx context.Context, siteID int, siteURL string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	ccfg := &container.Config{
		Env: []string{
			fmt.Sprintf("AVBL_SITE_ID=%d", siteID),
			fmt.Sprintf("AVBL_SITE_URL=%s", siteURL),
		},
		Image: "availability:job",
	}
	hcfg := &container.HostConfig{
		AutoRemove: true,
	}
	resp, err := cli.ContainerCreate(
		ctx,
		ccfg,
		hcfg,
		nil,
		nil,
		getJobName(siteID, siteURL))
	if err != nil {
		log.Printf("Error starting container: %v", err)
		log.Println("Possibly a runaway task. Re-starting")
		if err := Stop(ctx, siteID, siteURL); err != nil {
			log.Println("Giving up")
			return err
		}

		log.Println("Let's try create one more time")
		resp, err = cli.ContainerCreate(
			ctx,
			ccfg,
			hcfg,
			nil,
			nil,
			getJobName(siteID, siteURL))
		if err != nil {
			log.Printf("Error re-attempting container create: %v", err)
			log.Println("Giving up")
			return err
		}
	}

	log.Printf("Successfully created task: %#v", resp)

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	log.Println("Job started")
	return nil
}

func Stop(ctx context.Context, siteID int, siteURL string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	return stop(cli, ctx, siteID, siteURL)
}

type stopper interface {
	ContainerStop(context.Context, string, container.StopOptions) error
	ContainerWait(context.Context, string, container.WaitCondition) (<-chan container.WaitResponse, <-chan error)
}

func stop(cli stopper, ctx context.Context, siteID int, siteURL string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	jobName := getJobName(siteID, siteURL)

	resCh, errCh := cli.ContainerWait(ctx, jobName, container.WaitConditionRemoved)

	tmout := 0
	opts := container.StopOptions{Timeout: &tmout}
	if err := cli.ContainerStop(ctx, jobName, opts); err != nil {
		log.Printf("Error stopping container: %v", err)
		return err
	}

	select {
	case <-resCh:
		log.Println("Container apparently stopped")
	case err := <-errCh:
		log.Printf("Error stopping container: %v", err)
	}
	return nil
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
