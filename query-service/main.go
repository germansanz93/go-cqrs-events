package main

import (
	"fmt"
	"germansanz93/go/cqrs/database"
	"germansanz93/go/cqrs/events"
	"germansanz93/go/cqrs/repository"
	"germansanz93/go/cqrs/search"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PostgresDB           string `envconfig:"POSTGRES_DB"`
	PostgresUser         string `envconfig:"POSTGRES_USER"`
	PostgresPassword     string `envconfig:"POSTGRES_PASSWORD"`
	NatsAddress          string `envconfig:"NATS_ADDRESS"`
	ElasticSearchAddress string `enconfig:"ELASTICSEARCH_ADDRESS"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatalf("%v", err)
	}

	addr := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	repo, err := database.NewPostgresRepository(addr)

	if err != nil {
		log.Fatal(err)
	}

	repository.SetRepository(repo)

	es, err := search.NewElastic(fmt.Sprintf("http://%s", cfg.ElasticSearchAddress))
	if err != nil {
		log.Fatal(err)
	}

	search.SetSearchRepository(es)

	defer search.Close()

	n, err := events.NewNats(fmt.Sprintf("nats://%s", cfg.NatsAddress))

	if err != nil {
		log.Fatal(err)
	}

	err = n.OnCreateFeed(onCreatedFeed)

	if err != nil {
		log.Fatal(err)
	}

	events.SetEventStore(n)

	defer events.Close()

	router := newRouter()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
