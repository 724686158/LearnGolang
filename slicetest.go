package main

import "fmt"

func main() {
	a1 := [12]int{}
	fmt.Println(a1)
	s1 := a1[:] // a[5, 6, 7, 8, 9]
	s2 := make([]int , 3, 10)
	fmt.Println(s1)
	fmt.Println(s2)

	a2 := []byte{'a', 'b', 'c', 'd', 'e' , 'f', 'g', 'h'}
	s3 := a2[3:6]
	s4 := a2[4:6]
	s5 := s3[2:5]
	s4[1] = 'x'
	fmt.Println(len(s4), cap(s4))

	s4 = append(s4, 'x', 'x', 'x', 'x', 'x', 'x', 'x')
	copy(s3, s4)
	fmt.Println(len(s4), cap(s4))
	fmt.Println(len(s3), cap(s3))
	fmt.Println(string(s3), string(s4))
	fmt.Println(string(s5))




}

func main2() {
	s6 := make([]int, 3, 6)
	fmt.Println(&s6)
	s6 = append(s6, 1, 2)
	fmt.Println(&s6, len(s6), cap(s6))
	s6 = append(s6, 1, 2, 3, 4)
	fmt.Println(&s6, len(s6), cap(s6))


}
