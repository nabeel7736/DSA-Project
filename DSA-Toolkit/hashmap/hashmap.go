package hashmap

func WordFrequency(text string) map[string]int {
	freq := make(map[string]int)
	word := ""
	for _, ch := range text {
		if ch == ' ' {
			freq[word]++
			word = ""
		} else {
			word += string(ch)
		}
	}
	if word != "" {
		freq[word]++
	}
	return freq
}

type HashSet struct {
	set map[int]bool
}

func NewHashSet() *HashSet {
	return &HashSet{set: make(map[int]bool)}
}
func (h *HashSet) Add(v int)           { h.set[v] = true }
func (h *HashSet) Contains(v int) bool { return h.set[v] }
func (h *HashSet) Remove(v int)        { delete(h.set, v) }
