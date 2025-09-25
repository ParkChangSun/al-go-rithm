package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

var std = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

type Ticket struct {
	u, v, c, d int
}

type State struct {
	n, m, d int
}

func main() {
	defer std.Flush()

	var t int
	fmt.Fscan(std, &t)

	for i := 0; i < t; i++ {
		var n, m, k int
		fmt.Fscan(std, &n, &m, &k)
		tickets := map[int][]Ticket{}
		for i := 0; i < k; i++ {
			temp := Ticket{}
			fmt.Fscan(std, &temp.u, &temp.v, &temp.c, &temp.d)
			tickets[temp.u] = append(tickets[temp.u], temp)
		}
		for _, s := range tickets {
			slices.SortFunc(s, func(a, b Ticket) int { return a.d - b.d })
		}

		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, m+1)
			for j := range dp[i] {
				dp[i][j] = math.MaxInt
			}
		}
		dp[1][0] = 0
		queue := []State{{1, 0, 0}}

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			if cur.n == n || cur.d > dp[n][m] {
				continue
			}

			for _, ticket := range tickets[cur.n] {
				newTime := dp[cur.n][cur.m] + ticket.d
				newCost := cur.m + ticket.c

				if newCost <= m && dp[ticket.v][newCost] > newTime {
					for i := newCost; i <= m; i++ {
						if dp[ticket.v][i] > newTime {
							dp[ticket.v][i] = newTime
						} else {
							break
						}
					}

					queue = append(queue, State{ticket.v, newCost, newTime})
				}

			}
		}

		res := slices.Min(dp[n])
		if res == math.MaxInt {
			fmt.Fprint(std, "Poor KCM")
		} else {
			fmt.Fprint(std, res)
		}
	}
}

// 1
// 3 100 3
// 1 2 1 1
// 2 3 1 1
// 1 3 3 30

// 1
// 4 10 4
// 1 2 5 3
// 2 3 5 4
// 3 4 1 5
// 1 3 10 6
