// struct
package main

import (
	"fmt"
	"strings"
)

type User struct {
	name     string
	email    string
	password string
}

var user User
var userref *User

type File struct {
	fd       int
	filename string
}

type file1 struct {
	fd       int
	filename string
}

type TagType struct { //tags
	field1 bool   "An important answer"
	field2 string "the name of the string"
	field3 int    "how much there are"
}

type anonymousStruct struct {
	name string
	int
	string
	File
}

func main() {

	fmt.Println(user)
	fmt.Println(userref)
	userref = new(User)
	fmt.Println(userref)

	//Initailize
	user = User{name: "liming", email: "liming@gmail.com", password: "pw123456"}
	userref = &User{name: "liming", email: "liming@gmail.com", password: "pw123456"}
	fmt.Println(user)
	fmt.Println(userref)

	user1 := User{name: "liming", email: "liming@gmail.com", password: "pw123456"}
	userref1 := &User{name: "liming", email: "liming@gmail.com", password: "pw123456"}
	fmt.Println(user1)
	fmt.Println(userref1)

	//Fixme
	upUser(user1)
	upUserRef(userref1)

	fmt.Printf("name is %s, %s\n", user1.name, userref1.name)

	//tags can get from reflect

	anonymous := new(anonymousStruct)
	anonymous.name = "hanmeimei"
	anonymous.int = 88
	anonymous.string = "english"
	anonymous.File.fd = 10
	anonymous.File.filename = "xxoo.avi"
	fmt.Println(anonymous)

}

func upUser(user User) {
	strings.ToUpper(user.name)
}

func upUserRef(ur *User) {
	strings.ToUpper(ur.name)
}
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	return &File{fd, name}
}
func NewFile1(fd int, name string) *file1 {
	if fd < 0 {
		return nil
	}
	return &file1{fd, name}
}
