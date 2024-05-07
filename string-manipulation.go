package main

// Split string at the event of the separator
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

// Sort string characters in ASCII order (least to most)
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

// Check whether a string contains the specified substring
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

// Check whether a string starts with the specified substring
func stringPrefix(str, subStr string) bool {
	var status bool

	if str[:len(subStr)] == subStr {
		status = true
	}
	return status
}
