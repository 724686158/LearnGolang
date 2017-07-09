package main

import "fmt"

type human struct {
	Sex int
	Age int
}

type teacher struct {
	human
	Name string
	Age int
}

type student struct {
	human
	Name string
	Age int
}



func main() {
	a := teacher{human: human{Sex: 0, Age: 99}, Name: "joe", Age:19}
	b := student{human: human{Sex: 1, Age: 99}, Name: "joe", Age:19}
	fmt.Println(a, b)
	a.Age = 100
	b.human.Age = 100
	fmt.Println(a, b)
}