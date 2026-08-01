package findthesafestpathinagrid

import "container/list"

func maximumSafenessFactor(grid [][]int) int {
	type Point struct {
		r, c int
	}
	n := len(grid)

	distance := make([][]int, n)
	for i := 0; i < n; i++ {
		distance[i] = make([]int, n)
		for j := 0; j < n; j++ {
			distance[i][j] = -1
		}
	}

	queue := list.New()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				distance[i][j] = 0
				queue.PushBack(Point{i, j})
			}
		}
	}

	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)

		p := elem.Value.(Point)
		for _, d := range directions {
			newR, newC := p.r+d.r, p.c+d.c
			if newR >= 0 && newR < n && newC >= 0 && newC < n && distance[newR][newC] == -1 {
				distance[newR][newC] = distance[p.r][p.c] + 1
				queue.PushBack(Point{newR, newC})
			}
		}
	}

	canReach := func(safeness int) bool {
		if distance[0][0] < safeness || distance[n-1][n-1] < safeness {
			return false
		}

		visited := make([][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([]bool, n)
		}

		q := list.New()
		q.PushBack(Point{0, 0})
		visited[0][0] = true

		for q.Len() > 0 {
			elem := q.Front()
			q.Remove(elem)
			p := elem.Value.(Point)

			if p.r == n-1 && p.c == n-1 {
				return true
			}

			for _, d := range directions {
				newR, newC := p.r+d.r, p.c+d.c
				if newR >= 0 && newR < n && newC >= 0 && newC < n && !visited[newR][newC] && distance[newR][newC] >= safeness {
					visited[newR][newC] = true
					q.PushBack(Point{r: newR, c: newC})
				}
			}
		}
		return false
	}

	left := 0
	right := min(distance[0][0], distance[n-1][n-1])
	ans := 0
	for left < right {
		mid := (left + right + 1) / 2
		if canReach(mid) {
			left = mid
			ans = mid
		} else {
			right = mid - 1
		}
	}

	return ans
}
