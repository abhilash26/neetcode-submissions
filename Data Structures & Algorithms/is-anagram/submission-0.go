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
	freqMap := make(map[byte]int, LOWERCASE_EN_LETTERS)

	for index, _ := range s {
		freqMap[s[index]]++
		freqMap[t[index]]--
	}

	for _, count := range freqMap {
		if count != 0 {
			return false
		}
	}
	return true
}