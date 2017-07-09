package main

import "fmt"

func main2() {
	a := [2]int{1, 2}
	b := [2]int{1, 2}
	fmt.Println(a == b)
	p := new([10]int)
	fmt.Println(p)
	x := [10]int{}
	x[1] = 2
	fmt.Println(x)
	p[2] = 2
	fmt.Println(p)
	y := [2][3]int{
		{1, 2, 3},
		{4, 5, 6}}
	fmt.Println(y)

	z := [...][2]int{
		{1, 2},{2, 3},{4, 6}}
	fmt.Println(z)
}
func main() {
	a := [...]int{2, 8, 4, 7, 1, 9, 5, 6, 3, 0}
	fmt.Println(a)
	len := len(a)
	for i := 0; i < len; i++{
		for j := i + 1; j < len; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
				fmt.Println(a)
			}
		}
	}
	fmt.Println(a)
}
