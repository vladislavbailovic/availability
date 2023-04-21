package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	if err := Run(ctx, 1312, "http://puppychowfoo.rocks"); err != nil {
		log.Fatalf("Error running a task: %v", err)
	}
	log.Println("-- all good --")
}
