package main

import (
	"encoding/json"
	"flag"
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
	var filePath *string
	if len(os.Args) > 2 {
		filePath = flag.String("startWith", os.Args[3], "File to read to start")
	}
	var GameData HangManData
	args := os.Args
	if filePath != nil {
		jsonFile, err := os.Open(*filePath)
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &GameData)
		fmt.Println("Welcome Back, you have " + string(rune(GameData.Attempts)+48) + " attempts remaining.")
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
	game(GameData, filePath)
}

func game(data HangManData, path *string) {
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
		} else if letter == "STOP" || letter == "stop" || letter == "Stop" {
			if path != nil {
				content, _ := json.MarshalIndent(data, "", " ")
				_, err := os.Stat(*path)
				if err == nil {
					os.Remove(*path)
					_ = ioutil.WriteFile(*path, content, 0644)
				} else {
					os.Remove(*path + ".txt")
					_ = ioutil.WriteFile(*path+".txt", content, 0644)
				}
				fmt.Println("Game Saved in " + *path)
				return
			}
			var name string
			goodName := false
			for !goodName {
				fmt.Print("Choose a filename to save : ")
				fmt.Scanln(&name)
				goodName = save(data, name)
			}
			fmt.Println("Game Saved in " + name + ".txt")
			return
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
			data.Used = append(data.Used, letter)
		}
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

func save(data HangManData, name string) bool {
	isExist := false
	_, err := os.Stat(name + ".txt")
	if err != nil {
		isExist = true
	}
	_, err = os.Stat(name)
	if err != nil {
		isExist = true
	}
	if isExist {
		var choice string
		fmt.Print("File already exist overwrite it ? (Y/N) ")
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" || choice == "Yes" || choice == "yes" {
			os.Remove(name + ".txt")
		} else {
			return false
		}
	}
	content, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(name+".txt", content, 0644)
	return true
}
