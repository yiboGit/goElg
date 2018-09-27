package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	// interfaceName := os.Getenv("interface")
	var ptr interface{}
	ptr = (*UserServer)(nil)
	t := reflect.TypeOf(ptr).Elem()
	var methods string
	paramType := t.Method(0).Type.In(0).String()
	params := strings.Split(paramType, ".")
	paramType = params[1]

	methodStr := `func (client *%s)request(p %s, address string, method string) ( %s ,error){
	var b bytes.Buffer
	u, _ := json.Marshal(p)
	b.Write(u)
	url := address + "/" + "%s" + "/" + method
	resp, err := http.Post(url , "application/json", &b)
	if err != nil {
		return %s{},err
	}
	b.ReadFrom(resp.Body)
	var result %s
	json.Unmarshal(b.Bytes(), &result)
	return result,nil
}
	`
	clientType := fmt.Sprintf("%s%s", t.Name(), "Client")
	methods = methods + fmt.Sprintf(methodStr, clientType, paramType, paramType, t.Name(), paramType, paramType)
	title := `package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type %s struct {}

%s
	`
	flieName := fmt.Sprintf("./%s_client.go", t.Name())
	f, _ := os.Create(flieName)
	defer f.Close()
	info := fmt.Sprintf(title, clientType, methods)
	f.Write([]byte(info))
}
