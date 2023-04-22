package main

import (
	"context"
	"log"
	"sync"
	"time"

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
	for {
		if err := run(ctx); err != nil {
			panic(err)
		}
		time.Sleep(time.Duration(pingTimeoutSecs) * time.Second)
	}
}

func run(ctx context.Context) error {
	log.Println("task scheduling:")
	ctx, cancel := context.WithTimeout(ctx, time.Duration(pingTimeoutSecs)*time.Second)
	defer cancel()

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	query := new(sql.TaskCollection)
	if err := query.Connect(); err != nil {
		return err
	}
	defer query.Disconnect()

	tmout := pingTimeoutSecs + maxResponseDurationSecs
	tasks, err := collections.GetActiveTasks(query, maxActiveTasks, tmout)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for _, t := range tasks {
		log.Printf("\t- scheduling task %v", t)
		go func(t *model.Task) {
			if err := Run(ctx, cli, t); err != nil {
				log.Printf("Error running a task %v: %v (%T)", t, err, err)
			}
			wg.Done()
		}(t)
	}

	wg.Wait()
	log.Printf("done scheduling: %s", time.Now())
	return nil
}
