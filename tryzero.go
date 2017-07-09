package main

import (
	"fmt"
	"math"
)
var (
	a int
	b string
	c bool
	d float64
	e byte
	arr []int
	com complex128
	ptr uintptr
)

type (
	文本 string
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(arr)
	fmt.Println(com)
	fmt.Println(ptr)
	fmt.Println(math.MaxFloat64)
	fmt.Println(b)
	a1, _, c3, d4 := 1, 2, 3, 4
	fmt.Println(c3)
	fmt.Println(a1)

	fmt.Println(d4)

	var aa float64 = 100.02
	num := int64(aa)
	fmt.Println(num)

}