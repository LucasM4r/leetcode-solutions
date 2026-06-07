package addtwonumbers

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

func TestAddTwoNumbers(t *testing.T) {
	// Usando Table-Driven Tests
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		// Casos de teste clássicos do LeetCode
		{"Exemplo 1 Clássico", []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{"Zeroes", []int{0}, []int{0}, []int{0}},
		{"Tamanhos Diferentes e Carry no final", []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := sliceToList(tt.l1)
			list2 := sliceToList(tt.l2)

			gotList := AddTwoNumbers(list1, list2)

			got := listToSlice(gotList)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	l1 := sliceToList([]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9})
	l2 := sliceToList([]int{9, 9, 9, 9})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		AddTwoNumbers(l1, l2)
	}
}
