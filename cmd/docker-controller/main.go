package main

import (
	"context"
	"log"
	"sync"

	"availability/pkg/data/model"

	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	tasks := getSitesToPing()

	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for _, t := range tasks {
		go func(t *model.Task) {
			if err := Run(ctx, cli, t); err != nil {
				log.Printf("Error running a task %v: %v (%T)", t, err, err)
			}
			wg.Done()
		}(t)
	}

	wg.Wait()
	log.Println("-- all good --")
}

func getSitesToPing() []*model.Task {
	// TODO: implement fetching pings
	// This is going to be something like:
	// SELECT * FROM sites WHERE toPing=1 AND somehow-last-pinged WITHIN <PING_INTERVAL+1>
	return []*model.Task{
		&model.Task{
			Source: &model.Source{
				SiteID: 1312,
				URL:    "https://snap42.wpmudev.host",
			},
		},
		&model.Task{
			Source: &model.Source{
				SiteID: 161,
				URL:    "http://puppychowfoo.rocks",
			},
		},
	}
}
