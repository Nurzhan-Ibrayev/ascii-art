package Modify

import (
	"fmt"
	"strings"
)

func PrintLetter(s string, arr [][]string) {
	word := strings.ReplaceAll(s, "\n", "\\n")
	splitedWord := strings.Split(word, "\\n")
	splitedWord = checkLenWord(splitedWord)
	for _, w := range splitedWord {
		if w == "" {
			fmt.Println()
			// continue
		} else {
			for i := 0; i < 8; i++ {
				for _, w2 := range w {
					if w2 >= 32 {
						fmt.Print(arr[w2-32][i])
					}
				}
				fmt.Println()
			}
		}
	}
}

func checkLenWord(words []string) []string {
	counter := 0
	for _, word := range words {
		if word == "" {
			counter++
		}
	}
	if counter == len(words) {
		return words[1:]
	}
	return words
}
