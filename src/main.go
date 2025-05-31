package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func run(input string) {
	fmt.Println(input)
}

func runFile(filePath string) error {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	text := string(fileData)
	run(text)
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
		run(scanner.Text())
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
