build: main.go config.go
	go build -o bin/term-games

run: main.go config.go
	go run .
