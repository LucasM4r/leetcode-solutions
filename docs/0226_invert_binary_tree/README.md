## Invert Binary Tree

To solve this problem efficiently, I used a **Recursive (Pre-order Traversal)** approach. This method allows us to visit each node, swap its left and right children, and then recursively apply the exact same logic to invert the remaining subtrees.

### Algorithm Visualization
The GIF below demonstrates the pre-order traversal in action. It highlights the current node being visited and animates the geometric swap of the entire left and right subtrees before descending further down the tree:

![Invert Binary Tree](../../assets/0226_invert_binary_tree/algorithm_invert_binary_tree.gif)

### Complexity Analysis
* **Time Complexity:** $O(N)$, where $N$ is the number of nodes in the tree. The algorithm visits every single node exactly once to perform the swap.
* **Space Complexity:** $O(N)$ in the worst case (a completely skewed tree) due to the maximum depth of the recursion stack. In the best/average case (a completely balanced tree), the space complexity is $O(\log N)$.

### Implementations

* 🐹 **[Go Solution](../../go/trees/0226_invert_binary_tree/invert_binary_tree.go)**