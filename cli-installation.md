# How to install RPN calculator in Linux systems

## Unizip file
```console
$ tar -xvf rpn.tar.gz
```
**Info**: You can use -C flag to extract in a specific directory
```console
$ tar -vf rpn.tar.gz -C ~/Downloads/your_directory
```

## Move bin/rpn binary
- In this step you will need to add the RPN binary to you user local optional softwares directory.
```console
$ mv bin/rpn /usr/local/opt
```
- After that, you need to add the /usr/loca/opt path in your $PATH environment to expose the path of the new software for the whole thesystem.
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
