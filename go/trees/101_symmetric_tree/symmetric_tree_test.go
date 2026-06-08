package symmetrictree

import (
	"math"
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

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "Empty tree",
			input:    []int{},
			expected: true,
		},
		{
			name:     "Single root node",
			input:    []int{1},
			expected: true,
		},
		{
			name:     "Perfectly symmetric tree",
			input:    []int{1, 2, 2, 3, 4, 4, 3},
			expected: true,
		},
		{
			name:     "Asymmetric tree (different structure)",
			input:    []int{1, 2, 2, NULL, 3, NULL, 3},
			expected: false,
		},
		{
			name:     "Asymmetric tree (different values)",
			input:    []int{1, 2, 2, 3, 4, 4, 5},
			expected: false,
		},
		{
			name:     "Large symmetric tree",
			input:    []int{1, 2, 2, 3, 4, 4, 3, 5, 6, 7, 8, 8, 7, 6, 5},
			expected: true,
		},
		{
			name:     "One side null and the other is not",
			input:    []int{1, 2, NULL},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)

			result := isSymmetric(root)
			if result != tt.expected {
				t.Errorf("isSymmetric() = %v, expected %v for input %v", result, tt.expected, tt.input)
			}
		})
	}
}

func BenchmarkIsSymmetric(b *testing.B) {
	input := []int{1, 2, 2, 3, 4, 4, 3, 5, 6, 7, 8, 8, 7, 6, 5}
	root := buildTree(input)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		isSymmetric(root)
	}
}
