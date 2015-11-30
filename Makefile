all: build

build: export GOOS=linux
build:
	go build -o importer main.go database.go