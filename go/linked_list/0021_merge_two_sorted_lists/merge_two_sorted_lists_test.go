package mergetwosortedlists

import (
	"reflect"
	"testing"
)

func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := &ListNode{}
	current := dummy
	for _, val := range nums {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return dummy.Next
}

func listToSlice(node *ListNode) []int {
	var res []int
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	return res
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{
			name:     "Both lists empty",
			list1:    nil,
			list2:    nil,
			expected: nil,
		},
		{
			name:     "First list empty",
			list1:    nil,
			list2:    []int{0},
			expected: []int{0},
		},
		{
			name:     "Second list empty",
			list1:    []int{1, 2, 3},
			list2:    nil,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Lists of same size with overlapping values",
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "Lists of different sizes (L1 larger)",
			list1:    []int{1, 5, 9, 10},
			list2:    []int{2, 3},
			expected: []int{1, 2, 3, 5, 9, 10},
		},
		{
			name:     "Lists of different sizes (L2 larger)",
			list1:    []int{5},
			list2:    []int{1, 2, 4, 6},
			expected: []int{1, 2, 4, 5, 6},
		},
		{
			name:     "Lists with negative numbers",
			list1:    []int{-9, -3, 0},
			list2:    []int{-5, 2, 8},
			expected: []int{-9, -5, -3, 0, 2, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := sliceToList(tt.list1)
			list2 := sliceToList(tt.list2)

			gotList := mergeTwoLists(list1, list2)
			got := listToSlice(gotList)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("mergeTwoLists(%v, %v) = %v, want %v", tt.list1, tt.list2, got, tt.expected)
			}
		})
	}
}

func generateSortedSlice(size, start, step int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = start + i*step
	}
	return slice
}

func BenchmarkMergeTwoLists(b *testing.B) {
	s1 := generateSortedSlice(10000, 0, 2)
	s2 := generateSortedSlice(10000, 1, 2)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		l1 := sliceToList(s1)
		l2 := sliceToList(s2)

		b.StartTimer()
		mergeTwoLists(l1, l2)
	}
}
