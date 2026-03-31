package main

import "fmt"

func caesarCipher(s string, shift int) string {
	shift = shift % 26
	result := ""

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			newChar := 'a' + (ch-'a'+rune(shift))%26
			result += string(newChar)
		} else if ch >= 'A' && ch <= 'Z' {
			newChar := 'A' + (ch-'A'+rune(shift))%26
			result += string(newChar)

			// Non-alphabet characters (space, symbols)
		} else {
			result += string(ch)
		}
	}

	return result
}

func main() {
	fmt.Println(caesarCipher("abc", 2))        // cde
	fmt.Println(caesarCipher("xyz", 2))        // zab
	fmt.Println(caesarCipher("Hello, Go!", 3)) // Khoor, Jr!
}
