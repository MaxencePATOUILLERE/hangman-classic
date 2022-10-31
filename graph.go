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
}

func printWord(w string) {
	for i := 0; i < len(w); i++ {
		fmt.Print(string(w[i]))
		fmt.Print(" ")
	}
	fmt.Println()
}

func printStart() {
	fmt.Println("------------------------------")
	fmt.Println("      Welcome to Hangman      ")
	fmt.Println("------------------------------")
}
