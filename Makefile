# 'setup' requirements
install:
	go mod tidy

run:
	go run main.go

build: install
	go build -o bin/rpn

release: build
	cp -R bin reverse-polish-notation
	cp cli-installation.md reverse-polish-notation
	tar -cvzf rpn.tar.gz reverse-polish-notation
