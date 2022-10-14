package main

import (
	"fmt"
)

func printHangMan(failAttempts int) {

	fmt.Println("Not present in the word,", 10-failAttempts, "attempts remaining")
}
