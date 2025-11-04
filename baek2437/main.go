package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var std = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

func main() {
	defer std.Flush()

	n := 0
	fmt.Fscan(std, &n)
	weights := make([]int, n)
	for i := range n {
		fmt.Fscan(std, &weights[i])
	}

	slices.Sort(weights)

	sum := 0

	for _, w := range weights {
		if w > sum+1 {
			break
		}
		sum += w
	}

	fmt.Fprint(std, sum+1)
}
