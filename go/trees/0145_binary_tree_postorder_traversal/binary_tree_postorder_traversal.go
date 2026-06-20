package binarytreepostordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	var traversal func(node *TreeNode)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)
		traversal(node.Right)
		result = append(result, node.Val)
	}
	traversal(root)

	return result
}
