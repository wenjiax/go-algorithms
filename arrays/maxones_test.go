package arrays_test

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	// 给定一个二进制数组， 计算其中最大连续1的个数。

	// 示例 1:
	// 输入: [1,1,0,1,1,1]
	// 输出: 3
	// 解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
	// 注意：
	// 输入的数组只包含 0 和1。
	// 输入数组的长度是正整数，且不超过 10,000。

	t.Log(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}

// 计算最大连续 1 的个数
func findMaxConsecutiveOnes(nums []int) int {
	var max, count int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count = 0
		}

		if count > max {
			max = count
		}
	}

	return count
}
