package main

import (
	"fmt"
	"reflect"
)

func main() {
	// interfaceName := os.Getenv("interface")
	var ptr interface{}
	ptr = (*UserServer)(nil)
	t := reflect.TypeOf(ptr).Elem()
	for i := 0; i < t.NumMethod(); i++ {
		name := t.Method(i).Name
		paramType := t.Method(i).Type.In(0).Name()
		paramType := t.Method(i).Type.In(0).
		fmt.Println(name)
		fmt.Println(paramType)
	}
}
