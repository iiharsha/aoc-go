package main

import (
	"fmt"
	"github.com/wolfeidau/stringtokenizer"
	"log"
	"os"
	"strings"
)

var count int

func main() {
	file, err := os.ReadFile("day4.test")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		testtok(line)
	}
	fmt.Println(count)

}

// horizontal, vertical, diagonal, overlapping
func testtok(line string) {
	input := line
	tokenizer := stringtokenizer.NewStringTokenizer(strings.NewReader(input), "\n", false)

	for tokenizer.HasMoreTokens() {
		token := tokenizer.NextToken()
		fmt.Println(token)
	}
}

func horizontalCheck(line string) {
	horizontalFwd := strings.Contains(line, "XMAS")
	if horizontalFwd {
		count++
		fmt.Println(line)
	}
}
