package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
}

func (this *MyStruct) GetName() string {
	return this.name
}

func main() {

	fmt.Println("---------------------")
	var a MyStruct
	b := new(MyStruct)

	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.ValueOf(b))

	a.name = "hello"
	b.name = "hello"

	val := reflect.ValueOf(a).FieldByName("Name")
	fmt.Println(val)
	val = reflect.ValueOf(a).FieldByName("name")
	fmt.Println(val)
	val = reflect.ValueOf(*b).FieldByName("name")
	fmt.Println(val)

	fmt.Println(reflect.ValueOf(a).FieldByName("name").CanSet())
	//	fmt.Println(reflect.ValueOf(b).FieldByName("name").CanSet())
	fmt.Println(reflect.ValueOf(&(a.name)).Elem().CanSet())

	fmt.Println("---------------------")
	var c string = "hello"
	p := reflect.ValueOf(c)
	fmt.Println(p)
	fmt.Println(p.CanSet())
	//	fmt.Println(p.Elem().CanSet())

	p = reflect.ValueOf(&c)
	fmt.Println(p)
	fmt.Println(p.CanSet())
	fmt.Println(p.Elem().CanSet())

	p.Elem().SetString("world")
	fmt.Println(p)
	fmt.Println(c)

}
