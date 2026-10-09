/**
 * Constraints
 * 2 <= nums.length <= 100,000
 * -30 <= nums[i] <= 30
 * The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
 */
func productExceptSelf(nums []int) []int {
	numsLength := len(nums)
	// Solve using prefix and suffix
	var result = make([]int, numsLength)

	prefix := 1
	for i, num := range nums {
		result[i] = prefix
		prefix *= num
	}

	suffix := 1
	for i := numsLength - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}