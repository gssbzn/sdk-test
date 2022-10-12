package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mongodb-forks/digest"
	"go.mongodb.org/atlas/mongodbatlas"
)

func main() {
	apiKey := os.Getenv("MDB_API_KEY")
	apiSecret := os.Getenv("MDB_API_SECRET")
	baseURL := os.Getenv("MDB_BASE_URL")
	os.Getenv("MDB_API_KEY")
	t := digest.NewTransport(apiKey, apiSecret)
	tc, err := t.Client()
	if err != nil {
		log.Fatalf(err.Error())
	}

	client, err := mongodbatlas.New(tc, mongodbatlas.SetBaseURL(baseURL))
	if err != nil {
		log.Fatalf(err.Error())
	}

	projects, _, err := client.Projects.GetAllProjects(context.Background(), nil)
	fmt.Printf("%#v\n", projects)
}
