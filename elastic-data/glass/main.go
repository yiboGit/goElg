package main

import (
	"context"
	"database/sql"
	"log"

	"eglass.com/utils"
	"github.com/olivere/elastic"
)

var (
	mysql, err    = utils.NewMysql(false, false)
	request       = utils.InitRequest()
	index         = "m_sku"
	client, esErr = elastic.NewClient(elastic.SetURL("http://elastic:changeme@101.132.73.244:9200"), elastic.SetSniff(false))
	ctx           = context.Background()
)

var (
	skuIndex = `{
		"mappings": {
		  "glass": {
			"properties": {
			  "brand": {
				"type": "text",
				"analyzer": "ik_max_word",
				"search_analyzer": "ik_max_word"
			  },
			  "name": {
				"type": "text",
				"analyzer": "ik_max_word",
				"search_analyzer": "ik_max_word"
			  },
			  "refraction":{
				  "type":"text"
			  }
			}
		  }
		}
	  }
	`
)

type glass struct {
	Id         int32           `db:"id" json:"id"`
	Name       sql.NullString  `db:"name" json:"name"`
	Brand      sql.NullString  `db:"brand" json:"brand"`
	Refraction sql.NullFloat64 `db:"refraction" json:"refraction"`
}

func main() {
	// deleteIndex()
	// createIndex()
	inputData()
}

func deleteIndex() {
	deleteIndex, err := client.DeleteIndex(index).Do(ctx)
	if err != nil {
		log.Panic(err)
	}
	if !deleteIndex.Acknowledged {
		log.Panicf("Not acknowledged,acknowledged is %v \n", deleteIndex.Acknowledged)
	}
}

func createIndex() {
	exits, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		log.Panic(err)
	}
	if exits {
		log.Panic("index is exists")
	}
	createIndex, err := client.CreateIndex(index).BodyString(skuIndex).Do(ctx)
	if err != nil {
		log.Panic(err)
	}
	if !createIndex.Acknowledged {
		// Not acknowledged
		log.Panicf("Not acknowledgedf,acknowledged is %v \n", createIndex.Acknowledged)
	}
}

func inputData() {
	var glasses []glass

	mysql.Select("id", "name", "refraction", "brand").From("m_sku").All(&glasses)
	if err != nil {
		log.Printf("post err : %v \n", err)
	}
	// urlStr := "http://elastic:changeme@101.132.73.244:9200/msku/glass"

	for _, glass := range glasses {
		// url := fmt.Sprintf("%s/%v", urlStr, glass.Id)
		body := utils.Body{
			"id":         glass.Id,
			"name":       glass.Name.String,
			"brand":      glass.Brand.String,
			"refraction": glass.Refraction.Float64,
		}
		log.Printf("member : %v \n", body)
		// _, err := request.JPost(url, utils.Query{}, body)
		// if err != nil {
		// 	log.Printf("post err : %v \n", err)
		// }
	}
}
