
	package main
	import (
		"github.com/labstack/echo"
	)

	func ( *)bind() {
		 .e.Add("POST","/UserServer/AddUser",func(c echo.Context) error {
			var p UserReq
			c.Bind(&p)
			result,err := .AddUser(p)
			if err != nil {
				return c.JSON(500,result)
			}
			return c.JSON(200,result)
		})
		 .e.Add("POST","/UserServer/DeleteUser",func(c echo.Context) error {
			var p UserReq
			c.Bind(&p)
			result,err := .DeleteUser(p)
			if err != nil {
				return c.JSON(500,result)
			}
			return c.JSON(200,result)
		})
		 .e.Add("POST","/UserServer/UpdateUser",func(c echo.Context) error {
			var p UserReq
			c.Bind(&p)
			result,err := .UpdateUser(p)
			if err != nil {
				return c.JSON(500,result)
			}
			return c.JSON(200,result)
		})
		
	}
	