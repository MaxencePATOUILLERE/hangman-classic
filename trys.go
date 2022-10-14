package main

func trys(data HangManData, testLetter string) HangManData {
	if isUsed(data, testLetter) {
		return data
	}
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
	printASCIIArt(data, "thinkertoy")
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
	if data.word == data.toFind {
		return true
	}
	return false
}

func isUsed(data HangManData, letter string) bool {
	for i := 0; i < len(data.used); i++ {
		if data.used[i] == rune(letter[0]) {
			return true
		}
	}
	return false
}
