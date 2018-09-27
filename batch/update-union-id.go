package main

import (
	"errors"
	"log"
	"runtime"
	"sync"

	"upper.io/db.v3/lib/sqlbuilder"
)

// UpdateUnionId 更新201所有用户的 unionid
func UpdateUnionId(sess sqlbuilder.Database) {
	thirdID, appid := 5, "wx7cc34ee2a43fe772"
	if isProd {
		thirdID, appid = 207, "wxbaed49c053350603"
	}
	cpus := runtime.NumCPU() * 50
	batchSize := cpus
	var w sync.WaitGroup
	for {
		var users []User
		error := sess.Select("id", "wx_id").
			From("e_optometry_user").
			Where("third_id=?", thirdID).
			And("unionid is null").
			And("wx_id is not null").
			// And("is_subscribe = 1").
			Limit(batchSize).All(&users)
		if error != nil {
			log.Printf("limit error: %v", error)
			break
		}
		if len(users) == 0 {
			log.Printf("no more users ")
			break
		}
		for _, i := range users {
			w.Add(1)
			go func(u User) {
				defer w.Done()
				handle(appid, u)
			}(i)
		}
		w.Wait()
	}
}

func handle(appid string, u User) error {
	result, error := req.GetUserInfo(appid, u.WxId.String)
	if error != nil {
		return error
	}
	if errcode, ok := result["errcode"].(float64); ok && errcode > 0 {
		log.Printf("invalid appid or openid: %v, %v\n", result, u.WxId.String)
		return errors.New("invalid appid")
	}
	unionID, ok := result["unionid"].(string)
	if !ok {
		return nil
	}
	_, error = sess.Update("e_optometry_user").Set("unionid=?", unionID).Where("id=?", u.Id).Exec()
	if error == nil {
		log.Printf("process id: %d ", u.Id)
	}
	return error
}
