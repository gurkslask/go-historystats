package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("/home/alex/.bash_history")
	if err != nil {
		log.Fatal(err)
	}
	var content_string []string
	result_map := make(map[string]int)
	content_string = strings.Split(string(content), "\n")
	for _, text := range content_string {
		temp_string := strings.Split(text, " ")
		if temp_string[0] != "sudo" {
			result_map[temp_string[0]] += 1
		} else {
			result_map[temp_string[1]] += 1
		}
	}
	for key, val := range result_map {
		fmt.Printf("%s: %d \n", key, val)
	}
}
