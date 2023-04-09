# 'setup' requirements
build:
	go build -o bin/rpn

package: build
	cp bin/rpn dist
	cp installation.md dist
	tar -cvzf rpn.tar.gz dist
