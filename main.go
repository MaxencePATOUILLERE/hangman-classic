package main

import (
	"fmt"
)

type HangManData struct {
	word     string // Word composed of '_', ex: H_ll_
	toFind   string // Final word chosen by the program at the beginning. It is the word to find
	attempts int    // Number of attempts left
}

func main() {
	word := formatWord(getFileWords("words.txt"))
	hidden := ""
	for i := 0; i < len(word); i++ {
		if word[i] == ' ' {
			hidden += " "
		} else {
			hidden += "_"
		}
	}
	GameData := HangManData{
		word:     hidden,
		toFind:   word,
		attempts: 0,
	}
	game(GameData)
}

func game(data HangManData) {
	var letter string
	for i := 0; i < 10; i++ {
		fmt.Scanln(&letter)
		//s_letter := []rune(letter)
		if isGood(data.toFind, string(letter)) {
			trys(data, data.toFind)
		} else {
			data.attempts++
			printHangMan(data.attempts)
		}
	}
}
