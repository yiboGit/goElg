package main

import (
	"log"

	"eglass.com/entities"
	"eglass.com/utils"
)

type User entities.EOptometryUser

var (
	isProd         = os.Getenv("prod") == "true"
	req            = utils.InitRequest()
	sess, sqlerror = utils.NewMysql(isProd, !isProd)
)

func main() {
	if sqlerror != nil {
		log.Fatal(sqlerror)
	}
	defer sess.Close()
	UpdateUnionId(sess)
}
