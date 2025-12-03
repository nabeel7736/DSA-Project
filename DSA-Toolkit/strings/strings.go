package strings

func ReverseString(s string) string {
	r := []rune(s)
	l, rIndex := 0, len(r)-1
	for l < rIndex {
		r[l], r[rIndex] = r[rIndex], r[l]
		l++
		rIndex--
	}
	return string(r)
}

func IsPalindrome(s string) bool {
	r := []rune(s)
	l, rIndex := 0, len(r)-1
	for l < rIndex {
		if r[l] != r[rIndex] {
			return false
		}
		l++
		rIndex--
	}
	return true
}

func CountVowelsConsonants(s string) (int, int) {
	vowels := "aeiouAEIOU"
	v, c := 0, 0
	for _, ch := range s {
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
			if containsRune(vowels, ch) {
				v++
			} else {
				c++
			}
		}
	}
	return v, c
}

func containsRune(s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}
