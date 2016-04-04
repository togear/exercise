package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	filename := "../extras/rabbits.txt"

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file", filename)
	}

	fmt.Println(string(content))

	outfile := "../extras/output.txt"
	err = ioutil.WriteFile(outfile, content, 0644)
	if err != nil {
		log.Fatalln("Error writing file:", err)
	}
}
