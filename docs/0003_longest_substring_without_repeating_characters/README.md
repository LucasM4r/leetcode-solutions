## Longest Substring Without Repeating Characters

To solve this problem efficiently, I used the **Sliding Window** technique. This approach allows us to find the result by traversing the string only once, optimizing the time complexity from $O(N^2)$ to $O(N)$.

### Algorithm Visualization
The GIF below demonstrates how the `Left` (L) pointer "jumps" to the correct position whenever a repeated character is encountered, keeping the window valid at all times:

![Sliding Window Visualization](../../assets/0003_longest_substring_without_repeating_characters/algorithm_sliding_window.gif)

### Complexity Analysis
* **Time Complexity:** $O(N)$, where $N$ is the length of the string. The algorithm processes each character exactly once using the `right` pointer.
* **Space Complexity:** $O(min(N, M))$, where $M$ is the size of the character set (e.g., 256 for extended ASCII), used by the map to store the last seen index of each character.

### Implementations

* 🐹 **[Go Solution](../../go/sliding_window/0003_longest_substring_without_repeating_characters/longest_substring_without_repeating_characters.go)**