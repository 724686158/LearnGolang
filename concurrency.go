package main

import (
	"fmt"
	"time"
)

func main() {
	go Go()
	go func () {
		fmt.Println("GO GO GO!!!!")
	}()
	time.Sleep(2 * time.Second)
}

func Go()  {
	fmt.Println("GO GO GO!!!!")
}