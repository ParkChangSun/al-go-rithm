package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	std := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer std.Flush()

	var n int
	fmt.Fscanln(std, &n)

	if n%2 == 1 {
		fmt.Fprintln(std, 0)
		return
	}

	memo := make([]int, n+1)
	memo[0] = 1
	memo[2] = 3
	for i := 4; i <= n; i += 2 {
		acc := 0
		for j := 0; j <= i-4; j += 2 {
			acc += 2 * memo[j]
		}
		memo[i] = 3*memo[i-2] + acc
	}
	fmt.Fprintln(std, memo[n])
}
