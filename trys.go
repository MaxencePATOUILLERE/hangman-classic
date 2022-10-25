package main

func trys(data HangManData, testLetter rune, print bool) HangManData {
	if testLetter >= 'A' && testLetter <= 'Z' {
		testLetter += 32
	}
	Index := findIndexLetter(testLetter, data.toFind)
	addedLetter := []rune(data.word)
	for i := 0; i < len(data.word); i++ {
		for j := 0; j < len(Index); j++ {
			if i == Index[j] {
				addedLetter[i] = testLetter
			}
		}
	}
	data.word = convertInStr(addedLetter)
	data.used = append(data.used, testLetter)
	if print {
		printWord(data.word)
	}
	return data
}

func findIndexLetter(TestLetter rune, Words string) []int {
	Index := []int{}
	for i := 0; i < len(Words); i++ {
		if TestLetter == rune(Words[i]) {
			Index = append(Index, i)
		}
	}
	return Index
}

func convertInStr(liste []rune) string {
	finalWord := ""
	for i := 0; i < len(liste); i++ {
		finalWord = finalWord + string(liste[i])
	}
	return finalWord
}

func isGood(str string, test rune) bool {
	for i := 0; i < len(str); i++ {
		if rune(str[i]) == test || rune(str[i]) == test+32 {
			return true
		}
	}
	return false
}

func isUsed(data HangManData, letter rune) bool {
	for i := 0; i < len(data.used); i++ {
		if data.used[i] == letter {
			return true
		}
	}
	return false
}
