package tools

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func Length(s string) int {
	return len(s)
}

func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, r := range s1 {
		m1[r]++
	}
	for _, r := range s2 {
		m2[r]++
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
