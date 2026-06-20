package binarytreepostordertraversal

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
		current := queue[0]
		queue = queue[1:]

		if nums[i] != NULL {
			current.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, current.Left)
		}
		i++

		if i < len(nums) && nums[i] != NULL {
			current.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, current.Right)
		}
		i++
	}
	return root
}

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, NULL, 2, 3},
			want: []int{3, 2, 1},
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4, 5, NULL, 8, NULL, NULL, 6, 7, 9},
			want: []int{4, 6, 7, 5, 2, 9, 8, 3, 1},
		},
		{
			name: "Empty Tree",
			nums: []int{},
			want: []int{},
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
			got := postorderTraversal(root)

			if !reflect.DeepEqual(got, tt.want) && !(len(got) == 0 && len(tt.want) == 0) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPostorderTraversal(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, NULL, 8, NULL, NULL, 6, 7, 9}
	root := buildTree(nums)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postorderTraversal(root)
	}
}
