package main

import (
	"bufio"
	"fmt"
	"os"
)

var std = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

func main() {
	defer std.Flush()

	sum := 0

	var n, m int
	fmt.Fscan(std, &n)
	weights := make([]int, n)
	for i := range n {
		fmt.Fscan(std, &weights[i])
		sum += weights[i]
	}
	fmt.Fscan(std, &m)
	marbles := make([]int, m)
	for i := range m {
		fmt.Fscan(std, &marbles[i])
	}

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}
	dp[0][0] = true

	for i, w := range weights {
		for diff := 0; diff < sum && dp[i][diff]; diff++ {
			dp[i+1][diff] = true

			if diff+w <= sum {
				dp[i+1][diff+w] = true
			}
			if diff-w > 0 {
				dp[i+1][diff-w] = true
			}
			if w-diff > 0 {
				dp[i+1][w-diff] = true
			}
		}
	}

	for _, v := range marbles {
		if v <= sum && dp[n][v] {
			fmt.Fprint(std, "Y")
		} else {
			fmt.Fprint(std, "N")
		}
		fmt.Fprint(std, " ")
	}
}
