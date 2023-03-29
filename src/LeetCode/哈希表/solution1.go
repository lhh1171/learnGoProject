package 哈希表

// 暴力解法
func twoSum(nums []int, target int) []int {
	for k1, _ := range nums {
		for k2 := k1 + 1; k2 < len(nums); k2++ {
			if target == nums[k1]+nums[k2] {
				return []int{k1, k2}
			}
		}
	}
	return []int{}
}

// 使用map方式解题，降低时间复杂度
// 利用map找到之前的index
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		//找之前的index
		if preIndex, ok := m[target-val]; ok {
			return []int{preIndex, index}
		} else {
			m[val] = index
		}
	}
	return []int{}
}
