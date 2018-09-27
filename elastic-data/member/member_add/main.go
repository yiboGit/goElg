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
	mysql, err    = utils.NewMysql(true, false)
	request       = utils.InitRequest()
	index         = IsPord(true)
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

type Member struct {
	Id      int32          `db:"id" json:"id"`
	Name    string         `db:"name" json:"name"`
	Remark  sql.NullString `db:"remark" json:"remark"`
	PhoneWx sql.NullString `db:"phone_wx" json:"phoneWx"`
	ThirdId int32          `db:"third_id" json:"thirdId"`
}

func main() {
	inputData()
}

func inputData() {
	var members []Member
	err := mysql.Select("id", "name", "remark", "phone_wx", "third_id").From("e_optometry_user").Where("create_time between '2018-09-03 22:07:06' and '2018-09-04 09:45:23'").All(&members)
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
