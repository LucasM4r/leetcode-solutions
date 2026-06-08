package symmetrictree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetric checks whether a binary tree is a mirror of itself (symmetric around its center).
// Time Complexity: O(N), where N is the number of nodes in the tree. Each node is visited once in the worst case.
// Space Complexity: O(N) in the worst case (skewed tree) due to the recursion stack, or O(log N) in the best case (balanced tree).
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return compare(root.Left, root.Right)

}

func compare(node1 *TreeNode, node2 *TreeNode) bool {

	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val && compare(node1.Left, node2.Right) && compare(node1.Right, node2.Left)
}
