package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

//func (this *P) GobEncode() ([]byte, error) {
//	return []byte{}, nil
//}
func (this *P) GobEncode() ([]byte, error) {
	return []byte{}, nil
}

type Q struct {
	X, Y *int32
	Name string
}

func main() {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	//Encode(send)the value
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	//Decode(receive)the value
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Println(q)
	fmt.Printf("%s: {%d,%d}\n", q.Name, *q.X, *q.Y)
}
