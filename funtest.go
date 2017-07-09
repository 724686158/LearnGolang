package main

import (
	"fmt"
	"math"
)

func main() {
	c, d := Swap(1, 2)
	fmt.Println(c, d)
	f := func (numbers ...int) (max int)  {

		max = math.MinInt64
		for _, k := range numbers{
			if k > max{
				max = k
			}
		}
		return max
	}
	fmt.Println(Max(1, 23, 3, 4, 5 ,6 , 99, -1, -3), f(-2, -3, -6, -11))
}

func Swap(a int64, b int64) (c int64, d int64) {
	c, d = b, a
	return c, d
}

func Max(numbers ...int) (max int)  {

	max = math.MinInt64
	for _, k := range numbers{
		if k > max{
			max = k
		}
	}
	return max
}