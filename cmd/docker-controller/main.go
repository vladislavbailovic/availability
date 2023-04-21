package main

import (
	"context"
	"log"
	"sync"

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

	for _, task := range tasks {
		go func(task ping) {
			if err := Run(ctx, cli, task.SiteID, task.URL); err != nil {
				log.Printf("Error running a task %v: %v (%T)", task, err, err)
			}
			wg.Done()
		}(task)
	}

	wg.Wait()
	log.Println("-- all good --")
}

type ping struct {
	SiteID int
	URL    string
}

func getSitesToPing() []ping {
	// TODO: implement fetching pings
	// This is going to be something like:
	// SELECT * FROM sites WHERE toPing=1 AND somehow-last-pinged WITHIN <PING_INTERVAL+1>
	return []ping{
		ping{
			SiteID: 1312,
			URL:    "https://snap42.wpmudev.host",
		},
		ping{
			SiteID: 161,
			URL:    "http://puppychowfoo.rocks",
		},
	}
}
