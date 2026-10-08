/*
- Constaints:
- 1 <= s.length, t.length <= 5 * 10^4
- s and t consist of lowercase English letters.
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	const LOWERCASE_EN_LETTERS = 26
	const A_ASCII = 'a'
	var freq [26]int

	for i := 0; i < len(s); i++ {
		freq[s[i]-A_ASCII]++
		freq[t[i]-A_ASCII]--
	}

	for _, count := range freq {
		if count != 0 {
			return false
		}
	}
	return true
}
