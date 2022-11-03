package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func printASCIIArt(data HangManData) {
	file := openFile(data.WhichAsciiArt)
	for j := 0; j < 8; j++ {
		line := []string{}
		for i := 0; i < len(data.Word); i++ {
			if data.Word[i] == '_' {
				line = append(line, file[116+j])
			} else if data.Word[i] == ' ' {
			} else {
				line = append(line, file[586+j+int(rune(data.Word[i]-97)*9)])
			}
		}
		finalWord := ""
		for i := 0; i < len(line); i++ {
			finalWord = finalWord + line[i]
		}
		fmt.Println(finalWord)
	}
}

func openFile(letterFile string) []string {
	asciiType := whichTypeOfAsciiArt(letterFile)
	scanner := bufio.NewScanner(asciiType)
	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

func whichTypeOfAsciiArt(letterFile string) io.Reader {
	var f *os.File
	if letterFile == "standard.txt" {
		f, _ = os.Open("standard.txt")
	} else if letterFile == "shadow.txt" {
		f, _ = os.Open("shadow.txt")
	} else if letterFile == "thinkertoy.txt" {
		f, _ = os.Open("thinkertoy.txt")
	}
	return f
}
