package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jonathanlloyd/crafting_interpreters/golox/internal/lexer"
)

func run(source string) error {
	tokens, err := lexer.Scan(source)
	if err != nil {
		return err
	}

	for _, token := range tokens {
		fmt.Println(token)
	}
	return nil
}

func runFile(filename string) {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = run(string(source))
	if err != nil {
		panic(err)
	}
}

func runPrompt() {
	fmt.Println("=================")
	fmt.Println("golox interpreter")
	fmt.Println("=================")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(">>> ")
	for scanner.Scan() {
		source := scanner.Text()
		err := run(source)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		}
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
