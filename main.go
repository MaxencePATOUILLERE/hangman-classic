package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
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
	var filePath *string
	if len(os.Args) > 3 && isAlreadyExist(os.Args[3]) && isFileValid(os.Args[3]) {
		filePath = flag.String("startWith", os.Args[3], "File to read to start")
		GameData = getFileData(filePath)
	} else {
		GameData = createGameData(os.Args)
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
				saveWithPath(data, path)
				return
			}
			good := true
			for good {
				good = !save(data)
			}
			return
		} else {
			fmt.Println("Bad input")
		}
	}
	if data.Attempts == 10 {
		print("You failed the word was : " + data.ToFind)
		return
	}
	if path != nil {
		saveWithPath(data, path)
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
	var listemystery []string
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

func save(data HangManData) bool {
	var name string
	fmt.Print("Choose a valid filename to save : ")
	fmt.Scanln(&name)
	if !isAlreadyExist(name) {
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
	_ = os.WriteFile(name+".txt", content, 0644)
	fmt.Println("Game Saved in " + name + ".txt")
	return true
}

func isAlreadyExist(path string) bool {
	_, err := os.Stat(path + ".txt")
	if err != nil {
		return true
	}
	return false
}

func saveWithPath(data HangManData, path *string) {
	if path != nil {
		content, _ := json.MarshalIndent(data, "", " ")
		if isAlreadyExist(*path) || isAlreadyExist(*path+".txt") {
			_ = os.Remove(*path)
			_ = os.WriteFile(*path, content, 0644)
		} else {
			_ = os.Remove(*path + ".txt")
			_ = os.WriteFile(*path+".txt", content, 0644)
		}
		fmt.Println("Game Saved in " + *path)
		return
	}
}

func createGameData(args []string) HangManData {
	var GameData HangManData
	if !isFileValid(args[3]) {
		panic("Bad argument")
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
	GameData = HangManData{
		Word:     hidden,
		ToFind:   word,
		Attempts: 0,
	}
	GameData = reveal(GameData)
	return GameData
}

func getFileData(file *string) HangManData {
	var GameData HangManData
	jsonFile, err := os.Open(*file)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &GameData)
	fmt.Println("Welcome Back, you have " + string(rune(GameData.Attempts)+48) + " attempts remaining.")
	return GameData
}

func isFileValid(file string) bool {
	fileinfo, _ := os.Stat(file)
	fileContent, _ := os.Open(file)
	scanner := bufio.NewScanner(fileContent)
	var result []string
	cpt := 0
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
		cpt++
	}
	if fileinfo.Size() == 0 {
		return false
	}
	return true
}
