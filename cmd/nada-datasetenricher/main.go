package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: 5 * time.Minute,
	}

	request, err := http.NewRequest("POST", "http://nada-backend/api/bigquery/tables/sync", nil)
	if err != nil {
		log.Fatalf("building request: %v", err)
	}

	_, err = client.Do(request)
	if err != nil {
		log.Fatalf("sending request: %v", err)
	}
}
