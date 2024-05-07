package main

import "github.com/01-edu/z01"

func Print(str string) {
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}

func printSlice(slc []string) {
	for _, ch := range slc {
		Print(ch)
	}
}
