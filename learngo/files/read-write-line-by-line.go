package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filepath := "../extras/rabbits.txt"

	f, _ := os.Open(filepath)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("../extras/output2.txt")
	if err != nil {
		log.Fatalln("Error Creating file:", err)
	}
	defer f.Close()

	for _, str := range []string{"apple", "banana", "carrot"} {
		bytes, err := f.WriteString(str)
		if err != nil {
			log.Fatalln("Error writing string:", err)
		}
		fmt.Printf("Wrote %d bytes to file\n", bytes)
	}
}
