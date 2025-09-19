package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var std = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

type Ticket struct {
	u, v, c, d int
}

type Item struct {
	n, m int
}

func main() {
	defer std.Flush()

	var t int
	fmt.Fscan(std, &t)

	for i := 0; i < t; i++ {
		var n, m, k int
		fmt.Fscan(std, &n, &m, &k)
		tickets := make([][]*Ticket, n+1)
		for i := 0; i <= n; i++ {
			tickets[i] = make([]*Ticket, n+1)
		}
		for i := 0; i < k; i++ {
			temp := &Ticket{}
			fmt.Fscan(std, &temp.u, &temp.v, &temp.c, &temp.d)
			tickets[temp.u][temp.v] = temp
		}

		dp := make([][]int, n+1)
		for i := range dp {
			dp[i] = make([]int, m+1)
			for j := range dp[i] {
				dp[i][j] = m + 1
			}
		}
		dp[1][0] = 0
		queue := []Item{{1, 0}}

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			for _, v := range tickets[cur.n] {
				if v == nil || cur.m+v.c > m {
					continue
				}
				if dp[v.v][cur.m+v.c] > dp[cur.n][cur.m]+v.d {
					dp[v.v][cur.m+v.c] = dp[cur.n][cur.m] + v.d
					queue = append(queue, Item{v.v, cur.m + v.c})
				}
			}
		}
		// fmt.Print(dp[n])

		res := slices.Min(dp[n])
		if res > m {
			fmt.Fprint(std, "Poor KCM")
		} else {
			fmt.Fprint(std, res)
		}
	}
}
