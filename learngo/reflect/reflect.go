package main

import (
	"fmt"
	"reflect"
)

func say(text string) {
	fmt.Println("say", text)
}

func tell(text string) {
	fmt.Println("tell", text)
}

func imply(text string) {
	fmt.Println("imply", text)
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value) {
	f := reflect.ValueOf(m[name])
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}

	result = f.Call(in)
	return
}

func main() {
	var funcMap = make(map[string]interface{})
	funcMap["say"] = say
	funcMap["tell"] = tell
	funcMap["imply"] = imply
	Call(funcMap, "say", "hello")
	Call(funcMap, "tell", "helloworld")
	Call(funcMap, "imply", "hello, golang,world")
}
