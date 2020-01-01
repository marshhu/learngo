package main

import (
	"fmt"
	"strconv"
)

func convertToBin(n int) string{
	if n == 0{
		return "0"
	}
	result := ""
	for ;n > 0;n/=2{
		lst := n%2
		result = strconv.Itoa(lst)+result
	}
	return  result
}
func main() {
	fmt.Println(convertToBin(5),convertToBin(85),convertToBin(5685),convertToBin(0))
}
