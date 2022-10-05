package main

import (
	"fmt"
)

func printHangMan(failAttempts int) {
	switch {
	case failAttempts == 1:
		for i := 0; i < 7; i++ {
			fmt.Print("\n")
		}
		fmt.Println("=========")
	case failAttempts == 2:
		for i := 0; i < 5; i++ {
			fmt.Println("      |")
		}
		fmt.Println("=========")
	case failAttempts == 3:
		fmt.Println("  +---+")
		for i := 0; i < 5; i++ {
			fmt.Println("      |")
		}
		fmt.Println("=========")
	case failAttempts == 4:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		for i := 0; i < 4; i++ {
			fmt.Println("      |")
		}
		fmt.Println("=========")
	case failAttempts == 5:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		for i := 0; i < 3; i++ {
			fmt.Println("      |")
		}
		fmt.Println("=========")
	case failAttempts == 6:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println("  |   |")
		for i := 0; i < 2; i++ {
			fmt.Println("      |")
		}
		fmt.Println("=========")
	case failAttempts == 7:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|   |")
		for i := 0; i < 2; i++ {
			fmt.Println("      |")
		}
		fmt.Println("=========")
	case failAttempts == 8:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|\\  |")
		for i := 0; i < 2; i++ {
			fmt.Println("      |")
		}
		fmt.Println("=========")
	case failAttempts == 9:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|\\  |")
		fmt.Println(" /    |")
		for i := 0; i < 1; i++ {
			fmt.Println("      |")
		}
		fmt.Println("=========")
	case failAttempts == 10:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|\\  |")
		fmt.Println(" / \\  |")
		for i := 0; i < 1; i++ {
			fmt.Println("      |")
		}
		fmt.Println("=========")
	}
}
