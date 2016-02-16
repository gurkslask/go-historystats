package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("/home/alex/.bash_history")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
}
