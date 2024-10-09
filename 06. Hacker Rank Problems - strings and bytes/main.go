package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	for {
		fmt.Println("\n----MENU---")
		fmt.Println("\n1. Camel Case")
		fmt.Println("2. Caesar Cipher")
		fmt.Println("3. Exit")

		var op int
		fmt.Printf("\nChoose one of the above options:- ")
		fmt.Scanln(&op)

		switch op {

		case 1:
			fmt.Printf("\nEnter the input string for Camel Case:- ")
			var s string
			fmt.Scanln(&s)

			out := camelcase(s)

			fmt.Printf("Output :- %v", out)

		case 2:
			fmt.Printf("\nEnter the input string for Caesar Cipher:- ")
			var s string
			fmt.Scanln(&s)

			fmt.Printf("\nEnter the offset/rotation factor for Caesar Cipher:- ")
			var n int32
			fmt.Scanln(&n)

			out := caesarcipher(s, n)

			fmt.Printf("Output %v", out)

		case 3:
			os.Exit(1)

		default:
			fmt.Println("Invalid Input")

		}
	}

}

func caesarcipher(s string, k int32) string {
	var result strings.Builder
	k = k % 26

	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			result.WriteRune((char-'A'+k)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			result.WriteRune((char-'a'+k)%26 + 'a')
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func camelcase(s string) int32 {

	var wordCount int32 = 1
	for _, ch := range s {

		if unicode.IsUpper(ch) {
			wordCount++
		}
	}

	return wordCount

}
