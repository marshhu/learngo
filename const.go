package main

import "fmt"

func main() {
	//iota只能在const中使用，初始为0，每次出现自动递增1
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b, kb, mb, gb, tb, pb)
}
