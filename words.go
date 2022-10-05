package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func getFileWords(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	text := string(file)
	fmt.Println(text)
	return ""
}
