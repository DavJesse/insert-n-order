package main

import (
	"os"
)

func main() {
	var errSlc []string
	var inStr, argStr, rsltStr string
	var signal1, signal2, signal3 int
	help := []string{
		"--insert",
		"  -i",
		"\t This flag inserts the string into the string passed as argument.",
		"--order",
		"  -o",
		"\t This flag will behave like a boolean, if it is called it will order the argument.",
	}

	formatErr := "Cannot compute error ; wrong format"

	// Capture valid commandline arguments
	args := os.Args[1:]

	// Exit when user runs the program with no arguments attached
	if len(args) == 0 {
		printSlice(help)
		return
	}

	// Range over arguments, looking for flags
	for _, arg := range args {
		// Check for the "-i" and "--insert" flags
		if stringContains(arg, "-i") {
			if len(arg) < len("-i") || len(arg) < len("--insert") {
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

			// Check for the "-o" and "--order" flags
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

			// Check for the "-h" and "--help" flags
		} else if stringContains(arg, "-h") {
			if arg == "-h" || arg == "--help" {
				printSlice(help)
				return
			}

			// Identify the argument strings
		} else {
			argStr += arg
			signal3 = 7
		}
	}

	//Execute flags
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
	// Print result
	Print(rsltStr)
}
