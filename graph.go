package main

import (
	"bufio"
	"fmt"
	"os"
)

func printHangMan(failAttempts int) {
	var f *os.File
	f, _ = os.Open("hangman.txt")
	scanner := bufio.NewScanner(f)
	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	for i := 0; i < 8; i++ {
		fmt.Println(result[i+8*(failAttempts-1)])
	}
	fmt.Println("Not present in the word,", 10-failAttempts, "attempts remaining")
}
