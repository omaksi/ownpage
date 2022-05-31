package es

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

//STZ2KZ5e3wD1upqPI3rW
func InitES() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)
}
