package main

import (
	// "database/sql"
	// "time"
	"eglass.com/entities"
	// "database/sql"
	// "encoding/json"
	"context"
	_ "eglass.com/entities"
	// "log"
	// "time"
	"eglass.com/utils"
	"upper.io/db.v3/lib/sqlbuilder"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"upper.io/db.v3/mysql"
)

var (
	settings = mysql.ConnectionURL{
		Host:     "t.epeijing.cn",
		Database: "db_new",
		User:     "root",
		Password: "epeijing",
	}
)

type User struct {
	ID		uint   `db:"id" json:"id"`
	Phone	entities.NullString   `db:"phone" json:"phone"`
	StoreID entities.NullInt64 `db:"store_id" json:"store_id"`
	CreateTime entities.NullTime `db:"create_time" json:"create_time"`
}


var logger = utils.NewMysqlLogger(true)
var sess, err = mysql.Open(settings)


func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// sess, err := mysql.Open(settings)
	if err != nil {
		panic(err)
	}
	sess.SetLogging(true)
	sess.SetLogger(logger)
	defer sess.Close()
	// u := User{1, "jj", entities.NullTime{time.Now(), true}}
	// r, _ := json.Marshal(u)
	// log.Println(string(r))
	// return
  	// Route
  	e.GET("/hello", func(c echo.Context) (err error) {
		go sess.Tx(context.Background(), func(tx sqlbuilder.Tx)error{
			_, error := tx.Update("e_optometry_user").Set("name", "王博1").Where("id", 183).Exec()
			if error != nil { return error }
			
			_, error = tx.Update("e_optometry_user").Set("name", "王博2").Where("id", 183).Exec()
			if error != nil {
				return error
			}
			return nil
		})
		_, err = sess.Update("e_optometry_user").Set("name", "王博2").Where("id", 183).Exec()
		if err != nil {
			return
		}
		var u User
		err = sess.Select("id", "name", "store_id", "create_time").From("e_optometry_user").Limit(2).One(&u)
		if err != nil {
			return
		}
		c.JSON(200, u)
		return nil
	})
	e.GET("/h2", StaffHandler)
	e.Logger.Fatal(e.Start(":3002"))
	
}


