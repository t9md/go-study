package main

import (
	"fmt"
	"github.com/kr/pretty"
)

type Human struct {
	name string
}

type Singer struct {
	Human
	song string
}

func (s *Singer) Sing() {
	fmt.Printf("la la la...")
}

type Painter struct {
	Human
	picture string
}

func (p *Painter) Paint() {
	fmt.Printf("peta peta peta...")
}

type Singable interface {
	Sing()
}

type Paintable interface {
	Paint()
}

func main() {
	var bowie Singable = &Singer{Human: Human{name: "bowie"}, song: "starman"}
	var picasso Paintable = &Painter{Human: Human{name: "picasso"}, picture: "guernica"}
	if p, ok := bowie.(Singable); ok {
		pretty.Println(p.(*Singer).name)
	}
	if p, ok := picasso.(Paintable); ok {
		pretty.Println(p.(*Painter).picture)
	}
}
