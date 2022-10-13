package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mongodb-forks/digest"
	mongodbatlas "go.mongodb.org/atlas"
)

func main() {
	apiKey := os.Getenv("MDB_API_KEY")
	apiSecret := os.Getenv("MDB_API_SECRET")
	baseURL := os.Getenv("MDB_BASE_URL")

	t := digest.NewTransport(apiKey, apiSecret)
	tc, err := t.Client()
	if err != nil {
		log.Fatalf(err.Error())
	}
	cfg := &mongodbatlas.Configuration{
		HTTPClient: tc,
		Servers: mongodbatlas.ServerConfigurations{
			{
				URL: baseURL,
			},
		},
	}
	client := mongodbatlas.NewAPIClient(cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	resp, r, err := client.ProjectsApi.ReturnAllProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ReturnAllProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnAllProjects`: PaginatedAtlasGroupView
	fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ReturnAllProjects`: %v\n", resp)
	fmt.Printf("%#v\n", resp)
}
