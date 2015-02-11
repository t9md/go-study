package main

import (
	"fmt"
	"github.com/kr/pretty"
	"time"
)

var pp = pretty.Println
var p = fmt.Println

func main() {
	// pp(time.Second)
	withTimeout(time.Second*10, func() {
		helloN(1000)
	})
	withTimeout(time.Second*10, hello)
}

func hello() {
	n := 1
	for {
		fmt.Println("hello", n)
		n++

		time.Sleep(time.Second * 1)
	}
}

func helloN(num int) {
	n := num
	for {
		fmt.Println("helloN", n)
		n++

		time.Sleep(time.Second * 1)
	}
}

func withTimeout(t time.Duration, f func()) {
	fmt.Printf("Timeout after %s\n", t)
	go f()
	select {
	case <-time.After(t):
		fmt.Println("Timeout!!")
		return
	}
}
