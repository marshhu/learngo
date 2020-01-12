package main

import "fmt"

//go语言中，数组是值类型
func printArray(arr [5]int){
	arr[0] = 100
	for i,v := range arr{
		fmt.Println(i,v)
	}
}

func main() {
	var arr1 [5]int //声明一位数组
	var grid [4][5]int //声明二位数组

	arr2 := [3]int{1,3,5} //初始化数组
	arr3 := [...]int{2,4,6,8,10} //初始化数组

	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)

	fmt.Println("printArray(arr1)")
	printArray(arr1)



}
