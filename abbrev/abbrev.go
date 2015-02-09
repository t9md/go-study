package abbrev

// New retrurn abbrev map form abbrev to non-abbrev strings.
// ex) New(["abc", "def"])
//   => map[def:def abc:abc a:abc ab:abc d:def de:def]
func New(words []string) map[string]string {
	s := ""
	m := make(map[string]int)
	table := make(map[string]string)
	str := []rune{}

	for _, word := range words {
		str = []rune(word)
		for i := 0; i < len(str); i++ {
			s = string(str[0:i])
			m[s] += 1
			table[s] = word
		}
	}
	for k, v := range m {
		if v > 1 {
			delete(table, k)
		}
	}
	for _, word := range words {
		table[word] = word
	}
	return table
}
