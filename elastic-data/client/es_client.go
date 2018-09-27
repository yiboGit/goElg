package main

import (
	"context"
	"log"

	"github.com/olivere/elastic"
)

func main() {

	client, err := elastic.NewClient(elastic.SetURL("http://elastic:changeme@101.132.73.244:9200"),
		elastic.SetSniff(false))

	if err != nil {
		log.Printf("error : %v \n", err)
	}

	ctx := context.Background()

	boolQuery := elastic.NewBoolQuery()
	boolQuery.Should(elastic.NewMatchQuery("name", "蔡司新三维博锐1.5"),
		elastic.NewMatchQuery("brand", "蔡司新三维博锐1.5"),
		elastic.NewWildcardQuery("refraction", "*蔡司新三维博锐1.5*"))
	boolQuery.MinimumNumberShouldMatch(1)
	searchResult, err := client.Search().Index("msku").Type("glass").
		Query(boolQuery).Size(50).Do(ctx)
	if err != nil {
		log.Printf("err %v \n", err)
	} else {
		log.Printf("Result %v \n", searchResult.Hits)
	}

}
