package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jonathanlloyd/crafting_interpreters/golox/internal/lexer"
)

func run(source string) {
	fmt.Printf("%s\n", source)
}

func runFile(filename string) {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	run(string(source))
}

func runPrompt() {
	fmt.Println("=================")
	fmt.Println("golox interpreter")
	fmt.Println("=================")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(">>> ")
	for scanner.Scan() {
		source := scanner.Text()
		_ = source
		fmt.Println(lexer.Token{Type: lexer.IF, Line: 3, Literal: "if"})
		fmt.Printf(">>> ")
	}
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: golox [script]")
		os.Exit(1)
	}
	if len(os.Args) == 2 {
		filename := os.Args[1]
		runFile(filename)
	} else {
		runPrompt()
	}
}
