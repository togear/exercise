package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
	xing string
}

func (this *MyStruct) GetName() string {
	return this.name
}

func (this *MyStruct) GetXing() string {
	return this.xing
}

func (this *MyStruct) GetNameAndXing() (string, string) {
	return this.xing, this.name
}

func main() {
	s := "this is a string"
	fmt.Println(reflect.TypeOf(s))
	fmt.Println("---------------------")
	fmt.Println(reflect.ValueOf(s))

	var x float64 = 3.4
	fmt.Println(reflect.TypeOf(x))
	fmt.Println("---------------------")
	fmt.Println(reflect.ValueOf(x))

	a := new(MyStruct)
	a.name = "hello"
	a.xing = "wang"
	typeOfA := reflect.TypeOf(a)

	fmt.Println(typeOfA.NumMethod())
	fmt.Println("---------------------")
	b := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(b[0])

	b = reflect.ValueOf(a).MethodByName("GetXing").Call([]reflect.Value{})
	fmt.Println(b[0])

	fmt.Println("---------------------")
	b = reflect.ValueOf(a).MethodByName("GetNameAndXing").Call([]reflect.Value{})
	fmt.Println(b[0])
	fmt.Println(b[1])

}
