/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {

	var res []int
	var m = map[int]int{}
	for i:= 0; i <len(nums); i++ {
		 
		if v, ok := m[target - nums[i]]; !ok {
			m[nums[i]] = i
		} else {
			res = append(res, v, i)
			break
		}
	}

	return res
}

// @lc code=end

