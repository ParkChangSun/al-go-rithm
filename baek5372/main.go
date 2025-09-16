package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

var std = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

var n, k, m int
var min int

var graph []uint64

func main() {
	defer std.Flush()

	var cases int
	fmt.Fscan(std, &cases)
	for i := 0; i < cases; i++ {
		fmt.Fscan(std, &n, &k, &m)

		min = 51
		graph = make([]uint64, n+1)

		for j := 0; j < m; j++ {
			a, b := 0, 0
			fmt.Fscan(std, &a, &b)
			graph[a] = graph[a] | (1 << b)
			graph[b] = graph[b] | (1 << a)
		}

		backtrack(uint64(1<<(n+1)-2), k)

		if min == 51 {
			fmt.Fprintln(std, "IMPOSSIBLE")
		} else {
			fmt.Fprintln(std, min)
		}
	}
}

func backtrack(alive uint64, kleft int) {
	for change := true; change; {
		change = false

		for i := 1; i <= n; i++ {
			if alive&(1<<i) == 0 {
				continue
			}

			edges := bits.OnesCount64(alive & graph[i])

			if edges == 0 {
				alive -= 1 << i
			} else if edges == 1 {
				alive -= 1 << i
				alive -= 1 << bits.TrailingZeros64(alive&graph[i])
				kleft--
				change = true
			} else if edges > kleft {
				alive -= 1 << i
				kleft--
				change = true
			}
		}
	}

	if alive == 0 {
		if kleft >= 0 && k-kleft < min {
			min = k - kleft
		}
		return
	}
	if kleft <= 0 {
		return
	}

	p := bits.TrailingZeros64(alive)
	t := bits.TrailingZeros64(alive & graph[p])

	backtrack(alive-(1<<p), kleft-1)
	backtrack(alive-(1<<t), kleft-1)
}

// 3
// 5
// 5
// 2
// 1 3
// 2 3
// 3
// 1
// 3
// 1 2
// 2 3
// 1 3
// 6
// 3
// 6
// 1 3
// 2 3
// 2 6
// 3 5
// 3 6
// 4 5
