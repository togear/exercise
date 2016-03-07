package main

import "fmt"
import "bytes"

func main() {
	fmt.Println("%T\n",'a')
	var c1, c2 rune = '\u6211', '们'
	fmt.Println(c1 == '我',string(c2) == "\xe4\xbb\xac")

	a := "我爱go编程"
	for i:=0 ;i < len(a); i++ {
		fmt.Println("%c",a[i])
	}
	fmt.Println("\n")

	b := []rune(a)
	for i:=0 ;i < len(b); i++ {
		fmt.Println("%c",b[i])
	}
	fmt.Println("\n")

	for _,m := range a {
		fmt.Println("%c",m)
	}
	
	d := "你"
	e := d + "好"
	fmt.Println(e)

	c := bytes.Buffer{}
	c.WriteString("你")
	c.WriteString("好")
	fmt.Println(c.String())


}
