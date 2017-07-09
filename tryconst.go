package main

import "fmt"

const a int = 1
const b = 'A' + 2

const (
	c = a
	d = a + 2
)
/*
6 : 0110
11: 1011
------------------
&   0010 2
|   1111 15
^   1101 13
&^  0100 4
*/

const (
	MONDAY = 'A'
	TUESDAY = iota
	WEDNESDAY
	THURDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

const (
	BYTE float64 = 1 << (iota * 10)
	KILO_BYTE
	MEGA_BYTE
	GIGA_BYTE
	TERA_BYTE
	PETA_BYTE
	EXA_BYTE
	ZETTA_BYTE
	YOTTA_BYTE
)
func main() {
	fmt.Println(BYTE)
	fmt.Println(KILO_BYTE)
	fmt.Println(MEGA_BYTE)
	fmt.Println(GIGA_BYTE)
	fmt.Println(TERA_BYTE)
	fmt.Println(PETA_BYTE)
	fmt.Println(EXA_BYTE)
	fmt.Println(ZETTA_BYTE)
	fmt.Println(YOTTA_BYTE)
	/*

	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)
	a := 1
	if(a != 0 && 1 / a > 1){
		fmt.Println("OK")
	}
	fmt.Println(0.7 / 0.2)
	 */

}
