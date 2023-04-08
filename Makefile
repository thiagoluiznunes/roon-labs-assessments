# 'setup' requirements
build:
	go build -o bin/rpn

package: build
	tar -cvzf dist/rpn.tar.gz bin/rpn installation.md
