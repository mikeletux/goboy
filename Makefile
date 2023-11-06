build:
	mkdir -p bin/
	go build -o bin/goboy cmd/main/main.go

install:
	mkdir -p bin/
	go build -o bin/goboy cmd/main/main.go
	sudo cp bin/goboy /usr/bin/goboy

all:
	build
