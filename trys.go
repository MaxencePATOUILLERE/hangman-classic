package main

import "fmt"

func trys(data HangManData, testLetter string) HangManData {
	if isUsed(data, testLetter) {
		return data
	}
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
	fmt.Println(data.Word)
	return data
}
func findIndexLetter(TestLetter string, Words string) []int {
	Index := []int{}
	for i := 0; i < len(Words); i++ {
		if TestLetter[0] == Words[i] {
			Index = append(Index, i)
		}
	}
	return Index
}
func convertInStr(liste []string) string {
	finalWord := ""
	for i := 0; i < len(liste); i++ {
		finalWord = finalWord + liste[i]
	}
	return finalWord
}
func isGood(str string, test string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == test[0] {
			return true
		}
	}
	return false
}
func finish(data HangManData) bool {
	word := ""
	for i := 0; i < len(data.Word); i++ {
		if i < len(word)-3 && string(data.Word[i]) == " " && string(data.Word[i+1]) == " " && string(data.Word[i+2]) == " " {
			word += " "
		}
		if string(data.Word[i]) != " " && string(data.Word[i]) != "_" {
			word += string(data.Word[i])
		}
	}
	if word == data.ToFind {
		return true
	}
	if data.Attempts == 10 {
		return true
	}
	return false
}
func isUsed(data HangManData, letter string) bool {
	for i := 0; i < len(data.Used); i++ {
		if data.Used[i] == rune(letter[0]) {
			return true
		}
	}
	return false
}
