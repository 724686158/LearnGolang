package main

import "fmt"

func main() {
	f := closure(10)
	fmt.Println(f(2))
	fmt.Println(f(6))
}

func closure(x int) func(int) int {
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}