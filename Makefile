build:
	echo "building go application"
	go build -o ./bin/repl main.go

run:
	./bin/repl

start: build run