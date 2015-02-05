package main

import (
	"fmt"
	"os"
)

func main() {
	hn, _ := os.Hostname()
	fmt.Printf("Hello, world. my hostname is %s\n", hn)
}
