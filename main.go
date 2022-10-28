package main

import (
	"fmt"
	"math/rand"
	"os"
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
	args := os.Args
	if args[2] != "--letterFile" && len(args) != 2 || len(args) != 4 {
		fmt.Println("Bad Parameter")
		return
	}
	word := formatWord(getFileWords(args[1]))
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
		asciiType: "",
		word:      hidden,
		toFind:    word,
		attempts:  0,
	}
	if len(args) == 4 {
		GameData = HangManData{
			asciiType: os.Args[3],
			word:      hidden,
			toFind:    word,
			attempts:  0,
		}
		GameData = reveal(GameData)
		printASCIIArt(GameData)
	} else {
		GameData = reveal(GameData)
		fmt.Println(GameData.word)
	}
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
				if data.asciiType == "" {
					fmt.Println(data.word)
				} else {
					printASCIIArt(data)
				}
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
	return
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
