package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Shravan-1908/odyssey/lexer"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// lexer := lexer.NewLexer("5+55")
	// lexer.Tokenize()

	for {
		fmt.Print("odyssey> ")

		text, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToLower(text)) == "exit" {
			break
		}

		lexer := lexer.NewLexer(text)
		tokens, _ := lexer.Tokenize()

		for _, token := range tokens {
			fmt.Printf("%v: %v\n", token.Type, token.Value)
		}
	}
}