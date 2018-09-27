package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"
	"time"

	"eglass.com/utils"
)

const POST = "POST"
const GET = "GET"
const domain = "https://t.epeijing.cn/api/md/open/authenticate"

var reqs = utils.InitRequest()

type ResultFunc func() (bool, error)

func checkUrl(method, uri string, query utils.Query, body utils.Body) ResultFunc {
	return func() (bool, error) {
		var r utils.Result
		var err error
		if method == "POST" {
			r, err = reqs.JPost(uri, query, body)
		} else {
			r, err = reqs.JGet(uri, query)
		}
		if err != nil {
			return false, err
		}
		log.Print(r)
		success, ok := r["success"].(bool)
		if ok && success {
			return true, nil
		}
		return false, fmt.Errorf("%v", r)
	}
}

func checkServices() {
	serviceUrls := make(map[string][]ResultFunc)
	serviceUrls["server-new"] = []ResultFunc{
		checkUrl("POST", "https://epeijing.cn/api/md/open/authenticate", nil, utils.Body{
			"phone":    "15652813691",
			"password": "1",
		}),
	}
	serviceUrls["eglass-helper"] = []ResultFunc{
		checkUrl("POST", "https://epeijing.cn/api/micro/open/login", nil, utils.Body{
			"account":  "15652813693",
			"password": "111",
		}),
	}
	log.Printf("check start")
	var w sync.WaitGroup
	for service, fcs := range serviceUrls {
		w.Add(1)
		go func(s string, fcs []ResultFunc) {
			defer w.Done()
			var err error
			for _, fc := range fcs {
				_, err = fc()
				if err != nil {
					log.Printf("service %s status: failed, reason: %v", s, err)
					sendWarning(s, err)
					break
				}
			}
		}(service, fcs)
	}
}

func sendWarning(service string, err error) error {
	// 胡哥, 小智，小闯
	persons := []string{"og41I1jqmclx8rop3yczl9jp5SKc", "og41I1mv3VZF84-QYxdnnLiz2wuU", "og41I1tvkF5GuknoqvBtmMWAonAo", "og41I1qKtiiWcRuOhdzIJJaOIAuU"}
	var error error
	for _, p := range persons {
		_, error = reqs.SendTemplateMsg("wx7cc34ee2a43fe772", p,
			"8k4xukShFWdzj-hbEd405QfqD4HvaG_Nxue0k9Hek-4",
			utils.Body{
				"first": utils.Body{"value": "服务故障"},
				// 商户名称：{{keyword1.DATA}}
				"keyword1": utils.Body{"value": "Eglass"},
				// 机器编号：{{keyword2.DATA}}
				"keyword2": utils.Body{"value": service},
				// 故障内容：{{keyword3.DATA}}
				"keyword3": utils.Body{"value": err.Error()},
				"keyword4": utils.Body{"value": "登录epeijing.cn, docker restart <containerid> "},
			},
			"", nil)
	}

	return error
}

// 达达
const (
	appID     = "dada116ae2cd04fd419"
	appSecret = "6992c79a6e23b2449c51aadb42d20cd8"
)

var DadaErr = errors.New("dada api error")

func checkDadaBalance() error {
	uri := "/api/balance/query"
	r, err := dataRequest(uri, utils.Body{"category": 1})
	if err != nil {
		return DadaErr
	}
	balance, _ := r.(map[string]interface{})["deliverBalance"].(float64)
	log.Printf("data balance: %d", int64(balance))
	if balance < 200 {
		r, err = dataRequest("/api/recharge", utils.Body{
			"amount":   500,
			"category": "H5",
		})
		if err != nil {
			return err
		}
		sendDataCharge(balance, r.(string))
	}
	return err
}
func dataRequest(endpoint string, body utils.Body) (interface{}, error) {
	bodyStr := ""
	if body != nil {
		bs, _ := json.Marshal(body)
		bodyStr = string(bs)
	}
	a := utils.Body{
		"app_key":   appID,
		"timestamp": fmt.Sprintf("%d", time.Now().Unix()),
		"format":    "json",
		"v":         "1.0",
		"source_id": "3673",
		"body":      bodyStr,
	}
	keys := make([]string, 0)
	for k := range a {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	b.WriteString(appSecret)
	for _, key := range keys {
		b.WriteString(key + a[key].(string))
	}
	b.WriteString(appSecret)
	signature := fmt.Sprintf("%x", md5.Sum([]byte(b.String())))
	a["signature"] = strings.ToUpper(signature)
	r, err := reqs.JPost("http://newopen.imdada.cn"+endpoint, nil, a)
	code := r["code"].(float64)
	if err != nil || code > 0 {
		log.Printf("dada request fail: %v, %v", r, err)
		return nil, DadaErr
	}
	return r["result"], nil
}

func sendDataCharge(amount float64, url string) error {
	// 胡哥, 飞飞和小雯
	persons := []string{"og41I1jqmclx8rop3yczl9jp5SKc", "og41I1sElCa_2qler5Mo08zmDEB8", "og41I1n7P6esAHH54oRbvwAF167Y"}
	var error error
	for _, p := range persons {
		_, error = reqs.SendTemplateMsg("wx7cc34ee2a43fe772", p,
			"8k4xukShFWdzj-hbEd405QfqD4HvaG_Nxue0k9Hek-4",
			utils.Body{
				"first": utils.Body{"value": "达达余额不足"},
				// 商户名称：{{keyword1.DATA}}
				"keyword1": utils.Body{"value": "Eglass 达达余额"},
				// 机器编号：{{keyword2.DATA}}
				"keyword2": utils.Body{"value": "达达余额"},
				// 故障内容：{{keyword3.DATA}}
				"keyword3": utils.Body{"value": fmt.Sprintf("当前余额%.2f", amount)},
				"keyword4": utils.Body{"value": "点我打开充值页面"},
			},
			url, nil)
	}
	return error
}

func main() {
	checkDadaBalance()
	checkServices()
	ticker := time.Tick(time.Second * 60)
	for {
		select {
		case <-ticker:
			log.Printf("check start")
			// 服务检测
			checkServices()
			// 达达余额检测
			checkDadaBalance()
		}

	}

}
