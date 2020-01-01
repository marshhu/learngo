package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	if contents, err := ioutil.ReadFile("abc.txt"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", contents)
	}
}
