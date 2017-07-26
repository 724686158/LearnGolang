package main

import "fmt"

type USB interface {
	Name() string
	Connect()
}
type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() (string) {
	return pc.name
	
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func Disconnecter(usb USB) {
	fmt.Println("Disconnected.", usb.Name())
	pc, ok := usb.(PhoneConnecter);
	if ok {
		fmt.Println("Dis", pc.name )

	}
}

func main() {
	var a USB
	a = PhoneConnecter{"SONY"}
	a.Connect()
	Disconnecter(a)
}