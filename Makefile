#

# Run all tests in the project
eval:
	go test -v ./tests/...

# Compiles project to ./target/main
build:
	go build -o target/main src/main.go

# Removes builds artifacts
clean:
	go mod tidy
	rm -rf target

run:
	go run src/main.go