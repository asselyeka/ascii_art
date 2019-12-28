package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	arg := os.Args
	arguments := []string(arg[1:])
	str := ""
	lenArg := 0
	for range arguments {
		lenArg++
	}
	for i, input := range arguments {
		if i == lenArg-1 {
			str += input
			break
		}
		str += input + " "
	}
	strAsByte := []byte(str)

	fileName := "standard.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer file.Close()

	rawBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var rawRune []rune
	for _, elem := range rawBytes {
		rawRune = append(rawRune, rune(elem))
	}

	lines := strings.Split(string(rawRune), "\n")

	for h := 0; h < 9; h++ {
		for _, l := range strAsByte {
			for i, line := range lines {
				if i == (int(l)-32)*9+h {
					fmt.Print(line)
				}
			}
		}
		fmt.Println()
	}
}
