package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

type Manager struct {
	User
	title string
}

func (u User) Hello() {
	fmt.Println("Hello world", u.Name)
}

func Info(o interface{})  {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields:吧")

	for i := 0; i < t.NumField(); i++{
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v", f.Name, f.Type, val)
	}
}

func main() {
	u := User{1, "ok", 12}
	Info(u)
	fmt.Println()
	m := Manager{User:User{1, "ok", 12}, title:"123"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0))
}