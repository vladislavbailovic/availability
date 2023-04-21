package main

import (
	"log"
	"os"
	"time"
)

func main() {
	siteID := os.Getenv("AVBL_SITE_ID")
	siteURL := os.Getenv("AVBL_SITE_URL")
	for true {
		log.Printf("Time: %s / siteID: %s / site URL: %s", time.Now(), siteID, siteURL)
		time.Sleep(time.Second * 2)
	}
}
