package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"runtime"
	"sync"

	"database/sql"

	"eglass.com/utils"
)

var (
	mysql, mysqlErr = utils.NewMysql(true, false)
	wechatOpenCache []wechatOpen
	request         = utils.InitRequest()
)

type wechatOpen struct {
	Appid   sql.NullString `db:"appid" json:"appid"`
	ThirdId sql.NullInt64  `db:"third_id" json:"third_id"`
}

func main() {
	defer mysql.Close()
	if mysqlErr != nil {
		log.Panic(mysqlErr)
	}

	err := mysql.Select("appid", "third_id").From("n_wechat_open").All(&wechatOpenCache)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("cache: %v", wechatOpenCache)
	// 根据appid 查出会员卡id，将会员卡id与appid关系，更新在表e_member_card上
	insertEMemberCard()

	// 根据third_id 查出 其对应的wx_id
	for _, wechatOpen := range wechatOpenCache {
		if wechatOpen.ThirdId.Valid && wechatOpen.Appid.Valid {
			updateMemberCode(wechatOpen.ThirdId.Int64, wechatOpen.Appid.String)
		}
	}
}

type CardID struct {
	CardId string `db:"card_id" json:"cardId"`
}

type WxID struct {
	WxId string `db:"wx_id" json:"wxId"`
	ID   string `db:"id" json:"id"`
}

func updateMemberCode(thirdId int64, appid string) {
	var (
		cardID CardID
	)

	qy := utils.Query{"appid": appid}
	err := mysql.Select("card_id").From("e_member_card").Where("third_id=?", thirdId).One(&cardID)
	if err != nil {
		log.Printf("thirdId: %d, get card id error, Err : %v \n", thirdId, err)
		return
	}
	if cardID.CardId == "" {
		log.Printf("thirdId: %d, card_id empty \n", thirdId)
		return
	}

	cpus := runtime.NumCPU() * 50
	batchSize := cpus
	var w sync.WaitGroup

	for {
		var wxIDList []WxID
		err = mysql.Select("wx_id", "id").From("e_optometry_user").
			Where("third_id=?", thirdId).
			And("wx_id is not null").
			And("member_code is null").
			And("phone_wx is not null").Limit(batchSize).All(&wxIDList)
		if err != nil {
			log.Printf("thirdId: %d, select wx_id has error, Err : %v \n", thirdId, err)
			break
		}
		if len(wxIDList) == 0 {
			log.Printf("thirdId: %d ; select wx_id is empty \n", thirdId)
			break
		}

		for _, wxID := range wxIDList {
			w.Add(1)
			go func(wxID WxID) {
				defer w.Done()
				body := utils.Body{
					"openid":  wxID.WxId,
					"card_id": cardID.CardId,
				}
				result, err := request.JPost("https://t.epeijing.cn/api/v2/wechat/card/user/getcardlist", qy, body)
				if err != nil {
					log.Printf("err %v \n", err)
					return
				}
				if result["errcode"].(float64) > 0 {
					log.Printf("result has error %v,errMsg: %v, card_id: %s, wx_id: %s, appid: %s \n", result["errcode"].(float64), result["errmsg"], cardID.CardId, wxID.WxId, appid)
					//错误处理 member_code 空白
					_, err := mysql.Update("e_optometry_user").
						Set("member_code=?", "").
						Where("id=?", wxID.ID).Exec()
					if err != nil {
						log.Printf("errorcode is not nil; wx_id : %s;update e_optometry_user error: %v", wxID.WxId, err)
						return
					}
					return
				}
				log.Printf("result ok, has not error \n")

				if result["card_list"] == nil {
					log.Printf("thrid_id: %d , cardList is empty \n", thirdId)
					_, err := mysql.Update("e_optometry_user").
						Set("member_code=?", "").
						Where("id=?", wxID.ID).Exec()
					if err != nil {
						log.Printf("cardList empty! wx_id : %s;update e_optometry_user error: %v", wxID.WxId, err)
						return
					}
					return
				}
				cardList := result["card_list"].([]interface{})
				if len(cardList) == 0 {
					log.Printf("thrid_id: %d , cardList is empty \n", thirdId)
					_, err := mysql.Update("e_optometry_user").
						Set("member_code=?", "").
						Where("id=?", wxID.ID).Exec()
					if err != nil {
						log.Printf("cardList empty! wx_id : %s;update e_optometry_user error: %v", wxID.WxId, err)
						return
					}
					return
				}

				memberCard := cardList[0].(map[string]interface{})
				memberCode := memberCard["code"].(string)
				_, err = mysql.Update("e_optometry_user").
					Set("member_code=?", memberCode).
					Where("id=?", wxID.ID).Exec()
				if err != nil {
					log.Printf("wx_id : %s;update e_optometry_user error: %v", wxID.WxId, err)
					return
				}
				log.Println("update ok")
			}(wxID)
		}
		w.Wait()
	}
}

