/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	stdReader := bufio.NewReader(os.Stdin)
	var sendData string
	var err error
	for {
		fmt.Printf("%s > ", strings.TrimSpace(sendData))
		sendData, err = stdReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin")
			panic(err)
		}
	}
}
