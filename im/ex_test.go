package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"testing"

	"eglass.com/utils"

	"github.com/stretchr/testify/assert"
)

func TestSql(t *testing.T) {
	var s sql.NullString
	assert.Equal(t, "", s.String)
}
func TestTemplate(t *testing.T) {
	tem := `{"first":{"title":"你好","content":"你好，有新的客人等待你接入"},"keyword1":{"title":"客户称谓","content":"xxx"},"keyword2":{"title":"咨询时间","content":"xxxx年xx月xx日"},"remark":{"title":"备注","content":"点击查看详情"}}`
	type Item struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	type Template struct {
		First    Item `json:"first"`
		Keyword1 Item `json:"keyword1"`
		Keyword2 Item `json:"keyword2"`
		Remark   Item `json:"remark"`
	}
	var m Template
	json.Unmarshal([]byte(tem), &m)
	p := map[string]interface{}{}
	json.Unmarshal([]byte(tem), &p)
	v := p["first"].(map[string]interface{})
	assert.Equal(t, "你好", v["title"])

	assert.Equal(t, "你好", m.First.Title)
}

func TestSendStaffTempalte(t *testing.T) {
	// error := PushTemplateMsgToStaff(69, "wo", "msgnotify")
	r, error := req.SendTemplateMsg("wx5db40e19a0bd5244", "o4h0Qw_G1Hiz8nieT0AwvzbvRzm0",
		"A7th0-Wf7T6EryAEpJrusFQ4HYNrwu6oHu6iToAM2PI",
		utils.Body{
			"first":    utils.Body{"value": "你好，你的眼镜正在加工"},
			"keyword1": utils.Body{"value": "厦门配镜服务中心"},
			"keyword2": utils.Body{"value": utils.ToFullTimeString(nil)},
			"remark":   utils.Body{"value": "hello"},
		},
		"", nil)
	log.Print(r, error)
	assert.Equal(t, nil, error)

}
