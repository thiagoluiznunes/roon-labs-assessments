# Reverse Polish Notation

The objective in this project is to build a complete, polished, but _small_ product: A 
[reverse polish notation calculator](https://en.wikipedia.org/wiki/Reverse_Polish_notation) in the style of a unix command line tool.

## Requirements ###

It is necessary to install previously Go language and Make.

* **[Golang 1.9.x](https://go.dev/doc/install)** :white_check_mark:
* **[Make](https://www.gnu.org/software/make/#download)** :white_check_mark:

## Release RPN CLI
- After installing the requirements run the followed command to build a new release of the RPN cli.
```console
$ make build
```
## Move bin/rpn binary
- In this step you will need to add the RPN binary to you user local optional softwares directory.
```console
$ mv bin/rpn /usr/local/opt
```
- After that, you need to add the /usr/local/opt path in your $PATH environment to expose the path of the new software for the whole thesystem.
```console
$ export PATH=$PATH:/usr/local/opt
```
**Info**: To apply this new PATH release in every new terminal session, you will need to update it in your **.bash_profile** or if your are using zsh you will need add in **.zshrc**. Example below:
```console
$ export PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/opt/
```

## Check installation
- You can check the installation running the command in your terminal:
```console
$ rpn -h
```
