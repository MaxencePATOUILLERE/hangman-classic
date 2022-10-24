package main

import (
	"bufio"
	"fmt"
	"os"
)

func printASCIIArt(data HangManData) {
	var f *os.File
	if data.asciiType == "standard.txt" {
		f, _ = os.Open("standard.txt")
	} else if data.asciiType == "shadow.txt" {
		f, _ = os.Open("shadow.txt")
	} else if data.asciiType == "thinkertoy.txt" {
		f, _ = os.Open("thinkertoy.txt")
	}
	scanner := bufio.NewScanner(f)
	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	for j := 0; j < 7; j++ {
		line := []string{}
		for i := 0; i < len(data.word); i++ {
			if data.word[i] == '_' {
				line = append(line, result[116+j])
			} else if data.word[i] == ' ' {
			} else {
				line = append(line, result[298+j+int(rune(data.word[i]-97)*9)])
			}
		}
		fmt.Println(convertInStr(line))
	}
}
