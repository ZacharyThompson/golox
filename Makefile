run: generate_ast
	go run .
input: generate_ast
	go run . ./input.lox
generate_ast:
	mkdir -p ./generated
	go run ./tool/generate_ast.go ./generated
clean:
	rm -rf ./generated
