package main

import (
	"fmt"
	"os"

	"github.com/t9md/go-learn/stringutil"
)

func main() {
	hn, _ := os.Hostname()
	fmt.Printf("Hello, world. my hostname is %s\n", hn)
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}
