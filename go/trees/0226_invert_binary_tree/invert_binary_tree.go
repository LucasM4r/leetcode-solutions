package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InvertTree inverts a binary tree by swapping the left and right children of all nodes.
// Time Complexity: O(N), where N is the number of nodes in the tree. Each node is visited once.
// Space Complexity: O(N) in the worst case (skewed tree) due to the recursion stack, or O(log N) in the best case (balanced tree).
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	InvertTree(root.Left)
	InvertTree(root.Right)
	return root
}
