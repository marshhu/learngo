package main

import "fmt"

func main() {
	m := map[string]string{
		"name":"ccmouse",
		"course":"golang",
		"site":"imooc",
		"quality":"notbad",
	}
	fmt.Println(m)

	m2 := make(map[string]int)
	fmt.Println(m2)

	var m3 map[string]int
	fmt.Println(m3)

	fmt.Println("Traversing map")
	for k,v := range m{
		fmt.Println(k,v)
	}

	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)

	if causeName,ok := m["cause"];ok{
		fmt.Println(causeName)
	}else{
		fmt.Println("key dose not exist")
	}

	fmt.Println("Deleting values")
	name,ok := m["name"]
	fmt.Println(name,ok)

	delete(m,"name")
	name,ok = m["name"]
	fmt.Println(name,ok)
}
