package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("/home/alexander/.bash_history")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
	strings.Split(content, "\n")
}
