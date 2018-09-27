package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"eglass.com/utils"
	"github.com/olivere/elastic"
)

var (
	esClient, esError = elastic.NewClient(elastic.SetURL("http://elastic:changeme@101.132.73.244:9200"),
		elastic.SetSniff(false))
	ctx       = context.Background()
	index     = IsPord(false)
	centerSku = `{
		"mappings": {
		  "center": {
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
			   "serial": {
				"type": "text",
				"analyzer": "ik_max_word",
				"search_analyzer": "ik_max_word"
			  },
			  "refraction":{
				  "type":"text"
				},
				"deleted":{
				  "type":"keyword"
				},
				"category":{
				  "type":"keyword"
				},
				"goods_status":{
				  "type":"keyword"
			  }
			}
		  }
		}
	  }`
)

func IsPord(prod bool) string {
	if prod {
		return "center_sku"
	} else {
		return "center_sku_test"
	}
}

func main() {
	if esError != nil {
		log.Panicf("elastic error : %v \n", esError)
	}
	deleteIndex()
	createIndex()
	inputData()
}

func deleteIndex() {
	exits, err := esClient.IndexExists(index).Do(ctx)
	if err != nil {
		log.Panicf("elastic index exits error : %v \n", err)
	}
	if exits {
		_, err := esClient.DeleteIndex(index).Do(ctx)
		if err != nil {
			log.Panicf("elastic delete index error : %v \n", err)
		}
	}
}

func createIndex() {
	createIndex, err := esClient.CreateIndex(index).Do(ctx)
	if err != nil {
		log.Panicf("elastic create index error : %v \n", err)
	}
	if !createIndex.Acknowledged {
		// Not acknowledged
		log.Panicf("Not acknowledgedf,acknowledged is %v \n", createIndex.Acknowledged)
	}
}

func inputData() {
	mysql, err := utils.NewMysql(false, false)
	if err != nil {
		log.Panicf("mysql has error : %v \n", err)
	}
	type center struct {
		Id          int             `db:"id" json:"id"`
		Brand       sql.NullString  `db:"brand" json:"brand"`
		Serial      sql.NullString  `db:"serial" json:"serial"`
		Name        sql.NullString  `db:"name" json:"name"`
		Refraction  sql.NullFloat64 `db:"refraction" json:"refraction"`
		GoodsStatus sql.NullInt64   `db:"goods_status" json:"goodsStatus"`
		Deleted     sql.NullInt64   `db:"deleted" json:"deleted"`
		Category    sql.NullString  `db:"category" json:"category"`
		CenterId    int             `db:"center_id" json:"centerId"`
		IsCustomize sql.NullInt64   `db:"is_customize" json:"isCustomize"`
	}

	var centers []center
	err = mysql.Select("id", "brand", "serial", "name", "refraction", "goods_status", "center_id", "deleted", "category", "is_customize").From("m_center_sku").All(&centers)
	if err != nil {
		log.Panicf("select has error : %v \n", err)
	}

	for _, v := range centers {
		if v.GoodsStatus.Int64 > 0 {
			v.GoodsStatus.Int64 = 1
		}
		center := fmt.Sprintf(`{"id": %d,"brand":"%s","serial":"%s","name":"%s","refraction": "%v","goods_status":"%v","center_id":"%v","deleted":"%v","category":"%v","is_customize":%v}`,
			v.Id, v.Brand.String, v.Serial.String, v.Name.String, v.Refraction.Float64, v.GoodsStatus.Int64, v.CenterId, v.Deleted.Int64, v.Category.String, v.IsCustomize.Int64)
		fmt.Printf("%v", center)
		id := fmt.Sprintf("%v", v.Id)
		// fmt.Printf("%v", id)
		_, err := esClient.Index().Index(index).Type("center").Id(id).BodyString(center).Do(ctx)
		if err != nil {
			log.Printf("input err :%v \n", err)
		}
	}
}
