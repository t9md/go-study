package main

import "fmt"

// This `person` struct type has `name` and `age` fields.
type person struct {
	name string
	age  int
}

func main() {

	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})
	fmt.Println(person{name: "38"})

	// An `&` prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// Access struct fields with a dot.
	fmt.Println("=== s =======")
	// s is not pointer
	s := person{name: "Sean", age: 50}
	fmt.Println(s) // => {Sean 50}

	// proper way
	fmt.Println((&s).name) // => Sean

	// if 's' is not pointer, compler automatically
	// treat as s is pointer so either way is legal
	fmt.Println(s.name) // => Sean

	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age) // 50

	fmt.Println("=== p_s =======")
	// p_s is pointer
	p_s := &person{name: "Sean", age: 50}
	fmt.Println(p_s) //=> &{Sean 50}

	// s_p is pointer
	fmt.Println(p_s.name) //=> Sean

	// p_s is pointer, if you try to pointerize further
	// it mean pionter to pointer to p_s cause error
	// so this **p_s have no field of 'name' cause compile err
	// fmt.Println((&p_s).name)

	fmt.Println(&p_s) //=> 0x2081b0020
	pp_s := &p_s
	fmt.Println(*pp_s) //=> &{Sean 50}

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
