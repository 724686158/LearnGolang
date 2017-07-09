package main

import "fmt"

type person struct {
	Name string
	Age int
}

func main() {
	a := &struct {
		Name string
		Age int
	}{
		Name:"joe",
		Age:19,
	}
	fmt.Println(a)

	b := person{
		Name:"wang",
		Age:17,
	}
	c := person{
		Name:"wang",
		Age:17,
	}
	fmt.Println(b == c)

	fmt.Println(b)
	BeOlder(&b)
	fmt.Println(b)

	fmt.Println(b == c)
}

func BeOlder(per *person)  {
	per.Age++
}