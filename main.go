package main

import (
	"fmt"
)

type HangManData struct {
	word     string // Word composed of '_', ex: H_ll_
	toFind   string // Final word chosen by the program at the beginning. It is the word to find
	attempts int    // Number of attempts left
	used     []rune
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
		hidden += " "
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
		if finish(data) {
			fmt.Println("Congrats !")
			return
		}
		fmt.Print("Choose: ")
		fmt.Scanln(&letter)
		if len(letter) > 1 || letter[0] == ' ' {

		}
		//s_letter := []rune(letter)
		if isUsed(data, letter) {
			fmt.Println("Letter already used.")
			i--
		} else if isGood(data.toFind, string(letter)) {
			newData := trys(data, letter)
			data.word = newData.word
			i--
		} else {
			data.attempts++
			printHangMan(data.attempts)
		}
		data.used = append(data.used, rune(letter[0]))
	}
	print("You failed the word was : " + data.toFind)
	return
}
