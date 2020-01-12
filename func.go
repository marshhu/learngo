package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval2(a, b int, op string) (int , error){
	switch op {
	case "+":
		return a + b,nil
	case "-":
		return a - b,nil
	case "*":
		return a * b,nil
	case "/":
		q,_ := div(a,b)
		return q,nil
	default:
		return 0,errors.New("unsupported operation: "+ op)
	}
}

func div(a,b int)(int,int){
	return a/b,a%b
}

//命名返回值，一旦命名
func div2(a,b int)(q int,r int){
	q = a/b
	r = a%b
	return
}

//函数作为参数
func apply(a,b int,op func(a,b int) int) int{
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d,%d)\n",opName,a,b)
	return op(a,b)
}


//可变参数列表
func sum(values ...int) int{
	s := 0
	for _,v := range values{
		s += v
	}
	return s
}

//交换两个数
func swap(a,b int)(int,int){
	return  b,a
}

func main() {
	if result,err := eval2(3,4,"x");err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}

	q,r := div(10,3)
	fmt.Println(q,r)

	fmt.Println(
	)

	apply(3,4, func(a, b int) int {
		return int(math.Pow(float64(a),float64(b)))
	})

	fmt.Println(sum(3,5,8,12))

	a,b := swap(3,4)
	fmt.Println(a,b)
}
