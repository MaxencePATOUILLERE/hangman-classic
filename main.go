package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type HangManData struct {
	word     string
	toFind   string
	attempts int
	used     []rune
}

func main() {
	var wordList string
	flag.StringVar(&wordList, "wordlist", "words.txt", "Wordlist to the hangman")
	flag.Parse()
	setup(wordList)
}

func setup(wl string) {
	word := formatWord(getFileWords(wl))
	if word == "" {
		fmt.Println("Invalid File : " + wl + "\nSupported files are json and txt")
		return
	}
	GameData := HangManData{
		word:     genHidden(word),
		toFind:   word,
		attempts: 0,
	}
	GameData = reveal(GameData)
	printWord(GameData.word)
	game(GameData)
}

func game(data HangManData) {
	var letterIn string
	for !finish(data) {
		fmt.Print("Choose: ")
		fmt.Scan(&letterIn)
		letter := rune(letterIn[0])
		data = checkInput(data, letter)
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
		if !isUsed(data, letter) {
			cpt++
			data = trys(data, letter, false)
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
	if data.word == data.toFind {
		return true
	} else if data.attempts == 10 {
		return true
	}
	return false
}

func checkInput(data HangManData, l rune) HangManData {
	if l != ' ' && isValid(l) {
		if isUsed(data, l) {
			fmt.Println("Letter already used.")
		} else if isGood(data.toFind, l) {
			data = trys(data, l, true)
		} else {
			data.attempts++
			data.used = append(data.used, l)
			printHangMan(data.attempts)
		}
	} else {
		fmt.Println("Bad input")
	}
	return data
}
