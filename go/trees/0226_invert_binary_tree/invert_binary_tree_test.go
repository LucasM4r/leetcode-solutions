package invertbinarytree

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

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1 - Full Tree",
			input:    []int{4, 2, 7, 1, 3, 6, 9},
			expected: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name:     "Example 2 - Small Tree",
			input:    []int{2, 1, 3},
			expected: []int{2, 3, 1},
		},
		{
			name:     "Example 3 - Empty Tree",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			expectedTree := buildTree(tt.expected)

			result := InvertTree(root)

			if !reflect.DeepEqual(result, expectedTree) {
				t.Errorf("got %v, want %v", result, expectedTree)
			}
		})
	}
}

func BenchmarkInvertTreeRecursive(b *testing.B) {
	input := []int{4, 2, 7, 1, 3, 6, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	root := buildTree(input)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		InvertTree(root)
	}
}
