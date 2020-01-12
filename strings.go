package main

import "fmt"

func printBytes(s string){
	for i:= 0;i < len(s);i++{
		fmt.Printf("%x ",s[i])
	}
}

func printChars(s string){
	for i:= 0;i < len(s);i++{
		fmt.Printf("%c ",s[i])
	}
}
func printCharsRune(s string){
	runes :=[]rune(s)
	for i:= 0;i < len(runes);i++{
		fmt.Printf("%c ",runes[i])
	}
}

//for range循环遍历字符串，循环返回的是当前rune的字节位置
func printCharsAndBytes(s string){
	for i,v := range s{
		fmt.Printf("%c starts at byte %d\n",v,i)
	}
}


func main() {
	str :="Hello World"
	printBytes(str)

	fmt.Printf("\n")
	printChars(str)

	str ="Señor"
	fmt.Printf("\n")
	printBytes(str)
	//ñ的Unicode编码占用了两个字节 c3 b1
	fmt.Printf("\n")
	printChars(str)
	fmt.Printf("\n")
	printCharsRune(str)//使用rune
	fmt.Printf("\n")
	printCharsAndBytes(str)
}
