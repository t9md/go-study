package main

import (
	"fmt"
	"strings"
)

// Indent is used to keep track of depth of recursive call
type Indent struct {
	base  string
	depth int
}

func (i *Indent) String() string {
	return strings.Repeat(i.base, i.depth)
}

func fact(n int, i *Indent) int {
	fmt.Printf("%s%d: enter\n", i, n)

	if n == 0 {
		fmt.Printf("%s%d: leave\n", i, n)
		return 1
	}

	i.depth += 1
	answer := fact(n-1, i)
	i.depth -= 1
	fmt.Printf("%s%d: leave\n", i, n)
	return n * answer
}

// By power of defer, easily visualizable for depth of recursive call.
// Without depending on Indent custom type.
// This idea is borrowed from effective-go
func trace(n, depth int) (int, int) {
	fmt.Printf("%s%d: enter\n", strings.Repeat("  ", depth), n)
	return n, depth
}
func un(n, depth int) {
	fmt.Printf("%s%d: leave\n", strings.Repeat("  ", depth), n)
}

func factWithDefer(n, depth int) int {
	defer un(trace(n, depth))
	if n == 0 {
		return 1
	}
	return n * factWithDefer(n-1, depth+1)
}

func main() {
	fmt.Println(fact(7, &Indent{base: "  ", depth: 0}))
	fmt.Println("----------")
	fmt.Println(factWithDefer(7, 0))
}

/*
=>
     7: called
       6: called
         5: called
           4: called
             3: called
               2: called
                 1: called
                   0: called
                   0: returned
                 1: returned
               2: returned
             3: returned
           4: returned
         5: returned
       6: returned
5040
*/
