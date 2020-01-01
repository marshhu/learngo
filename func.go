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

func main() {
	if result,err := eval2(3,4,"x");err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}

	q,r := div(10,3)
	fmt.Println(q,r)

	fmt.Println(apply(3,4, func(a, b int) int {
		return int(math.Pow(float64(a),float64(b)))
	}))

	fmt.Println(sum(3,5,8,12))
}
