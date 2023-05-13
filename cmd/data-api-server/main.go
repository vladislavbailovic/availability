package main

import (
	"availability/pkg/env"
	"log"
)

func main() {
	apiPort := env.ApiPortData.WithFallback("3667")
	auth := env.ApiSecretData.Value()

	log.Println(apiPort, auth)
}
