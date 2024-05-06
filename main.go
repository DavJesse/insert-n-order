package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		help := []string{"--insert", "  -i", "\tThis flag inserts the string into the string passed as argument.", "--order", "  -o", "\tThis flag will behave like a boolean, if it is called it will order the argument."}
		for _, ch := range help {
			PrintLn(ch)
		}
	}
	var inStr, argStr, rsltStr string
	var signal1, signal2 int
	var count int

	for _, arg := range args {
		if stringContains(arg, "-i") {
			signal1 += 3
			if arg[:2] == "-i" || arg[:8] == "--insert" {
				inStr += stringSplit(arg, "=")[1]
			}
		} else if stringContains(arg, "-o") {
			if arg[:2] == "-o" || arg == "--order" {
				signal2 = 5
			}
		} else if stringContains(arg, "-h") {
			if arg == "-h" || arg == "--help" {
				help := []string{"--insert", "  -i", "\tThis flag inserts the string into the string passed as argument.", "--order", "  -o", "\tThis flag will behave like a boolean, if it is called it will order the argument."}
				for _, ch := range help {
					PrintLn(ch)
				}
			}
		} else {
			argStr = arg
		}
	}

	if signal1%3 == 0 {
		count = signal1 / 3
		rsltStr = argStr + inStr

	} else if ((signal1+signal2)-5)%3 == 0 {
		count = ((signal1 + signal2) - 5) / 3

		rsltStr = argStr + inStr

		for count-1 > 0 {
			rsltStr += inStr
			count--
		}
		rsltStr = stringSort(rsltStr)
	} else if signal1+signal2 == 5 {
		rsltStr = stringSort(argStr)
	} else {
		rsltStr = argStr
	}
	PrintLn(rsltStr)
}

func PrintLn(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func stringContains(str, subStr string) bool {
	var status bool

	for i := 0; i < len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			if subStr == str[i:j] {
				status = true
			}
		}
	}
	return status
}

func stringSplit(str, sep string) []string {
	var strSlc []string

	for i := 0; i < len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			if sep == str[i:j] {
				strSlc = append(strSlc, str[:i])
				strSlc = append(strSlc, str[j:])
			}
		}
	}
	return strSlc
}

func stringSort(str string) string {
	rnd := []rune(str)

	for i := 0; i < len(rnd)-1; i++ {
		for j := 0; j < len(rnd)-1-i; j++ {
			if rnd[j] > rnd[j+1] {
				rnd[j], rnd[j+1] = rnd[j+1], rnd[j]
			}
		}
	}
	result := string(rnd)
	return result
}
