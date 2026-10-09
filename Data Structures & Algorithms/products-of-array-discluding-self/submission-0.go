/**
 * Constraints
 * 2 <= nums.length <= 100,000
 * -30 <= nums[i] <= 30
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
 */
func productExceptSelf(nums []int) []int {
	numsLength := len(nums)
	// Solve using prefix and suffix
	var prefix = make([]int, numsLength)
	var suffix = make([]int, numsLength)
	var result = make([]int, numsLength)

	prefix[0] = 1
	for i := 1; i < numsLength; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	suffix[numsLength-1] = 1
	for j := numsLength - 2; j >= 0; j-- {
		suffix[j] = suffix[j+1] * nums[j+1]
	}

	for index, _ := range nums {
		result[index] = prefix[index] * suffix[index]
	}

	return result
}
