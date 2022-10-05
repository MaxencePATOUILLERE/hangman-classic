package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println(getFileWords("words.txt"))
}

func getFileWords(path string) string {
	f, _ := os.Open(path)
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
