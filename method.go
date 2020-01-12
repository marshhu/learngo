package main

import "fmt"

type Employee struct{
	name string
	salary int
	currency string
	age int
}
func (e Employee) displaySalary(){
	fmt.Printf("Salary of %s is %s %d\n",e.name,e.currency,e.salary)
}

/*
使用值接收器的方法
 */
func (e Employee) changeName(newName string){
	e.name = newName
}
/*
使用指针接收器的方法
 */
func (e *Employee) changeAge(newAge int){
	e.age = newAge
}

func main() {
	e1 := Employee{
		name:"lily",
		salary:5000,
		currency:"$",
	}
	e1.displaySalary()

	e2 := Employee{
		name:"Mark andrew",
		age: 50,
	}

	fmt.Printf("Employee name before change: %s\n",e2.name)
	e2.changeName("Michael Andrew")
	fmt.Printf("Employee name after change: %s\n",e2.name)

	fmt.Printf("Employee age before change: %d\n",e2.age)
	e2.changeAge(51)
	fmt.Printf("Employee age after change: %d\n",e2.age)

}
