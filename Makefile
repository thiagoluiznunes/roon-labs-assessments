# 'setup' requirements
build:
	go build -o bin/rpn

package: build
	cp -R bin rever-polish-notation
	cp installation.md rever-polish-notation
	tar -cvzf rpn.tar.gz rever-polish-notation
