package main

import (
	"fmt"

	"github.com/Piyush01Bhatt/go-lsm-db/internal/langparser"
)

func main() {
	fmt.Println("Your server code goes here.")
	lexer := langparser.NewLexer("select name,roll from class where section = 'A'")
	tokens := lexer.Tokens()
	fmt.Println(tokens)
}
