package main

import "fmt"

func trys(testLetter string, toFind string, word string) bool {
	listemystery := []string{}
	Index := findIndexLetter(testLetter, toFind)
	for i := 0; i < len(word); i++ {
		listemystery = append(listemystery, string(word[i]))
	}
	for i := 0; i < len(Index); i++ {
		listemystery[Index[i]] = testLetter
	}
	fmt.Println(convertInStr(listemystery))
	return isGood(convertInStr(listemystery), word)
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
		if str[i] != test[i] {
			return true
		}
	}
	return false
}
