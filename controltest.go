package main

import "fmt"

func main1() {
	a := 11
	if a := 1; a > 10{
		fmt.Println(a)
	}else {
		fmt.Println(a)
	}
	println(a)
}
func main() {
	a :=0
	for {
		a++
		if a > 3{
			break
		}
		fmt.Println(a)
	}
	fmt.Println("over")
}