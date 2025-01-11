package main

import (
	"fmt"

	"github.com/Piyush01Bhatt/go-lsm-db/internal/langparser"
)

func main() {
	fmt.Println("Your server code goes here.")
	lexer := langparser.NewLexer("select name,")
	var tokens []langparser.Token
	for {
		token := lexer.NextToken()
		if token.Type == langparser.TokenEOF {
			break
		}
		tokens = append(tokens, token)
	}
	fmt.Println(tokens)
}
