package main

import "fmt"

func printSlice(s []int){
	fmt.Printf("len=%d,cap=%d,%v\n",len(s),cap(s),s)
}

func main() {
	var s []int
	for i:=0;i<100;i++{
		printSlice(s)
		s = append(s,2*i+1)
    }

	s1 := []int{2,4,6,8}
	printSlice(s1)

	s2 := make([]int,16)
	s3 := make([]int,10,32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2,s1)
	printSlice(s2)

}
