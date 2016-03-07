package main

import "fmt"

type Stringer interface {
	String() string
}

type Printer interface {
	Stringer
	Print()
}

type User struct {
	id int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("user %d,%s",self.id,self.name)
}

func (self *User)Print() {
	fmt.Println(self.String())
}

func Print(v interface{}) {
	fmt.Printf("%T: %v\n",v,v)
}

func main() {
	var t Printer = &User{1,"tom"}
	t.Print()
	
	Print(1)
	Print("Hello World!")

	u := User{1, "tom"}
	var vi,pi interface{} = u, &u

	pi.(*User).name="jack"

	fmt.Printf("vi %v\n",vi.(User))

	fmt.Printf("pi %v\n",pi.(*User))

	var t1 Printer = &u
	t1.Print()
}
