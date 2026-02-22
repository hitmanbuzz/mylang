package main

import (
	"fmt"
	"mylang/lexer"
	"mylang/utils"
	"os"
)

func main() {
	lexer := lexer.NewLexer()
	fname := "test.lox"

	if !utils.IsFileExist(fname) {
		fmt.Printf("File `%s` doesn't exist\n", fname)
		return
	}

	data, err := os.ReadFile(fname)
	if err != nil {
		fmt.Println("EOF  null")
		return
	}

	source := string(data)
	lexer.Source = source
	idx := 0

	for {
		if idx+1 <= len(source)-1 {
			idx = lexer.Tokenize()
		} else {
			break
		}

		lexer.CurrIdx = idx
	}

	lexer.Display()
	os.Exit(lexer.ExitCode)
}
