package main

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

func getFileWords(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic("Bad parameter")
	}
	scanner := bufio.NewScanner(f)
	var result []string
	cpt := 0
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
		cpt++
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return result[r.Intn(cpt)]
}

func formatWord(word string) string {
	sSlice := []rune(word)
	for i := 0; i < len(sSlice); i++ {
		if sSlice[i] < 'a' || sSlice[i] > 'z' {
			sSlice[i] = ' '
		}
	}
	return string(sSlice)
}
