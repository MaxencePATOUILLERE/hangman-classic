package main

import "fmt"

func printASCIIArt(data HangManData) {
	_PrintStandard := []string{"           ", "           ", "           ", "           ", "           ", "  ______   ", " |______|  "}
	aPrintStandard := []string{"           ", "    /\\     ", "   /  \\    ", "  / /\\ \\   ", " / ____ \\  ", "/_/    \\_\\ ", "           ", "           "}
	bPrintStandard := []string{"  ____     ", " |  _ \\    ", " | |_) |   ", " |  _ <    ", " | |_) |   ", " |____/    ", "           ", "           "}
	cPrintStandard := []string{"   _____   ", "  / ____|  ", " | |       ", " | |       ", " | |       ", "  \\_____|  ", "           ", "           "}
	dPrintStandard := []string{"  _____    ", " |  __ \\   ", " | |  | |  ", " | |  | |  ", " | |__| |  ", " |_____/   ", "           ", "           "}
	ePrintStandard := []string{"  ______   ", " |  ____|  ", " | |__     ", " |  __|    ", " | |____   ", " |______|  ", "           ", "           "}
	fPrintStandard := []string{"  ______   ", " |  ____|  ", " | |__     ", " |  __|    ", " | |       ", " |_|       ", "           ", "           "}
	hPrintStandard := []string{"  _    _   ", " | |  | |  ", " | |__| |  ", " |  __  |  ", " | |  | |  ", " |_|  |_|  ", "           ", "           "}
	iPrintStandard := []string{"  _____    ", " |_   _|   ", "   | |     ", "   | |     ", "   | |     ", " |_____|   ", "           ", "           "}
	jPrintStandard := []string{"       _   ", "      | |  ", "      | |  ", "  _   | |  ", " | |__| |  ", "  \\____/   ", "           ", "           "}
	kPrintStandard := []string{"  _  __    ", " | |/ /    ", " | ' /     ", " |  <      ", " | . \\     ", " |_|\\_\\    ", "           ", "           "}
	lPrintStandard := []string{"  _        ", " | |       ", " | |       ", " | |       ", " | |____   ", " |______|  ", "           ", "           "}
	mPrintStandard := []string{"  __  __   ", " |  \\/  |  ", " | \\  / |  ", " | |\\/| |  ", " | |  | |  ", " |_|  |_|  ", "           ", "           "}
	nPrintStandard := []string{"  _   _    ", " | \\ | |   ", " |  \\| |   ", " | . ` |   ", " | |\\  |   ", " |_| \\_|   ", "           ", "           "}
	oPrintStandard := []string{"   ____    ", "  / __ \\   ", " | |  | |  ", " | |  | |  ", " | |__| |  ", "  \\____/   ", "           ", "           "}
	pPrintStandard := []string{"  _____    ", " |  __ \\   ", " | |__) |  ", " |  ___/   ", " | |       ", " |_|       ", "           ", "           "}
	qPrintStandard := []string{"   ____    ", "  / __ \\   ", " | |  | |  ", " | |  | |  ", " | |__| |  ", "  \\___\\_\\  ", "           ", "           "}
	rPrintStandard := []string{"  _____    ", " |  __ \\   ", " | |__) |  ", " |  _  /   ", " | | \\ \\   ", " |_|  \\_\\  ", "           ", "           "}
	sPrintStandard := []string{"   _____   ", "  / ____|  ", " | (___    ", "  \\___ \\   ", "  ____) |  ", " |_____/   ", "           ", "           "}
	tPrintStandard := []string{"  _______  ", " |__   __| ", "    | |    ", "    | |    ", "    | |    ", "    |_|    ", "           ", "           "}
	uPrintStandard := []string{"  _    _   ", " | |  | |  ", " | |  | |  ", " | |  | |  ", " | |__| |  ", "  \\____/   ", "           ", "           "}
	vPrintStandard := []string{" __      __", " \\ \\    / /", "  \\ \\  / / ", "   \\ \\/ /  ", "    \\  /   ", "     \\/    ", "           ", "           "}
	wPrintStandard := []string{" __          __", " \\ \\        / /", "  \\ \\  /\\  / / ", "   \\ \\/  \\/ /  ", "    \\  /\\  /   ", "     \\/  \\/    ", "           ", "           "}
	for j := 0; j < 7; j++ {
		line := ""
		for i := 0; i < len(data.word); i++ {
			switch {
			case data.word[i] == '_':
				line += _PrintStandard[j]
			case data.word[i] == 'a':
				line += aPrintStandard[j]
			case data.word[i] == 'b':
				line += bPrintStandard[j]
			case data.word[i] == 'c':
				line += cPrintStandard[j]
			case data.word[i] == 'd':
				line += dPrintStandard[j]
			case data.word[i] == 'e':
				line += ePrintStandard[j]
			case data.word[i] == 'f':
				line += fPrintStandard[j]
			case data.word[i] == 'h':
				line += hPrintStandard[j]
			case data.word[i] == 'i':
				line += iPrintStandard[j]
			case data.word[i] == 'j':
				line += jPrintStandard[j]
			case data.word[i] == 'k':
				line += kPrintStandard[j]
			case data.word[i] == 'l':
				line += lPrintStandard[j]
			case data.word[i] == 'm':
				line += mPrintStandard[j]
			case data.word[i] == 'n':
				line += nPrintStandard[j]
			case data.word[i] == 'o':
				line += oPrintStandard[j]
			case data.word[i] == 'p':
				line += pPrintStandard[j]
			case data.word[i] == 'q':
				line += qPrintStandard[j]
			case data.word[i] == 'r':
				line += rPrintStandard[j]
			case data.word[i] == 's':
				line += sPrintStandard[j]
			case data.word[i] == 't':
				line += tPrintStandard[j]
			case data.word[i] == 'u':
				line += uPrintStandard[j]
			case data.word[i] == 'v':
				line += vPrintStandard[j]
			case data.word[i] == 'w':
				line += wPrintStandard[j]
			}
		}
		fmt.Println(line)
	}
}
