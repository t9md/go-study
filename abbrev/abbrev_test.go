package abbrev

import "testing"

type expect struct {
	input  string
	output string
}

func TestNew(t *testing.T) {
	var abbrevTest = []expect{
		{"a", "ab"},
		{"ab", "ab"},
		{"c", "cd"},
		{"cd", "cd"},
	}

	var table = make(map[string]string)
	table = New([]string{"ab", "cd"})
	for _, tt := range abbrevTest {
		if got := table[tt.input]; got != tt.output {
			t.Errorf("Expected: %q, got: %q", tt.output, tt.input)
		}
	}
}

func TestNewComplex(t *testing.T) {
	var abbrevTest = []expect{
		{"Hy", "Hydrogen"},
		{"He", "Helium"},
		{"アイラン", "アイランド"},
		{"アイメ", "アイメイク"},
		{"ユニコー", "ユニコード"},
		{"ユニコ", "ユニコ"},
		{"Lithi", "Lithium"},
		{"Be", "Beryllium"},
		{"Bo", "Boron"},
		{"Carbon", "Carbon"},
		{"Nitrogen", "Nitrogen"},
		{"O", "Oxygen"},
		{"Fl", "Fluorine"},
		// ambiguous input
		{"B", ""},
		{"H", ""},
	}

	table := New([]string{
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
	})
	for _, tt := range abbrevTest {
		if got := table[tt.input]; got != tt.output {
			t.Errorf("Expected: %q for input %q, got: %q", tt.output, tt.input, got)
		}
	}
}
