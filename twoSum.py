from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in nums:
            for j in nums:
                if i + j == target:
                    return [nums.index(i), nums.index(j)]
        return []


sol = Solution()
print(sol.twoSum([2, 7, 11, 15], 9))  # [0, 1]