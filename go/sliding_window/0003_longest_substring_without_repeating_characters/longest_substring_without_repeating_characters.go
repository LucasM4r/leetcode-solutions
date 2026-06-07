package longestsubstringwithoutrepeatingcharacters

// LengthOfLongestSubstringHashMap finds the length of the longest substring without repeating characters using Brute Force.
// Time Complexity: O(N^2) - We check every possible substring using a nested loop.
// Space Complexity: O(min(N, M)) - Where M is the size of the character set (e.g., 256 for ASCII).
func LengthOfLongestSubstringHashMap(s string) int {
	maxLen := 0

	for i := 0; i < len(s); i++ {
		seen := make(map[byte]bool)

		for j := i; j < len(s); j++ {
			if seen[s[j]] {
				break
			}

			seen[s[j]] = true

			length := j - i + 1
			if length > maxLen {
				maxLen = length
			}
		}
	}

	return maxLen
}

// lengthOfLongestSubstringSlidingWindow finds the length of the longest substring without repeating characters using Sliding Window.
// Time Complexity: O(N) - We iterate through the string exactly once.
// Space Complexity: O(min(N, M)) - Where M is the size of the character set (e.g., 256 for ASCII).
func lengthOfLongestSubstringSlidingWindow(s string) int {
	lastSeen := make(map[byte]int)

	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		if pos, ok := lastSeen[s[right]]; ok && pos >= left {
			left = pos + 1
		}
		lastSeen[s[right]] = right

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
