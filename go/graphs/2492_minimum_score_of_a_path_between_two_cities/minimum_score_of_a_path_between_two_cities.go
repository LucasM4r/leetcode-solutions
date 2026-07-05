package minimumscoreofapathbetweentwocities

func minScore(n int, roads [][]int) int {
	type edge struct {
		to       int
		distance int
	}

	graph := make([][]edge, n+1)

	for _, road := range roads {
		u, v, distance := road[0], road[1], road[2]

		graph[u] = append(graph[u], edge{to: v, distance: distance})
		graph[v] = append(graph[v], edge{to: u, distance: distance})
	}

	visited := make([]bool, n+1)
	minDistance := 10001

	var depthFirstSearch func(node int)
	depthFirstSearch = func(node int) {
		visited[node] = true

		for _, neighbor := range graph[node] {
			if neighbor.distance < minDistance {
				minDistance = neighbor.distance
			}

			if !visited[neighbor.to] {
				depthFirstSearch(neighbor.to)
			}
		}
	}

	depthFirstSearch(1)
	return minDistance
}
