package clonegraph

import (
	"testing"
)

func buildGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	nodes := make(map[int]*Node)

	for i := range adjList {
		val := i + 1
		nodes[val] = &Node{Val: val}
	}

	for i, neighbors := range adjList {
		val := i + 1
		for _, neighborVal := range neighbors {
			nodes[val].Neighbors = append(nodes[val].Neighbors, nodes[neighborVal])
		}
	}

	return nodes[1]
}

func areGraphsEqualButDistinct(t *testing.T, original, clone *Node) bool {
	if original == nil && clone == nil {
		return true
	}
	if original == nil || clone == nil {
		return false
	}

	visitedOrig := make(map[*Node]bool)
	visitedClone := make(map[*Node]bool)

	var check func(n1, n2 *Node) bool
	check = func(n1, n2 *Node) bool {
		if n1.Val != n2.Val {
			return false
		}
		if len(n1.Neighbors) != len(n2.Neighbors) {
			return false
		}
		if n1 == n2 {
			t.Errorf("Deep copy failed: Node with value %d shares the same memory as the original", n1.Val)
			return false
		}

		visitedOrig[n1] = true
		visitedClone[n2] = true

		for i := range n1.Neighbors {
			neigh1 := n1.Neighbors[i]
			neigh2 := n2.Neighbors[i]

			if visitedOrig[neigh1] {
				if !visitedClone[neigh2] {
					return false
				}
				continue
			}

			if !check(neigh1, neigh2) {
				return false
			}
		}
		return true
	}

	return check(original, clone)
}

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name    string
		adjList [][]int
	}{
		{
			name:    "Complex graph with cycles (Example 1)",
			adjList: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			name:    "Single node without neighbors (Example 2)",
			adjList: [][]int{{}},
		},
		{
			name:    "Empty graph (Example 3)",
			adjList: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalGraph := buildGraph(tt.adjList)
			clonedGraph := cloneGraph(originalGraph)

			if !areGraphsEqualButDistinct(t, originalGraph, clonedGraph) {
				t.Errorf("Cloned graph does not match the expected for test: %s", tt.name)
			}
		})
	}
}

func BenchmarkCloneGraph(b *testing.B) {
	adjList := make([][]int, 100)
	for i := 0; i < 100; i++ {
		prev := i
		next := i + 2

		if i == 0 {
			prev = 100
		}
		if i == 99 {
			next = 1
		}
		adjList[i] = []int{prev, next}
	}

	largeGraph := buildGraph(adjList)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = cloneGraph(largeGraph)
	}
}
