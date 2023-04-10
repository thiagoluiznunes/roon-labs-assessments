# Reverse Polish Notation

The objective in this project is to build a complete, polished, but _small_ product: A 
[reverse polish notation calculator](https://en.wikipedia.org/wiki/Reverse_Polish_notation) in the style of a unix command line tool.

# How to run it

### Requirements ###

It is necessary to install previously Go language and Make.

* **[Golang 1.9.x](https://go.dev/doc/install)** :white_check_mark:
* **[Make](https://www.gnu.org/software/make/#download)** :white_check_mark:

### Running

- After installing the requirements run the followed command to install the go packages.
```console
$ go mod tidy
```
- You can run the rpn cli in the interactive mode by:
```console
$ go run main.go
```
- You also can run in the no-interactive mode, passing the expression:
```console
$ go run main.go [EXPRESSION]
```
- To access the help tab, use the followed command:
```console
$ go run main.go -h
```
