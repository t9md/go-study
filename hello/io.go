package main

import (
	"fmt"
	"os"

	"github.com/kr/pretty"
	"syscall"
)

var pp = pretty.Println

func main() {
	file, err := os.Open("/etc/hosts")
	if err != nil {
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	// pp(stat.Name())
	// pp(stat.Mode())
	// pp(stat.ModTime())
	// pp(stat.IsDir())
	_sys := stat.Sys()
	// pp(_sys)
	if v, ok := _sys.(*syscall.Stat_t); ok {
		pp(v)
		pp("---------")
		pp(v.Mode)
		pp(v.Blocks)
	}

	// pp(stat)
	bs := make([]byte, stat.Size())

	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	os.Exit(0)
	fmt.Println(str)
}
