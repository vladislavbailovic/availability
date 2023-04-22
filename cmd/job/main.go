package main

import (
	"log"
	"os"
	"time"
)

func main() {
	siteID := os.Getenv("AVBL_SITE_ID")
	siteURL := os.Getenv("AVBL_SITE_URL")
	prev := os.Getenv("AVBL_PREVIOUSLY_DOWN")
	for true {
		log.Printf("Time: %s / siteID: %s / site URL: %s / prev: %s", time.Now(), siteID, siteURL, prev)
		time.Sleep(time.Second * 2)
	}
}
