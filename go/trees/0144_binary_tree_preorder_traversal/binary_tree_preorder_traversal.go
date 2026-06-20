package binarytreepreordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	var traversal func(node *TreeNode)

	if root == nil {
		return []int{}
	}

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		result = append(result, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)

	return result
}
