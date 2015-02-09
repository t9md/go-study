package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

// with pointer dereference
func (h *Human) Hi() {
	fmt.Printf("Hi, I am %s, %d years old. you can call me on %s\n", (*h).name, (*h).age, (*h).phone)
}
func (h *Human) SetName(name string) {
	(*h).name = name
}

func (h Human) UselessSetName(name string) {
	h.name = name
	fmt.Printf("inner function, name is %s. but useless\n", h.name)
}

// without pointer dereference
func (h *Human) SetAge(age int) {
	h.age = age
}
func (h *Human) GoodBye() {
	fmt.Printf("GoodBye, I am %s, %d years old. you can call me on %s\n", h.name, h.age, h.phone)
}

func main() {
	taku := Human{name: "Taku", age: 25, phone: "222-222-XXX"}
	taku.Hi()
	taku.GoodBye()
	fmt.Println("-- useless change-------")
	fmt.Println(taku.name)
	taku.UselessSetName("Yoko")
	fmt.Println(taku.name)
	fmt.Println("--change-------")
	taku.SetName("Taro")
	taku.SetAge(10)
	taku.Hi()
	taku.GoodBye()

	fmt.Println("-- with assigning pointer to struct -------")
	yoko := &Human{name: "Yoko", age: 18, phone: "1234-222-XXX"}
	yoko.Hi()
	yoko.GoodBye()
	fmt.Println("-- useless change-------")
	fmt.Println(yoko.name)
	yoko.UselessSetName("Yoko2")
	fmt.Println(yoko.name)
	fmt.Println("--change-------")
	yoko.SetName("Taro")
	yoko.SetAge(10)
	yoko.Hi()
	yoko.GoodBye()
}

/* output
Hi, I am Mike you can call me on 222-222-XXX
GoodBye, I am Mike you can call me on 222-222-XXX
--change-------
Hi, I am Taro you can call me on 222-222-XXX
GoodBye, I am Taro you can call me on 222-222-XXX
*/