func insertEMemberCard() {
	log.Println("Insert into e_member_card")
	for _, wechatOpen := range wechatOpenCache {
		if !wechatOpen.Appid.Valid || !wechatOpen.ThirdId.Valid {
			continue
		}
		appid := wechatOpen.Appid.String
		thirdId := wechatOpen.ThirdId.Int64
		log.Printf("appid: %s get card_list", appid)
		//查出该appid下的会员卡id
		cardID, err := getCardList(appid)
		if err != nil {
			log.Printf("appid: %s get card_list error： %v \n", appid, err)
			continue
		}
		if cardID == "" {
			log.Printf("appid: %s do not has member_card \n", appid)
			continue
		}
		log.Printf("appid: %s , third_id: %d, cardID: %s", appid, thirdId, cardID)
		insertSql := InsertMemberShipSql
		insertSql = fmt.Sprintf(insertSql, thirdId, cardID, thirdId)
		_, err = mysql.Exec(insertSql)
		if err != nil {
			log.Printf("insert e_member_card error : %v", err)
		} else {
			log.Printf("insert e_member_card ok ;appid: %s , third_id: %d, cardID: %s", appid, thirdId, cardID)
		}
	}
}

var InsertMemberShipSql = `insert into e_member_card (third_id, card_id) 
select %d,'%s'
from dual where not exists(select 1 from e_member_card where third_id=%d)
`

func getCardList(appid string) (string, error) {
	var (
		resultCardID string
		totalNum     float64
		err          error
	)
	resultCardID, totalNum, err = cardBatchget(appid, 0)
	if err != nil {
		return "", err
	}
	if totalNum > 50 && resultCardID == "" {
		pages := int(math.Ceil(totalNum / 50))
		for i := 1; i <= pages; i++ {
			resultCardID, totalNum, err = cardBatchget(appid, i*50)
			if err != nil || resultCardID != "" {
				break
			}
		}
	}
	return resultCardID, err
}

func cardBatchget(appid string, offset int) (string, float64, error) {
	var resultCardID string
	qy := utils.Query{"appid": appid}
	body := utils.Body{
		"offset": offset,
		"count":  50,
		"status_list": []string{
			"CARD_STATUS_VERIFY_OK", "CARD_STATUS_DISPATCH",
		}}
	result, err := request.JPost("https://t.epeijing.cn/api/v2/wechat/card/batchget", qy, body)
	if err != nil {
		log.Printf("err %v \n", err)
		return "", 0, err
	}
	if result["errcode"].(float64) > 0 {
		msg := "the appid gat card_list error"
		log.Println(msg)
		return "", 0, errors.New(msg)
	}

	for _, cardId := range result["card_id_list"].([]interface{}) {
		log.Printf("list: %s", cardId.(string))
		result, err := request.JPost("https://t.epeijing.cn/api/v2/wechat/card/get", qy, utils.Body{
			"card_id": cardId.(string),
		})
		if err != nil {
			log.Printf("get cardlist error :%v", err)
			continue
		}
		if result["errcode"].(float64) > 0 {
			continue
		}
		cardType := result["card"].(map[string]interface{})["card_type"]
		if cardType.(string) == "MEMBER_CARD" {
			resultCardID = cardId.(string)
			break
		}
	}
	return resultCardID, result["total_num"].(float64), nil
}
