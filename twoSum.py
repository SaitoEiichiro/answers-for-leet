from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for index1, i in enumerate(nums):
            for index2, j in enumerate(nums):
                if index1 == index2:
                    continue
                if i + j == target:
                    return [index1, index2]
        return []


sol = Solution()
print(sol.twoSum([3, 3], 6))
