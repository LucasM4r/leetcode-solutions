import unittest
import timeit

from two_sum import Solution

class TestTwoSum(unittest.TestCase):
    def setUp(self):
        self.sol = Solution()

    def test_two_sum(self):
        tests = [
            {"name": "Example 1", "nums": [2, 7, 11, 15], "target": 9, "want": [0, 1]},
            {"name": "Example 2", "nums": [3, 2, 4], "target": 6, "want": [1, 2]},
            {"name": "Example 3", "nums": [3, 3], "target": 6, "want": [0, 1]},
            {"name": "No Solution", "nums": [1, 2, 3], "target": 7, "want": None},
        ]

        for tt in tests:
            with self.subTest(name=tt["name"]):
                
                got_dict = self.sol.twoSumDictionary(tt["nums"], tt["target"])
                self.assertEqual(got_dict, tt["want"], 
                                 f"twoSumDictionary() = {got_dict}, want {tt['want']}")

                got_array = self.sol.twoSumArray(tt["nums"], tt["target"])
                self.assertEqual(got_array, tt["want"], 
                                 f"twoSumArray() = {got_array}, want {tt['want']}")

def run_benchmarks():
    print("\n--- Running Benchmarks ---")
    
    setup_code = """
from __main__ import Solution
sol = Solution()
benchNums = [11, 15, 1, 8, 9, 14, 22, 33, 44, 55, 66, 77, 88, 99, 100, 101, 102, 103, 104, 105, 2, 7]
benchTarget = 9
    """
    
    iterations = 100000

    time_dict = timeit.timeit(
        "sol.twoSumDictionary(benchNums, benchTarget)", 
        setup=setup_code, 
        number=iterations
    )
    
    time_array = timeit.timeit(
        "sol.twoSumArray(benchNums, benchTarget)", 
        setup=setup_code, 
        number=iterations
    )

    print(f"Benchmark TwoSumDictionary: {time_dict:.4f} seconds ({iterations} loops)")
    print(f"Benchmark TwoSumArray:      {time_array:.4f} seconds ({iterations} loops)")


if __name__ == '__main__':
    unittest.main(exit=False, verbosity=2)
    
    run_benchmarks()