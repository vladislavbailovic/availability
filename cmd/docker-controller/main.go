package main

import (
	"context"
	"log"
	"sync"

	"availability/pkg/data/collections"
	"availability/pkg/data/fakes"
	"availability/pkg/data/model"

	"github.com/docker/docker/client"
)

const maxActiveTasks int = 5

func main() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	query := &fakes.TaskCollection{
		Sources: []fakes.Source{
			fakes.Source{ID: 1312, URL: "https://snap42.wpmudev.host"},
			fakes.Source{ID: 161, URL: "http://puppychowfoo.rocks"},
		},
	}
	tasks, err := collections.GetActiveTasks(query, maxActiveTasks)
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
