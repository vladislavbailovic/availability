package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func Run(ctx context.Context, siteID int, siteURL string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	lnk, err := url.Parse(siteURL)
	if err != nil {
		return err
	}

	jobName := fmt.Sprintf("ping-%d-%s", siteID, strings.Replace(lnk.Host, ":", "-", -1))
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
		jobName)
	if err != nil {
		return err
	}

	log.Printf("Successfully created task: %#v", resp)

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	log.Println("Job started")
	return nil
}
