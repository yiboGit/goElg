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
	mysql, err    = utils.NewMysql(false, false)
	request       = utils.InitRequest()
	index         = IsPord(false)
	client, esErr = elastic.NewClient(elastic.SetURL("http://elastic:changeme@101.132.73.244:9200"),
		elastic.SetSniff(false))
	ctx = context.Background()
)

func IsPord(prod bool) string {
	if prod {
		return "member"
	} else {
		return "member_test"
	}
}

var (
	pinyinAnalysis = `{
		"index" : {
			"analysis" : {
				"analyzer" : {
					"pinyin_analyzer" : {
						"tokenizer" : "my_pinyin"
						}
				},
				"tokenizer" : {
					"my_pinyin" : {
						"type" : "pinyin",
						"keep_separate_first_letter" : true,
						"keep_full_pinyin" : true,
						"keep_original" : true,
						"limit_first_letter_length" : 16,
						"lowercase" : true,
						"remove_duplicated_term" : true
					}
				}
			}
		}
	}`

	mapping = `{
		"optometryUser": {
		  "properties": {
			"name": {
			  "type": "text",
			  "analyzer": "ik_max_word",
			  "search_analyzer": "ik_max_word"
			},
			"name_py": {
			  "type": "text",
			  "analyzer": "pinyin_analyzer"
	
			},
			"remark": {
			  "type": "text",
			  "analyzer": "ik_max_word",
			  "search_analyzer": "ik_max_word"
			},
			"phone_wx":{
				"type":"text"
			}
		  }
		}
	  }`
)

type Member struct {
	Id      int32          `db:"id" json:"id"`
	Name    string         `db:"name" json:"name"`
	Remark  sql.NullString `db:"remark" json:"remark"`
	PhoneWx sql.NullString `db:"phone_wx" json:"phoneWx"`
	ThirdId int32          `db:"third_id" json:"thirdId"`
}

func main() {
	deleteIndex()
	createIndex()
	inputData()
}

func deleteIndex() {
	deleteIndex, err := client.DeleteIndex(index).Do(ctx)
	if err != nil {
		log.Panic(err)
	}
	if !deleteIndex.Acknowledged {
		// Not acknowledgedf
		log.Panicf("Not acknowledged,acknowledged is %v \n", deleteIndex.Acknowledged)
	}
}

func createIndex() {
	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		log.Panic(err)
	}
	if exists {
		log.Panic("index is exists")
	}

	createIndex, err := client.CreateIndex(index).BodyString(pinyinAnalysis).Do(ctx)
	if err != nil {
		log.Panic(err)
	}
	if !createIndex.Acknowledged {
		// Not acknowledged
		log.Panicf("Not acknowledgedf,acknowledged is %v \n", createIndex.Acknowledged)
	}
	putMapping, err := client.PutMapping().Index(index).Type("optometryUser").BodyString(mapping).Do(ctx)
	if err != nil {
		log.Panic(err)
	}
	if !putMapping.Acknowledged {
		// Not acknowledged
		log.Panicf("Not acknowledgedf,acknowledged is %v \n", putMapping.Acknowledged)
	}

}

func inputData() {
	var members []Member
	err := mysql.Select("id", "name", "remark", "phone_wx", "third_id").From("e_optometry_user").All(&members)
	if err != nil {
		log.Printf("post err : %v \n", err)
	}
	urlStr := fmt.Sprintf("http://elastic:changeme@101.132.73.244:9200/%s/optometryUser", index)
	for _, member := range members {
		url := fmt.Sprintf("%s/%v", urlStr, member.Id)
		body := utils.Body{
			"id":       member.Id,
			"name":     member.Name,
			"name_py":  member.Name,
			"remark":   member.Remark.String,
			"phone_wx": member.PhoneWx.String,
			"third_id": member.ThirdId,
		}
		log.Printf("member : %v \n", body)
		_, err := request.JPost(url, utils.Query{}, body)
		if err != nil {
			log.Printf("post err : %v \n", err)
		}
	}
}
