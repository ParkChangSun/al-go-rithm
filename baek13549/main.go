package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, k int

var dist []int

func main() {
	std := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer std.Flush()

	fmt.Fscanln(std, &n, &k)

	dist = make([]int, 100001)
	for i := range dist {
		dist[i] = -1
	}

	solve()

	fmt.Fprintln(std, dist[k])
}

func solve() {
	dist[n] = 0
	deque := []int{}
	deque = append(deque, n)

	for len(deque) > 0 {
		cur := deque[0]
		deque = deque[1:]

		if cur == k {
			return
		}

		if l := cur * 2; l <= 100000 && dist[l] == -1 {
			dist[l] = dist[cur]
			deque = append(deque, 0)
			copy(deque[1:], deque)
			deque[0] = l
		}

		if l := cur - 1; l >= 0 && dist[l] == -1 {
			dist[l] = dist[cur] + 1
			deque = append(deque, l)
		}

		if l := cur + 1; l <= 100000 && dist[l] == -1 {
			dist[l] = dist[cur] + 1
			deque = append(deque, l)
		}
	}
}

// 01 bfs는 반드시 한번씩만 방문한다 o(v+e)
// 그래서 한번씩만 방문하는 논리는 사용하려면 방문의 우선순위를 잘 따져봐야 한다
// 만약 +1을 우선시한다면
// 4-5-6 = 2

// 만약 -1을 우선시한다면
// 4-3-6 = 1

// 두배가 될 수 있기 때문에 -를 먼저 처리해야 한다.

// 현재 노드에서 두배로 접근할 수 있는 것들은 다 했으니까 빼고

// 이 노드에서 +로 접근할 수 있는 것들은 -를 통해 두배로 접근할 수 있다.
