package main

import "fmt"

type MYINT int

func (a *MYINT) Increase() {
	for i := 0; i < 100; i++{
		*a++
	}

}

func main() {
	var a MYINT
	a = 0
	a.Increase()
	fmt.Println(a)
}
