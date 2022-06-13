package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/graphql"
)

func main() {
	ctx := context.Background()
	nbURL := os.Getenv("NADA_BACKEND_URL")

	if nbURL == "" {
		fmt.Println("nada-backend URL not set")
		os.Exit(1)
	}
	client := graphql.NewClient(nbURL, nil)

	var m struct {
		TriggerMetadataSync bool `graphql:"triggerMetadataSync"`
	}

	if err := client.Mutate(ctx, &m, nil); err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	os.Exit(0)
}
