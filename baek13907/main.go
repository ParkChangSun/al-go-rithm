package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type Route struct {
	current int
	dist    int
	toll    int
}
type PriorityQueue []*Route

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Route))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

type Edge struct {
	dest  int
	value int
}

var edge = map[int][]Edge{}

var n, m, k, s, d int
var memo [][]int
var taxes []int

func main() {
	std := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer std.Flush()

	fmt.Fscanln(std, &n, &m, &k)
	for i := 1; i <= n; i++ {
		edge[n] = make([]Edge, 0)
	}
	fmt.Fscanln(std, &s, &d)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscanln(std, &a, &b, &c)
		edge[a] = append(edge[a], Edge{b, c})
		edge[b] = append(edge[b], Edge{a, c})
	}
	taxes = make([]int, k+1)
	for i := 1; i <= k; i++ {
		fmt.Fscanln(std, &taxes[i])
	}

	memo = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			memo[i][j] = math.MaxInt32
		}
	}

	dijkstra()

	arr := memo[d]
	for _, tax := range taxes {
		min := math.MaxInt32
		cut := 0
		for i := range arr {
			arr[i] += tax * i
			if arr[i] < min {
				min = arr[i]
				cut = i
			}
		}
		fmt.Fprintln(std, min)
		arr = arr[:cut+1]
	}
}

func dijkstra() {
	memo[s][0] = 0
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Route{s, 0, 0})

	for pq.Len() != 0 {
		r := heap.Pop(pq).(*Route)

		if r.toll > n || r.dist > memo[r.current][r.toll] {
			continue
		}

	OPTIMALFOUND:
		for _, e := range edge[r.current] {
			min := e.value + memo[r.current][r.toll]

			for i := 1; i < r.toll; i++ {
				if memo[e.dest][i] <= min {
					continue OPTIMALFOUND
				}
			}

			if min < memo[e.dest][r.toll+1] {
				memo[e.dest][r.toll+1] = min
				heap.Push(pq, &Route{e.dest, min, r.toll + 1})
			}
		}
	}
}

// 지금 가려고 하는 루트, 도시의
// 이 도시로 가는 길 중 톨이 지금 루트보다 작고 거리고 작다?
// 톨은 작지만 거리는 길어야 해
// 이 루트가 톨도 많고 거리도 더 길다?
// 이 뜻은 중간도시를 굳이 돌아간다는거지
// 중간도시를 굳이 돌아갈 필요는 없다
//

// 도로는 반드시 999여야 한다 memo[n][n] 인데 나는 처음에 memo[n][m]으로 함
