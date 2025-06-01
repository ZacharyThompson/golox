run:
	go run src/main.go src/scanner.go src/token.go
input:
	go run src/main.go src/scanner.go src/token.go ./input.lox
generate_ast:
	mkdir -p ast
	go run src/tool/generate_ast.go ./src
