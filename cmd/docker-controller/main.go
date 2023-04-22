package main

import (
	"context"
	"log"
	"sync"

	"availability/pkg/data/collections"
	"availability/pkg/data/model"
	"availability/pkg/data/sql"

	"github.com/docker/docker/client"
)

const (
	maxActiveTasks          int = 5
	pingTimeoutSecs         int = 120
	maxResponseDurationSecs int = 10
)

func main() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	query := new(sql.TaskCollection)
	if err := query.Connect(); err != nil {
		panic("unable to connect")
	}
	defer query.Disconnect()

	tmout := pingTimeoutSecs + maxResponseDurationSecs
	tasks, err := collections.GetActiveTasks(query, maxActiveTasks, tmout)
	if err != nil {
		panic(err)
	}

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
