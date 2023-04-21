package main

import (
	"context"
	"log"

	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	if err := Run(ctx, cli, 1312, "http://puppychowfoo.rocks"); err != nil {
		log.Fatalf("Error running a task: %v (%T)", err, err)
	}
	log.Println("-- all good --")
}
