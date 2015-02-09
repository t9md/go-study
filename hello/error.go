package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func process(fname string, f func(int) (int, error), numbers []int) {
	for _, i := range numbers {
		if r, e := f(i); e != nil {
			fmt.Println(fname, " failed:", e)
		} else {
			fmt.Println(fname, " worked:", r)
		}
	}
}

func main() {
	numbers := []int{7, 42}
	process("f1", f1, numbers)
	process("f2", f1, numbers)

	_, e := f2(42)
	fmt.Println("========")
	if ae, ok := e.(*argError); ok {
		fmt.Println("========")
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
