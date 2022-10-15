package main

import (
	"fmt"
	"math/rand"
	"time"
)

type HangManData struct {
	word      string // Word composed of '_', ex: H_ll_
	toFind    string // Final word chosen by the program at the beginning. It is the word to find
	attempts  int    // Number of attempts left
	asciiType string // Which type of ascii you want
	used      []rune
}

func main() {
	asciiType := ""
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
	fmt.Print("choose ascii type (thinkertoy, standard, shadow): ")
	fmt.Scanln(&asciiType)
	GameData := HangManData{
		word:      hidden,
		toFind:    word,
		attempts:  0,
		asciiType: asciiType,
	}
	GameData = reveal(GameData)
	game(GameData)
}
func game(data HangManData) {
	var letter string
	for !finish(data) {
		fmt.Print("Choose: ")
		fmt.Scanln(&letter)
		if len(letter) == 1 && letter[0] != ' ' && isValid(rune(letter[0])) {
			if isUsed(data, letter) {
				fmt.Println("Letter already used.")
			} else if isGood(data.toFind, string(letter)) {
				newData := trys(data, letter)
				data.word = newData.word
			} else {
				data.attempts++
				printHangMan(data.attempts)
			}
			data.used = append(data.used, rune(letter[0]))
		} else {
			fmt.Println("Bad input")
		}
	}
	if data.attempts == 10 {
		print("You failed the word was : " + data.toFind)
		return
	}
	print("Congrats !")
}

func reveal(data HangManData) HangManData {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	letter := rune(data.toFind[r.Intn(len(data.toFind))])
	data = trys(data, string(letter))
	data.used = append(data.used, letter)
	return data
}

func isValid(l rune) bool {
	if l >= 'A' && l <= 'Z' || l >= 'a' && l <= 'z' {
		return true
	}
	return false
}
