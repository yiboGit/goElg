package main

import (
	"log"

	"github.com/labstack/echo"
)

func (instance *UserInstace) bind() {
	instance.e.Add("POST", "/UserServer/AddUser", func(c echo.Context) error {
		var p UserReq
		c.Bind(&p)
		result := instance.AddUser(p)
		log.Printf("Add 判断有无错误")
		if result.Err.Message != "" {
			log.Printf("Add 判断有错")
			result.Err.StatusCode = 500
			return c.JSON(500, result)
		}
		log.Printf("Add 判断无错")
		result.Err.StatusCode = 200
		return c.JSON(200, result)
	})
	instance.e.Add("POST", "/UserServer/GetUser", func(c echo.Context) error {
		var p UserReq
		c.Bind(&p)
		result := instance.GetUser(p)
		if result.Err.Message != "" {
			log.Printf("Get 有错， resp : %v", result)
			result.Err.StatusCode = 500
			return c.JSON(500, result)
		}
		result.Err.StatusCode = 200
		return c.JSON(200, result)
	})

}
