/**
 * Constraints:
 * 0 <= nums.length <= 10^5
 * -10^9 <= nums[i] <= 10^9
 */
func hasDuplicate(nums []int) bool {
	numsLength := len(nums)
	if numsLength == 0 {
		return false
	}
	// Create a hashmap for seen entries
	seen := make(map[int]bool, numsLength)

	for _, num := range nums {
		if ok, _ := seen[num]; ok {
			return true
		} else {
			seen[num] = true
		}
	}
	return false
}
