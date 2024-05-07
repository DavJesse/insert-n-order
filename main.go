package main

import (
	"os"
)

func main() {
	var errSlc []string
	help := []string{
		"--insert",
		"  -i",
		"\t This flag inserts the string into the string passed as argument.",
		"--order",
		"  -o",
		"\t This flag will behave like a boolean, if it is called it will order the argument.",
	}

	formatErr := "Cannot compute error ; wrong format"

	args := os.Args[1:]
	if len(args) == 0 {
		printSlice(help)
		return
	}
	var inStr, argStr, rsltStr string
	var signal1, signal2, signal3 int

	for _, arg := range args {
		if stringContains(arg, "-i") {
			if len(arg) < len("-i") || len(arg) < len("--insert") {
				Print(arg)
				errSlc = stringSplit(formatErr, "error ")
				print(errSlc[0] + "\"" + arg + "\"" + errSlc[1] + "\n")
				return
			} else {
				if stringPrefix(arg, "-i") || stringPrefix(arg, "--insert") {
					if stringContains(arg, "=") {
						inStr += stringSplit(arg, "=")[1]
						signal1 += 3
					} else {
						errSlc = stringSplit(formatErr, "error ")
						print(errSlc[0] + "\"" + arg + "\"" + errSlc[1] + "\n")
						return
					}
				}
			}
		} else if stringContains(arg, "-o") {
			if len(arg) < len("-o") || len(arg) < len("--order") {
				errSlc = stringSplit(formatErr, "error ")
				print(errSlc[0] + "\"" + arg + "\"" + errSlc[1] + "\n")
				return
			} else {
				if stringPrefix(arg, "-o") || stringPrefix(arg, "--order") {
					signal2 = 5
				}
			}
		} else if stringContains(arg, "-h") {
			if arg == "-h" || arg == "--help" {
				printSlice(help)
				return
			}
		} else {
			argStr += arg
			signal3 = 7
		}
	}
	switch signal1 > 0 || signal3 > 0 {
	case ((signal1+signal3)-7)%3 == 0:
		rsltStr = argStr + inStr

	case signal1%3 == 0:
		rsltStr = inStr

	case signal3%7 == 0:
		rsltStr += argStr

	case signal1+signal2+signal3 == 7:
		rsltStr = argStr
	}

	if signal2 > 0 && signal2%5 == 0 {
		rsltStr = stringSort(rsltStr)
	}

	Print(rsltStr)
}
