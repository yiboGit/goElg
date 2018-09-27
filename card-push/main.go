package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis"

	"eglass.com/utils"
)

var prod = os.Getenv("prod") == "true"
var redisClient = utils.NewRedisClient(prod)
var mysql, _ = utils.NewMysql(prod, true)
var activityUserStats = make(map[string]int)
var reqs = utils.InitRequest()

var rw sync.RWMutex

type Rule struct {
	Num     int    `json:"num"`
	Below   string `json:"below"`
	TextMsg string `json:"text_msg"`
	CardMsg string `json:"card_msg"`
}
type HelpAction struct {
	Actid       int64 `json:"act_id"`
	CurrentSubs int
	Rules       []Rule `json:"rules"`
	// 助力user
	User struct {
		Name string `json:"name"`
	} `json:"user"`
	Appid string `json:"appid"`
	// 被助力的user
	FromUserOpenid string `json:"fromUserOpenid"`
}

func main() {
	pb := redisClient.Subscribe("channel")
	for {
		select {
		case event := <-pb.Channel():
			handle(event.Payload)
		}
	}
}
func handle(receMessage string) (err error) {
	log.Print(receMessage)
	var helpAction HelpAction
	err = json.Unmarshal([]byte(receMessage), &helpAction)
	if err != nil {
		return err
	}

	// 被助力的好友 fromUserOpenid
	keyHelpedUser := fmt.Sprintf("%d-%s", helpAction.Actid, helpAction.FromUserOpenid)
	var helperCount int
	helperCount, ok := activityUserStats[keyHelpedUser]
	if ok {
		helperCount++
	} else {
		row, err := mysql.QueryRow("select count(*) from e_user_activity a join e_optometry_user b on a.from_id = b.id where b.wx_id = ? and a.activity_id = ? ", helpAction.FromUserOpenid, helpAction.Actid)
		if err != nil {
			return err
		}
		row.Scan(&helperCount)
	}
	activityUserStats[keyHelpedUser] = helperCount
	helpAction.CurrentSubs = helperCount

	go checkRules(&helpAction)
	time.AfterFunc(time.Minute*10, func() {
		rw.Lock()
		delete(activityUserStats, keyHelpedUser)
		rw.Unlock()
	})
	return
}
func checkRules(ha *HelpAction) (err error) {
	log.Printf("user %s receive %d \n", ha.FromUserOpenid, ha.CurrentSubs)
	giftKey := fmt.Sprintf("gifts-%d-%s", ha.Actid, ha.FromUserOpenid)
	isNil := err == redis.Nil
	for _, rule := range ha.Rules {
		if ha.CurrentSubs < rule.Num && isExp2(ha.CurrentSubs) {
			s := strings.Replace(rule.Below, "{nickname}", ha.User.Name, -1)
			s = strings.Replace(s, "{num}", fmt.Sprintf("%d", rule.Num-ha.CurrentSubs), -1)
			log.Print(s)
			reqs.KfSendMessage(ha.Appid, ha.FromUserOpenid, "text", map[string]interface{}{"content": s})
			break
		}
		if ha.CurrentSubs == rule.Num {
			err = redisClient.SAdd(giftKey, rule.CardMsg).Err()
			if isNil {
				err = redisClient.Expire(giftKey, 31*time.Hour*24).Err()
			}
			if err != nil {
				break
			}
			kr, _ := reqs.KfSendMessage(ha.Appid, ha.FromUserOpenid, "text", map[string]interface{}{"content": fmt.Sprintf("您的好友%s帮您扫码了", ha.User.Name)})
			log.Printf("kr: %v", kr)
			reqs.KfSendMessage(ha.Appid, ha.FromUserOpenid, "text", map[string]interface{}{"content": rule.TextMsg})
			r, err := reqs.KfSendMessage(ha.Appid, ha.FromUserOpenid, "wxcard", map[string]interface{}{
				"card_id": rule.CardMsg,
			})
			if err != nil {
				log.Printf("send card req error: %v\n", err)
				break
			}
			log.Printf("send card result: %v\n", r)
			errMsg := ""
			if code, _ := r["errcode"].(float64); code > 0 {
				errMsg = r["errmsg"].(string)
			}
			uri := fmt.Sprintf("https://%s/api/call/card/sendCard", utils.GetHost(prod))
			reqs.JPost(uri, nil, utils.Body{
				"appid":       ha.Appid,
				"openid":      ha.FromUserOpenid,
				"card_id":     rule.CardMsg,
				"activity_id": ha.Actid,
				"err_msg":     errMsg,
			})

		}
	}

	return err
}
func isExp2(a int) bool {
	if a == 1 {
		return true
	}
	if a > 1024 {
		return true
	}
	r := 1

	for i := 1; i < 10; i++ {
		if a == r {
			return true
		}
		r = 2 * r
	}
	return false
}

/*
{"act_id":97,"rules":[{"num":1,"below":"您的好友{nickname}帮你扫码啦，还差{num}个好友助力，即可获得镜片八折抵扣券。","text_msg":"恭喜您获得镜片八折抵扣券！\n【兑换券使用方式】：\n领取之后到店，向店员出示兑换券进行核销。\n礼品多多，精彩不断，还差2个好友助力，可获得护眼贴一个。\n公众号回复【海伦凯勒】，了解最新活动情况。","card_msg":"pg41I1kQAz5cAOCvDbpQbGE6byUo"},{"num":3,"below":"您的好友{nickname}帮你扫码啦，还差{num}个好友助力，即可获得护眼贴一个。","text_msg":"恭喜您获得护眼贴一个！\n【兑换券使用方式】：\n领取之后到店，向店员出示兑换券进行核销。\n公众号回复【海伦凯勒】，了解最新活动情况。","card_msg":"pg41I1tQxijibV8lj_b9JZg2i7gg"}],"appid":"wx7cc34ee2a43fe772","user":{"id":3,"name":"年轻就要不停的奔跑","head_url":"http://img.schoolgater.com/weixin-head/og41I1tvkF5GuknoqvBtmMWAonAo.jpg"},"fromUserOpenid":"og41I1jqmclx8rop3yczl9jp5SKc"}

*/
