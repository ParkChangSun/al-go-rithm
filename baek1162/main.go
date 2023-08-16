package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
)

var n, m, k int

type E struct {
	dest int
	dist int
}

func main() {
	std := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer std.Flush()

	fmt.Fscanln(std, &n, &m, &k)
	edge := make(map[int][]E)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscanln(std, &a, &b, &c)
		if edge[a] == nil {
			edge[a] = make([]E, 0)
		}
		edge[a] = append(edge[a], E{b, c})
		if edge[b] == nil {
			edge[b] = make([]E, 0)
		}
		edge[b] = append(edge[b], E{a, c})
	}

	fmt.Println(dikjstra(edge))
}

type Item struct {
	value     int
	priority  int
	skippable int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func dikjstra(edge map[int][]E) int {
	memo := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			memo[i][j] = math.MaxInt64
		}
	}
	memo[1][k] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{value: 1, priority: 0, skippable: k})

	for len(pq) != 0 {
		t := heap.Pop(&pq)
		i := t.(*Item)

		if i.priority > memo[i.value][i.skippable] {
			continue
		}

		for _, v := range edge[i.value] {
			if i.skippable > 0 {
				if smin := memo[i.value][i.skippable]; smin < memo[v.dest][i.skippable-1] {
					memo[v.dest][i.skippable-1] = smin
					heap.Push(&pq, &Item{v.dest, smin, i.skippable - 1})
				}
			}

			if nmin := memo[i.value][i.skippable] + v.dist; nmin < memo[v.dest][i.skippable] {
				memo[v.dest][i.skippable] = nmin
				heap.Push(&pq, &Item{v.dest, nmin, i.skippable})
			}
		}
	}
	return min(memo[n])
}

func min(a []int) int {
	sort.Ints(a)
	return a[0]
}
