package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Token any
type Tokens []Token

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: jlox [script")
		os.Exit(64)
	}

	var reader io.Reader
	if len(os.Args) == 2 {
		file, err := os.Open(os.Args[0])
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
	}
	sc := bufio.NewScanner(reader)
	var input string
	for sc.Scan() {
		input += sc.Text()
	}

	tokens := scanTokens(input)

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func scanTokens(input string) Tokens {
	tokens := make(Tokens, 0)
	splitted := strings.Split(input, " ")
	for _, token := range splitted {
		tokens = append(tokens, Token(token))
	}
	return tokens
}
