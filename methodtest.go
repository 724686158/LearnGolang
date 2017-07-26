package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string
}

type TZ int

func main() {
	a := A{}
	a.Print()
	fmt.Println(a)
	b := B{}
	b.Print()
	var tz TZ
	tz.Print()
}

func (a *A) Print() {
	a.Name = "AA"
	fmt.Println("A")
}


func (b B) Print() {
	fmt.Println("B")
}

func (a *TZ) Print(){
	fmt.Println("TZ")
}