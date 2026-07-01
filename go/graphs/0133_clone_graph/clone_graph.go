package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)
	var DepthFirstSearch func(n *Node) *Node

	DepthFirstSearch = func(n *Node) *Node {
		if clone, exists := visited[n]; exists {
			return clone
		}

		clone := &Node{Val: n.Val}

		visited[n] = clone

		for _, neighbor := range n.Neighbors {
			clone.Neighbors = append(clone.Neighbors, DepthFirstSearch(neighbor))
		}

		return clone
	}

	return DepthFirstSearch(node)
}
