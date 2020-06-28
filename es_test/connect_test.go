package test

import (
	"testing"

	"github.com/olivere/elastic"
)

// connect to es v6.8
func TestConnectES(t *testing.T) {
	url := "http://localhost:9200"
	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	if client == nil {
		t.Fatal("es node is not found")
	}
}
