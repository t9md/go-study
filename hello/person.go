package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) Name() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	fmt.Println((&Person{"Taro", "Yamada"}).Name())
	// fmt.Println(person.Name())
}
