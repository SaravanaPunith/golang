package main

import (
	"fmt"
	"strings"
)

func main() {

	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	words := strings.Fields(text)
	fmt.Println(words)
	fmt.Println(len(words))

	counts := map[string]int{}
	fmt.Println(counts)
	fmt.Println("----------")
	for _,word := range words {
		counts[strings.ToLower(word)]++
	}

	fmt.Println(counts)
}
