class Solution(object):
    
    def twoSumDictionary(self, nums, target):
        """
        twoSumDictionary finds the indices of two numbers using a Dictionary and add up to the target.
        Time Complexity: O(N)
        Space Complexity: O(N)

        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        numDict = {}

        for i in range(len(nums)):
            complement = target - nums[i]

            if complement in numDict:
                return [numDict[complement], i]
            numDict[nums[i]] = i
        return None
    
    def twoSumArray(self, nums, target):
        """
        twoSumArray finds the indices of two numbers using a Array Brute Force and add up to the target.
        Time Complexity: O(N^2)
        Space Complexity: O(1)

        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for i in range(len(nums)):
            for j in range(i+1, len(nums)):
                if nums[i] + nums[j] == target:
                    return [i,j]
        return None
