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
	fmt.Printf("%s", content)
	var content_string string
	content_string = string(content)
	var string_list []string
	var result_list []string
	var result_map map[string]int
	for _, i := range content_string {
		string_list = strings.Split(string(i), "\n")
	}
	for i, _ := range string_list {
		result_list = strings.Split(string(i), " ")
		for j, _ := range result_list {
			if j[0] != "sudo" {
				result_map[j[0]] += 1
			} else {
				result_map[j[1]] += 1
			}
		}
	}
	fmt.Sprintf("%s", string_list)
}
