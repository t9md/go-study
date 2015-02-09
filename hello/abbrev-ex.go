package main

import (
	"fmt"
	"os"

	"github.com/t9md/go-learn/abbrev"
)

func main() {
	words := []string{
		"Hydrogen",
		"Helium",
		"アイランド",
		"アイメイク",
		"ユニコード",
		"ユニコ",
		"Lithium",
		"Beryllium",
		"Boron",
		"Carbon",
		"Nitrogen",
		"Oxygen",
		"Fluorine",
		"Neon",
	}
	table := abbrev.New(words)
	fmt.Println(abbrev.New([]string{"abc", "def"}))
	os.Exit(0)
	// fmt.Println(abbrev_table)
	var input string
	for {
		fmt.Printf("input? ")
		fmt.Scanf("%s", &input)

		switch input {
		case "quit":
			os.Exit(0)
		default:
			if val, ok := table[input]; ok {
				fmt.Printf("Found!: %s => %s\n", input, val)
			} else {
				fmt.Printf("NotFound: %s\n", input)
			}
		}
	}
}
