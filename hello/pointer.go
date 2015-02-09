package main

import "fmt"

// This `person` struct type has `name` and `age` fields.
type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("=== introduction")
	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initializing a struct.
	// This is better since its more resilient to field order change.
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})
	fmt.Println(person{name: "38"})

	// An `&` prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// Access struct fields with a dot.
	fmt.Println("=== s")
	s := person{name: "Sean", age: 50}
	fmt.Println(s) // => {Sean 50}
	// passing pointer
	fmt.Println(s.name)    // => Sean
	fmt.Println(&s.name)   // => 0x2081c60e0
	fmt.Println((&s).name) // => Sean

	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age) // 50

	// following line cause error, trying to dereferencing `sp.age`
	// but sp.age's type is int `50` not pointer(storing address of memory)
	// fmt.Println(*sp.age)

	fmt.Println((*sp).age) // 50

	fmt.Println("=== p_s =======")
	// p_s is pointer
	p_s := &person{name: "Sean", age: 50}
	fmt.Println(p_s) //=> &{Sean 50}

	// s_p is pointer
	// when you use, dot(.) operator against struct, if struct is pointer,
	// its automatically dereferenced since,
	// (*p_s).name == p_s.name
	fmt.Println((*p_s).name) //=> Sean
	fmt.Println(p_s.name)    //=> Sean

	// p_s is pointer, if you try to pointerize further
	// it mean pointer to p_s(this also pointer) cause error
	// so this **p_s have no field of 'name' cause compile err
	// fmt.Println((&p_s).name)

	// but with `.`, cause one level dereference, following code is work
	fmt.Println((*(&p_s)).name)  // Sean
	fmt.Println((**(&p_s)).name) // Sean

	fmt.Println(&p_s) //=> 0x2081b0020
	pp_s := &p_s
	fmt.Println(*pp_s)  //=> &{Sean 50}
	fmt.Println(**pp_s) //=> {Sean 50}

	// following code cause error because of affinity
	// that pp_s.name is evaluated first.
	// fmt.Println(*pp_s.name) //=> Sean

	// we need to evaluate *pp_s first so add parenthesis to tell compilier
	// order of evaluation.
	fmt.Println((*pp_s).name) //=> Sean

	// the address of name
	fmt.Println(&(*pp_s).name) //=> Sean

	// here further take off!
	fmt.Println("=== ppp_s ")

	ppp_s := &pp_s
	fmt.Println(ppp_s)    //=> 0x2081b2028
	fmt.Println(*ppp_s)   //=> 0x2081b2020
	fmt.Println(**ppp_s)  //=> &{Sean 50}
	fmt.Println(***ppp_s) //=> {Sean 50}

	fmt.Println((**ppp_s).name)  //=> &{Sean 50}.name => Sean
	fmt.Println((***ppp_s).name) //=>  {Sean 50}.name => Sean

	fmt.Println("=== end =======")

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)
}
