/*
*
  - Constraints
  - 1 <= nums.length <= 10^4.
  - -1000 <= nums[i] <= 1000
  - 1 <= k <= number of distinct elements in nums.
*/
func topKFrequent(nums []int, k int) []int {
	nums_length := len(nums)

	if nums_length == k {
		// According to constraints k is number of distict elements in nums
		// so distinct elements and length of array can't have duplicates
		return nums
	}

	const INDEX_OFFSET int = 1000

	// Get the frequencies for all the numbers
	var freqArray [2001]uint16
	for _, num := range nums {
		freqArray[num+INDEX_OFFSET]++
	}

	// Make the freq as index and numbers as values in a bucket sort fashion
	bucketMap := make([][]int16, nums_length+1)
	for n, freq := range freqArray {
		bucketMap[freq] = append(bucketMap[freq], int16(n-INDEX_OFFSET))
	}

	// Find out the result
	var result = make([]int, 0, k)

	for i := nums_length; i >= 0 && k > 0; i-- {
		// If the frequency exists
		for _, num := range bucketMap[uint16(i)] {
			result = append(result, int(num))
			k--

			if k == 0 {
				return result
			}
		}

	}
	return result
}