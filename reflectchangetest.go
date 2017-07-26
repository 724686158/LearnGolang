package main

import (
	"reflect"
	"fmt"
)

type User struct {
	Id int
	Name string
	Age int
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet(){
		fmt.Println("xxx")
		return
	} else{
		v = v.Elem()
	}

	if f := v.FieldByName("Name"); f.Kind() == reflect.String{
		f.SetString("BYEBYE")
	}

}

func main() {
	u := User{1, "Ok", 12}
	Set(&u)
	fmt.Println(u)
}

func main2() {
	x := 123
	fmt.Println(x)
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)
}
