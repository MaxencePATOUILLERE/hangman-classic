package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	Word     string
	ToFind   string
	Attempts int
	Used     []rune
}

func main() {
	var GameData HangManData
	args := os.Args
	if len(args) > 2 {
		jsonFile, err := os.Open("save.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &GameData)
	} else {
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
		GameData = HangManData{
			Word:     hidden,
			ToFind:   word,
			Attempts: 0,
		}
		GameData = reveal(GameData)
	}
	fmt.Println(GameData.Word)
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
			} else if isGood(data.ToFind, string(letter)) {
				newData := trys(data, letter)
				data.Word = newData.Word
			} else {
				data.Attempts++
				printHangMan(data.Attempts)
			}
			data.Used = append(data.Used, rune(letter[0]))
		} else {
			fmt.Println("Bad input")
		}
	}
	if data.Attempts == 10 {
		print("You failed the word was : " + data.ToFind)
		return
	}
	print("Congrats !")
}

func reveal(data HangManData) HangManData {
	var cpt = 0
	for cpt < len(data.ToFind)/2-1 {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		letter := rune(data.ToFind[r.Intn(len(data.ToFind))])
		if isNotIn(letter, data.Used) {
			cpt++
			data = tryWithoutOut(string(letter), data)
		}
		data.Used = append(data.Used, letter)
	}
	return data
}

func tryWithoutOut(testLetter string, data HangManData) HangManData {
	listemystery := []string{}
	editedToFind := ""
	for i := 0; i < len(data.ToFind); i++ {
		editedToFind = editedToFind + string(data.ToFind[i]) + " "
	}
	Index := findIndexLetter(testLetter, editedToFind)
	for i := 0; i < len(data.Word); i++ {
		listemystery = append(listemystery, string(data.Word[i]))
	}
	for i := 0; i < len(Index); i++ {
		listemystery[Index[i]] = testLetter
	}
	data.Word = convertInStr(listemystery)
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
