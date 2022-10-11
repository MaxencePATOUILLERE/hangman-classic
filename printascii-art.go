package main

import "fmt"

func printASCIIArt(data HangManData) {

	_Print := []string{"           ", "           ", "           ", "           ", "           ", "  ______   ", " |______|  "}
	aPrint := []string{"           ", "    /\\     ", "   /  \\    ", "  / /\\ \\   ", " / ____ \\  ", "/_/    \\_\\ ", "           ", "           "}
	bPrint := []string{"  ____     ", " |  _ \\    ", " | |_) |   ", " |  _ <    ", " | |_) |   ", " |____/    ", "           ", "           "}
	for j := 0; j < 7; j++ {
		ligne := ""
		for i := 0; i < len(data.word); i++ {
			switch {
			case data.word[i] == '_':
				ligne += _Print[j]
			case data.word[i] == 'a':
				ligne += aPrint[j]
			case data.word[i] == 'b':
				ligne += bPrint[j]
			}
		}
		fmt.Println(ligne)
	}
}
