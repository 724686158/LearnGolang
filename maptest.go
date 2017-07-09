package main

import "fmt"

func main2() {
	m1 := map[int]string{}
	m1[1] = "OK"
	delete(m1, 1)
	fmt.Println(m1)

	mm1 := map[int]map[int]string{}
	mm1[1] = make(map[int]string)
	mm1[1][1] = "OK"
	fmt.Println(mm1)
	a, ok := mm1[2][1]

	if ok {
		fmt.Println(a)
	} else {
		fmt.Println("empty")
	}
}

func main4() {
	arr1 := [...]int{2, 3, 4, 5, 6}
	s1 := arr1[2:]
	for i, v := range s1{
		fmt.Println(i, v)
	}

	sm := make([]map[int]string, 5)
	for i, v := range sm{
		fmt.Println(v)
		sm[i] = make(map[int]string)
		sm[i][1] = "HELLO"
	}
	fmt.Println(sm)
}

func main() {
	m1 := map[int]string{1:"a", 2:"b", 3:"c"}
	m2 := make(map[string]int, len(m1))
	for i, v := range m1{
		m2[v] = i
	}
	fmt.Println(m2)
}