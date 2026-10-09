/**
* Constraints
- 1 <= strs.length <= 10000.
- 0 <= strs[i].length <= 100
- strs[i] is made up of lowercase English letters.
*/

func groupAnagrams(strs []string) [][]string {
	const ALPHABET_SIZE = 26
	const ASCII_A = 'a'

	strs_length := len(strs)

	if strs_length == 0 {
		return [][]string{}
	}

	// Make the frequencies as key and the strings as values in a map
	freqMap := make(map[[ALPHABET_SIZE]int][]string, strs_length)

	// Set up frequency map for each string
	for _, str := range strs {
		var freq [ALPHABET_SIZE]int
		for _, ch := range str {
			freq[ch-ASCII_A]++
		}
		freqMap[freq] = append(freqMap[freq], str)
	}

	// Create a result slice to hold the grouped anagrams
	var result [][]string
	for _, group := range freqMap {
		result = append(result, group)
	}
	// Return result as order does not matter
	return result
}