package main

func twoSum(nums []int, target int) []int {
	ans := make([]int, 2)
	for i, _ := range nums {
		for k, _ := range nums {
			if i < k {
				if target == nums[i]+nums[k] {
					ans[0] = i
					ans[1] = k
					return ans
				}
			}
		}
	}
	return ans
}
