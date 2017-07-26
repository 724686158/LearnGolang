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

func (u User)WHOAREYOU(name string) {
	fmt.Println("Hello", name, ", My name is", u.Name)

	
}
func main() {
	// 常规方法调用
	u := User{1, "mengzicheng", 12}
	u.WHOAREYOU("laowang")

	v := reflect.ValueOf(u)
	mv := v.MethodByName("WHOAREYOU")
	args := []reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)
}
