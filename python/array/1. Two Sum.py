class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        
        numbers_map = dict()
        
        for i, n in enumerate(nums):
            diff = target - n
            
            if diff in numbers_map:
                return numbers_map[diff], i
            
            numbers_map[n] = i