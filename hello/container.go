package main

import (
	"container/list"
	"fmt"
	"github.com/kr/pretty"
)

var pp = pretty.Println

func main() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	pp("")
	for e := x.Front(); e != nil; e = e.Next() {
		// fmt.Println(e.Value.(int))
		// fmt.Println(e.Value.(int))
		fmt.Println(e.Value.(int))
	}
}
