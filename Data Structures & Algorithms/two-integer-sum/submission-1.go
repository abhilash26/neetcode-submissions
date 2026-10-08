/*
* 	Constraints:
  - 2 <= nums.length <= 1000
    -10,000,000 <= nums[i] <= 10,000,000
    -10,000,000 <= target <= 10,000,000
  - Only one valid answer exists.
*/
func twoSum(nums []int, target int) []int {

	numMap := make(map[int]int, len(nums))

	for i, num := range nums {
		// Check if other part is already in the map
		if j, ok := numMap[target-num]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}

	return nil
}
