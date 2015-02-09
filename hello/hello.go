package main

import (
	"fmt"
	"os"

	"github.com/t9md/go-learn/stringutil"
)

func main() {
	var yes stringutil.Hogehoge = 77
	hn, _ := os.Hostname()
	pf := fmt.Printf
	p := fmt.Println
	p(yes)
	p(stringutil.CallPrivateType())
	pf("Hello, world. my hostname is %s\n", hn)
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}
