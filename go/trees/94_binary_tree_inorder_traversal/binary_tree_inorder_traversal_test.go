package binarytreeinordertraversal

import (
	"math"
	"reflect"
	"testing"
)

const NULL = math.MinInt32

func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 || nums[0] == NULL {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(nums) {
		curr := queue[0]
		queue = queue[1:]

		if i < len(nums) && nums[i] != NULL {
			curr.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, curr.Left)
		}
		i++

		if i < len(nums) && nums[i] != NULL {
			curr.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, curr.Right)
		}
		i++
	}
	return root
}

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, NULL, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4, 5, NULL, 8, NULL, NULL, 6, 7, 9},
			want: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
		{
			name: "Empty Tree",
			nums: []int{},
			want: nil,
		},
		{
			name: "Single Node",
			nums: []int{1},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.nums)
			got := inorderTraversal(root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkInorderTraversal(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, NULL, 8, NULL, NULL, 6, 7, 9}
	root := buildTree(nums)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		inorderTraversal(root)
	}
}
