package main

import (
	"fmt"
	"math/rand"
	"time"
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
	GameData = reveal(GameData)
	fmt.Println(GameData.word)
	game(GameData)
}

func game(data HangManData) {
	var letter string
	for !finish(data) {
		fmt.Print("Choose: ")
		fmt.Scanln(&letter)
		if len(letter) > 1 {
			var response string
			fmt.Println("You have entered more than one letter want to try to guess the word ? (Y/N)")
			fmt.Scanln(&response)
			if response == "Y" || response == "Yes" || response == "y" || response == "yes" {
				if letter == data.toFind {
					fmt.Println("Congrats !")
					return
				} else {
					data.attempts += 2
					fmt.Println("Bad word try again ! " + string(rune(-data.attempts+58)) + " left.")
				}
			} else if response != "N" && response != "No" && response != "n" && response != "no" {
				fmt.Println("Bad input try again")
			}
		} else if letter[0] != ' ' && isValid(rune(letter[0])) {
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
	var cpt = 0
	for cpt < len(data.toFind)/2-1 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		letter := rune(data.toFind[r.Intn(len(data.toFind))])
		if isNotIn(letter, data.used) {
			cpt++
			data = tryWithoutOut(string(letter), data)
		}
		data.used = append(data.used, letter)
	}
	return data
}
func tryWithoutOut(testLetter string, data HangManData) HangManData {
	listemystery := []string{}
	editedToFind := ""
	for i := 0; i < len(data.toFind); i++ {
		editedToFind = editedToFind + string(data.toFind[i]) + " "
	}
	Index := findIndexLetter(testLetter, editedToFind)
	for i := 0; i < len(data.word); i++ {
		listemystery = append(listemystery, string(data.word[i]))
	}
	for i := 0; i < len(Index); i++ {
		listemystery[Index[i]] = testLetter
	}
	data.word = convertInStr(listemystery)
	return data
}
func isNotIn(l rune, lst []rune) bool {
	for i := 0; i < len(lst); i++ {
		if lst[i] == l {
			return false
		}
	}
	return true
}
func isValid(l rune) bool {
	if l >= 'A' && l <= 'Z' || l >= 'a' && l <= 'z' {
		return true
	}
	return false
}
