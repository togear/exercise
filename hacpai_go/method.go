package main

import "fmt"

type Data struct {
	x int
}

func (self Data)ValueTest() {
	fmt.Printf("Value： %p\n",&self)
}

func (self *Data)PointerTest() {
	fmt.Printf("Pointer: %p\n",self)
}

type User struct {
	id int
	name string
}

type Manager struct {
	User
}

func (self *User)ToString() string {
	return fmt.Sprintf("User: %p,%v",self,self)
}

func main() {
	
	d := Data{}
	p := &d
	fmt.Printf("Data : %p\n",p)

	d.ValueTest()
	d.PointerTest()

	p.ValueTest()
	p.PointerTest()

	m := Manager{User{1,"tom"}}

	fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.ToString())
}
