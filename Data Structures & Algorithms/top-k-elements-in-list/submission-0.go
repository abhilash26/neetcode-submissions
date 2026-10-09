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

	// Get the frequencies for all the numbers
	freqMap := make(map[int16]uint16, nums_length)
	for _, num := range nums {
		freqMap[int16(num)]++
	}

	// Make the freq as index and numbers as values in a bucket sort fashion
	bucketMap := make(map[uint16][]int16, nums_length)
	for n, freq := range freqMap {
		bucketMap[freq] = append(bucketMap[freq], n)
	}

	// Find out the result
	var result = make([]int, 0, k)

	for i := nums_length; i >= 0 && k > 0; i-- {
		if _, ok := bucketMap[uint16(i)]; !ok {
			continue
		}

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