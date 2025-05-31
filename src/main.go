package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func run(input string) []error {
	scanner := NewScanner(input)
	tokens := scanner.scanTokens()
	for i, t := range tokens {
		fmt.Println(i, t)
	}
	return nil
}

func runFile(filePath string) error {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	text := string(fileData)
	errs := run(text)
	if errs != nil {
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, e)
		}
		os.Exit(1)
	}
	return nil
}

func runPrompt() error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		moreTokens := scanner.Scan()
		if !moreTokens {
			break
		}
		errs := run(scanner.Text())
		// if errs != nil {
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, e)
		}
		// os.Exit(1)
		// }
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		err := runFile(os.Args[0])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		runPrompt()
	}
}
