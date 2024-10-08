package main

import (
	"fmt"
	"unicode"
)

func camelcase(s string) int32 {

	var wordCount int32 = 1
	for _, ch := range s {

		if unicode.IsUpper(ch) {
			wordCount++
		}
	}

	return wordCount

}

func main() {

	out := camelcase("saveChangesInTheEditor")

	fmt.Println(out)
}
