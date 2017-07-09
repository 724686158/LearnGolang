package main

import "fmt"

func main() {
	x, y := 1, 2
	a := [2]int{2, 4}
	b := [1]int{1}
	c := [...]int{99:19}
	d := [...]*int{&x, &y}
	var p *[100]int = &c
	fmt.Println(a, b, c)
	fmt.Println(p)
	fmt.Println(d)
}
