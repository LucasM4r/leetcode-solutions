package findasafewalkthroughagrid

import (
	"container/heap"
	"math"
)

type Point struct {
	r, c, damage int
}

type PriorityQueue []*Point

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].damage < pq[j].damage
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Point))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func findSafeWalk(grid [][]int, health int) bool {
	m, n := len(grid), len(grid[0])
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	minDamage := make([][]int, m)

	for i := range minDamage {
		minDamage[i] = make([]int, n)

		for j := range minDamage[i] {
			minDamage[i][j] = math.MaxInt32
		}
	}

	pq := make(PriorityQueue, 0, m*n)
	heap.Init(&pq)

	minDamage[0][0] = grid[0][0]
	heap.Push(&pq, &Point{0, 0, grid[0][0]})

	for pq.Len() > 0 {
		point := heap.Pop(&pq).(*Point)

		if point.r == m-1 && point.c == n-1 {
			return health-point.damage > 0
		}

		if point.damage > minDamage[point.r][point.c] {
			continue
		}

		for _, dir := range directions {
			newR, newC := point.r+dir[0], point.c+dir[1]
			if newR >= 0 && newR < m && newC >= 0 && newC < n {
				newDamage := point.damage + grid[newR][newC]

				if newDamage < minDamage[newR][newC] {
					minDamage[newR][newC] = newDamage
					heap.Push(&pq, &Point{newR, newC, newDamage})
				}
			}
		}
	}
	return false
}
