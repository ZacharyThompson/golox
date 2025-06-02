package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func handleErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func defineType(w *bufio.Writer, baseName string, structName string, fieldList string) error {
	_, err := w.WriteString("type " + structName + " struct {\n")
	handleErr(err)

	// Fields
	fields := strings.Split(fieldList, ",")
	for i := range fields {
		fields[i] = strings.TrimSpace(fields[i])
	}
	for _, field := range fields {
		_, err = w.WriteString("\t" + field + "\n")
		handleErr(err)
	}
	_, err = w.WriteString("}\n\n")
	handleErr(err)

	//Constructor
	_, err = w.WriteString("func New" + structName + "(" + fieldList + ") " + structName + " {\n")
	handleErr(err)
	_, err = w.WriteString("\t return " + structName + "{")
	handleErr(err)
	// Store parameters in fields
	for i, field := range fields {
		suffix := ", "
		if i == len(fields)-1 {
			suffix = ""
		}
		_, err = w.WriteString(strings.Split(field, " ")[0] + suffix)
		handleErr(err)
	}
	_, err = w.WriteString("}\n")
	handleErr(err)
	_, err = w.WriteString("}\n\n")
	handleErr(err)
	return nil
}

func defineAst(outputDir string, baseName string, types []string) error {
	path := outputDir + "/" + strings.ToLower(baseName) + ".go"
	f, err := os.Create(path)
	defer f.Close()
	handleErr(err)
	w := bufio.NewWriter(f)
	defer w.Flush()
	_, err = w.WriteString("package main\n\n")
	handleErr(err)
	_, err = w.WriteString("type " + baseName + " interface{}\n\n")
	handleErr(err)

	for _, t := range types {
		structName := strings.Split(t, ":")[0]
		structName = strings.TrimSpace(structName)
		fields := strings.Split(t, ":")[1]
		fields = strings.TrimSpace(fields)
		defineType(w, baseName, structName, fields)
	}

	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: generate_ast <output directory>")
		os.Exit(64)
	}
	outputDir := os.Args[1]
	fmt.Println(outputDir)
	err := defineAst(outputDir, "Expr", []string{
		"Binary   : left Expr, operator Token, right Expr",
		"Grouping : expression Expr",
		"Literal  : value any ",
		"Unary    : operator Token, right Expr",
	})
	handleErr(err)
}
