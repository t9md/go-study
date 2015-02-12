/*
Here I demonstrate several aspect of Go's language characteristics, which I got
from effective-go and other documents.

This is excutable code with descriptive function name.
These functions give me good refresher to recap language specifications.
*/
package main

import (
	"fmt"
	"os"

	"github.com/kr/pretty"
	"github.com/t9md/go-learn/abbrev"
)

var pp = pretty.Println
var p = fmt.Println

func opt(options ...int) []int {
	return options
}

// NoExpectation interface is the interface allways suffice for any type T.
type NoExpectation interface{}

func AnyTypeOfArgmeentIsOK(v ...NoExpectation) {
	pp(v)
}

func HowDeferWorks() {
	p("-normal")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	p("\n-reverse by power of `defer`")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func _return(_type string) interface{} {
	return map[string]interface{}{
		"bool":   true,
		"string": "hello",
		"int":    1,
	}[_type]
}

func TypeSwitchIdiom(_type string) {
	var t interface{}
	t = _return(_type)
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T", t) // %T prints whatever type t has
	case string:
		fmt.Printf("string %s\n", t) // t has type bool
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
	pp(t)
}

func InterestingOnlyKey() {
	a := [...]int{1, 2, 3}
	for i := range a {
		println(i)
	}
	m := map[string]string{"sex": "male", "id": "t9md"}
	for i := range m {
		println(i)
	}
}

func HowAbbrevWorks() {
	v := abbrev.New([]string{"ab", "cd"})
	pp(v)
}

// Idiom: Type State have string representation of its value.
// String() is automatically invoked withing fmt.Println().
type State int

const (
	Running State = iota
	Stopped
)

func (s State) String() string {
	switch s {
	case Running:
		return "Running"
	case Stopped:
		return "Stopped"
	default:
		return "Unknown"
	}
}

func TypeWithStringRepresentation() {
	state := Running
	pp(state)
	fmt.Println(state)
}

// [effective-go]
// In these examples, the initializations work regardless of the values of
// Enone, Eio, and Einval, as long as they are distinct.
func CompositeLiteral() {
	const (
		Eio = iota
		Enone
		Einval
	)
	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	pp(a)
	pp(s)
	pp(m)
}

// Array and Slice are complete defferent type
func ArrayAndSlice() {
	// array have fixed size in bracket[]
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// for slice, not size is specified in bracket[]
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pp(a)
	pp(s)
}
func ArrayIsCopiedOnAssignment() {
	// array have fixed size in bracket[]
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a
	a[0] = 99
	pp(a)
	pp(b) // here b is intact by change of a
}

func TypeCastToInt(v interface{}) {
	// invalid operation: v + 1 (mismatched types interface {} and int)
	// fmt.Println(v + 1)

	fmt.Println(v.(int) + 1)
}

type CustomInt int

func main() {
	// HowAbbrevWorks()
	// HowDeferWorks()
	// TypeSwitchIdiom("string")
	// TypeSwitchIdiom("bool")
	// InterestingOnlyKey()
	// TypeWithStringRepresentation()
	// CompositeLiteral()
	// ArrayAndSlice()
	// ArrayIsCopiedOnAssignment()
	// AnyTypeOfArgmeentIsOK(1, 1.5, "string", true, func() {}, []byte("byte"))
	// TypeCastToInt(1)

	os.Exit(0)
}
