package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const d = `{"act_id":97,"rules":[{"num":1,"below":"您的好友{nickname}帮你扫码啦，还差{num}个好友助力，即可获得镜片八折抵扣券。","text_msg":"恭喜您获得镜片八折抵扣券！\n【兑换券使用方式】：\n领取之后到店，向店员出示兑换券进行核销。\n礼品多多，精彩不断，还差2个好友助力，可获得护眼贴一个。\n公众号回复【海伦凯勒】，了解最新活动情况。","card_msg":"pg41I1kQAz5cAOCvDbpQbGE6byUo"},{"num":3,"below":"您的好友{nickname}帮你扫码啦，还差{num}个好友助力，即可获得护眼贴一个。","text_msg":"恭喜您获得护眼贴一个！\n【兑换券使用方式】：\n领取之后到店，向店员出示兑换券进行核销。\n公众号回复【海伦凯勒】，了解最新活动情况。","card_msg":"pg41I1tQxijibV8lj_b9JZg2i7gg"}],"appid":"wx7cc34ee2a43fe772","user":{"id":3,"name":"年轻就要不停的奔跑","head_url":"http://img.schoolgater.com/weixin-head/og41I1tvkF5GuknoqvBtmMWAonAo.jpg"},"fromUserOpenid":"og41I1jqmclx8rop3yczl9jp5SKc"}`

func TestPush(t *testing.T) {
	var w sync.WaitGroup
	w.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer w.Done()
			redisClient.Publish("channel", d)
		}()
	}
	w.Wait()
	assert.Equal(t, nil, nil)

}
