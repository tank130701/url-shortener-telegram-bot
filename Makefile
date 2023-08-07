.SILENT:

build:
	go build -o ./.bin cmd/main.go
run:
	go run cmd/main.go