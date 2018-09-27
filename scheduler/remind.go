package main

import (
	"log"

	"eglass.com/entities"
	"eglass.com/utils"
	"upper.io/db.v3/lib/sqlbuilder"
)

type ERemindRecord entities.ERemindRecord
type ERemindRecords []ERemindRecord

var request = utils.InitRequest()

func (ct ERemindRecord) DueTimestamp() int64 {
	if ct.RemindTime.Valid {
		return ct.RemindTime.Time.Unix()
	}
	return 0
}

func (ct ERemindRecord) IsCycle() bool {
	return false
}
func (ct ERemindRecord) getCycle() int64 {
	return 0
}
func (ct ERemindRecord) Run(mysql sqlbuilder.Database) (err error) {
	defer func() {
	}()
	var t ERemindRecord
	err = mysql.SelectFrom("e_remind_record").Where("id=?", ct.Id).One(&t)
	var user entities.EOptometryUser
	err = mysql.Select("id", "wx_id").From("e_optometry_user").Where("id=?", t.UserId).One(&user)
	openid := user.WxId
	if !openid.Valid {
		return
	}
	var wechat entities.NWechatOpen
	err = mysql.Select("appid").From("n_wechat_open").Where("third_id=?", t.ThirdId).One(&wechat)
	appid := wechat.Appid.String
	result, err := request.JPost("https://epeijing.cn/api/call/wechat/getDynamicTemplate", nil, utils.
		Body{"appid": appid, "id_short": "OPENTM411453554"})
	if err != nil {
		return
	}
	templateID := result["template_id"].(string)
	var store entities.TryStore
	err = mysql.Select("b.store_name").From("e_optometry a ").Join("try_store b").
		On("a.store_id = b.id").Where("a.id = ?", t.OptometryId).One(&store)
	storeName := ""
	if store.StoreName.Valid {
		storeName = store.StoreName.String
	}
	body := utils.Body{
		"first": utils.Body{
			"value": "您有一条复诊预约请求",
		},
		"keyword1": utils.Body{
			"value": "预约复诊",
		},
		"keyword2": utils.Body{
			"value": storeName,
		},
		"keyword3": utils.Body{
			"value": utils.ToFullTimeString(nil),
		},
		"remark": utils.Body{
			"value": "",
		},
	}
	result, err = request.SendTemplateMsg(appid, openid.String, templateID, body, "", nil)
	// fmt.Print(result)
	status := "success"
	if result["errcode"].(float64) > 0 {
		status = "fail"
		log.Printf("send fail %v", result)
	}
	_, err = mysql.Update("e_remind_record").Set("status=?", status).Where("id=?", ct.Id).Exec()
	return
}
func (ct ERemindRecord) GetID() int {
	return int(ct.Id)
}
