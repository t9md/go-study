package abbrev

// New generate abbrev map from `words` and return map from abbrev to word string.
// ex) abbrev.New([]string{"ab", "cd"})
//   => map[string]string{"ab":"ab", "a":"ab", "c":"cd", "cd":"cd"}
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
	// Even if word is part of other longer word, shorter one should always prioritized.
	for _, word := range words {
		table[word] = word
	}
	return table
}
