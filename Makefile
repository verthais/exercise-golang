#

eval:
	go test ./test

build:
	go build -o target/main src/main.go

clean:
	rm -rf target

run:
	go run src/main.go