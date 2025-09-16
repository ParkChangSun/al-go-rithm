package main

import (
	"bufio"
	"fmt"
	"os"
)

var std = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

type Stuff struct {
	w, v int
}

func main() {
	defer std.Flush()

	var n, m int
	fmt.Fscan(std, &n, &m)

	stuff := make([]Stuff, 0)
	for i := 0; i < n; i++ {
		var w, v, k int
		fmt.Fscan(std, &w, &v, &k)

		for k > 0 {
			num := 0
			for i := 1; k-1<<i+1 >= 0; i++ {
				stuff = append(stuff, Stuff{w * 1 << num, v * 1 << num})
				num++
			}
			k = k - 1<<num + 1
		}
	}

	dp := make([]int, m+1)
	for _, v := range stuff {
		for i := m; i-v.w >= 0; i-- {
			dp[i] = max(dp[i], dp[i-v.w]+v.v)
		}
	}

	fmt.Fprint(std, dp[m])
}
