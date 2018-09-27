package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Any("/service/test/resp", handleAll)
	e.Logger.Fatal(e.Start(":3030"))
}

func handleAll(c echo.Context) error {
	req := c.Request()
	body, error := ioutil.ReadAll(req.Body)
	// reqBody := bytes.NewReader(body)
	var respBody map[string]interface{}
	json.Unmarshal(body, &respBody)
	requestId := respBody["requestId"]

	fmt.Printf("body: %v", respBody)
	return error
}
