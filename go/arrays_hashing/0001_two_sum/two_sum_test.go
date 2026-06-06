package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"Example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"Example 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"Example 3", []int{3, 3}, 6, []int{0, 1}},
		{"No Solution", []int{1, 2, 3}, 7, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSumHashMap(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSumHashMap() = %v, want %v", got, tt.want)
			}

			if got := TwoSumArray(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSumArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

var benchNums = []int{11, 15, 1, 8, 9, 14, 22, 33, 44, 55, 66, 77, 88, 99, 100, 101, 102, 103, 104, 105, 2, 7}
var benchTarget = 9

func BenchmarkTwoSumHashMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSumHashMap(benchNums, benchTarget)
	}
}

func BenchmarkTwoSumArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSumArray(benchNums, benchTarget)
	}
}
