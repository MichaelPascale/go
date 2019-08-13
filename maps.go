package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var wordbin map[string]int
	wordbin = make(map[string]int)

	words := strings.Fields(s)

	for _, word := range words {
		wordbin[word]++
	}

	return wordbin
}

func main() {
	wc.Test(WordCount)
}
