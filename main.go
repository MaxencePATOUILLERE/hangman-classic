package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type HangManData struct {
	Save     string
	File     string
	Word     string
	ToFind   string
	Attempts int
	Used     []rune
}

func main() {
	var wordList string
	var saveFile string
	flag.StringVar(&wordList, "wordlist", "words.txt", "Wordlist to the hangman")
	flag.StringVar(&saveFile, "startWith", "", "Save file for the hangman")
	flag.Parse()
	if saveFile != "" && isFileValid(saveFile) {
		GameData := getFileData(&saveFile)
		game(GameData)
	} else {
		printStart()
		fmt.Println("No valid file given, start a new game ... ")
		setup(wordList)
	}
}

func setup(wl string) {
	word := formatWord(getFileWords(wl))
	if word == "" {
		fmt.Println("Invalid File : " + wl + "\nSupported files are json and txt")
		return
	}
	GameData := HangManData{
		Save:     "",
		File:     wl,
		Word:     genHidden(word),
		ToFind:   word,
		Attempts: 0,
	}
	GameData = reveal(GameData)
	printWord(GameData.Word)
	game(GameData)
}

func game(data HangManData) {
	var letterIn string
	for !finish(data) {
		fmt.Print("Choose: ")
		fmt.Scan(&letterIn)
		if letterIn == "STOP" || letterIn == "stop" || letterIn == "Stop" {
			if data.Save != "" {
				saveWithPath(data, data.Save)
				return
			}
			good := true
			for good {
				good = !save(data)
			}
			return
		}
		letter := rune(letterIn[0])
		fmt.Print("\033[H\033[2J")
		data = checkInput(data, letter)
		printWord(data.Word)
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
		if !isUsed(data, letter) {
			cpt++
			data = trys(data, letter)
		}
	}
	return data
}

func isValid(l rune) bool {
	if l >= 'A' && l <= 'Z' || l >= 'a' && l <= 'z' {
		return true
	}
	return false
}

func finish(data HangManData) bool {
	if data.Word == data.ToFind {
		return true
	} else if data.Attempts == 10 {
		return true
	}
	return false
}

func checkInput(data HangManData, l rune) HangManData {
	if l != ' ' && isValid(l) {
		if isUsed(data, l) {
			printHangMan(data.Attempts)
			fmt.Println("Letter already used.")
			return data
		} else if isGood(data.ToFind, l) {
			data = trys(data, l)

		} else {
			data.Attempts++
			data.Used = append(data.Used, l)
			printHangMan(data.Attempts)
			fmt.Println("Not present in the word,", 10-data.Attempts, "attempts remaining")
			return data
		}
	} else {
		fmt.Println("Bad input")

	}
	printHangMan(data.Attempts)
	return data
}
