package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	// instaceVar := "u"
	// instaceStruce := "UserInstace"
	instaceVar := os.Getenv("instaceVar")
	instaceStruce := os.Getenv("instaceStruce")
	var ptr interface{}
	ptr = (*UserServer)(nil)
	t := reflect.TypeOf(ptr).Elem()
	var s string
	for i := 0; i < t.NumMethod(); i++ {
		name := t.Method(i).Name
		paramName := t.Method(i).Type.In(0).String()
		paramNames := strings.Split(paramName, ".")
		if len(paramNames) > 1 {
			paramName = paramNames[1]
		} else {
			paramName = paramNames[0]
		}
		methodStr := ` %s.e.Add("POST","/%s/%s",func(c echo.Context) error {
			var p %s
			c.Bind(&p)
			result,err := %s.%s(p)
			if err != nil {
				return c.JSON(500,result)
			}
			return c.JSON(200,result)
		})
		`
		s = s + fmt.Sprintf(methodStr, instaceVar, t.Name(), name, paramName, instaceVar, name)
	}
	fmt.Println(s)
	//TODO 在指定目录生成文件
	// fileName := fmt.Sprintf("./bind.go")
	// f, _ := os.Create(fileName)
	f, _ := os.Create("../user_server/bind.go")
	defer f.Close()
	info := `
	package main
	import (
		"github.com/labstack/echo"
	)

	func (%s *%s)bind() {
		%s
	}
	`
	info = fmt.Sprintf(info, instaceVar, instaceStruce, s)
	f.Write([]byte(info))

}
