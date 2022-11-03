package main

import (
	"bufio"
	"fmt"
	"os"
)

func printHangMan(failAttempts int) {
	hangMan := openHangManTxt()
	if failAttempts == 0 {
		for k := 0; k < 8; k++ {
			fmt.Println()
		}
	} else {
		for i := 0; i < 8; i++ {
			fmt.Println(hangMan[i+8*(failAttempts-1)])
		}
	}
}

func printWord(w string, letterFile string) {
	if letterFile != "" {
		printASCIIArt(w, letterFile)
	} else {
		for i := 0; i < len(w); i++ {
			fmt.Print(string(w[i]))
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func printStart() {
	fmt.Println("------------------------------")
	fmt.Println("      Welcome to Hangman      ")
	fmt.Println("------------------------------")
}

func openHangManTxt() []string {
	var f *os.File
	f, _ = os.Open("hangman.txt")
	scanner := bufio.NewScanner(f)
	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}
