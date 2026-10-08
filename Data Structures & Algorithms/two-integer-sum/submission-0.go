/*
* 	Constraints:
  - 2 <= nums.length <= 1000
    -10,000,000 <= nums[i] <= 10,000,000
    -10,000,000 <= target <= 10,000,000
  - Only one valid answer exists.
*/
func twoSum(nums []int, target int) []int {

	numsLength := len(nums)
	numMap := make(map[int]int, numsLength)

	for i := 0; i < numsLength; i++ {
		diff := target - nums[i]
		// Check if other part is already in the map
		if j, ok := numMap[diff]; ok {
			return []int{j, i}
		}
		numMap[nums[i]] = i
	}

	return nil
}
