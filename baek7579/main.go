package main

import (
	"bufio"
	"fmt"
	"os"
)

var std = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

func main() {
	defer std.Flush()

	var n, m int
	fmt.Fscan(std, &n, &m)
	mem := make([]int, n)
	cost := make([]int, n)
	csum := 0
	for i := range n {
		fmt.Fscan(std, &mem[i])
	}
	for i := range n {
		fmt.Fscan(std, &cost[i])
		csum += cost[i]
	}

	dp := make([]int, csum+1)

	for i := range n {
		for c := csum; c >= cost[i]; c-- {
			dp[c] = max(dp[c-cost[i]]+mem[i], dp[c])
		}
	}

	for i, v := range dp {
		if v >= m {
			fmt.Fprintln(std, i)
			break
		}
	}
}
