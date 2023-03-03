package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

//Solution12865 No.12865
func Solution12865() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	val := make([]int, n+1)
	wt := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &wt[i], &val[i])
	}

	fmt.Fprint(writer, knapsack(val, wt, n, k))
}

// consider all subsets of items

func knapsack(val, wt []int, n, k int) int {
	store := make([][]int, n+1)
	for i := range store {
		store[i] = make([]int, k+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if wt[i] <= j {
				store[i][j] = max(store[i-1][j], val[i]+store[i-1][j-wt[i]])
			} else {
				store[i][j] = store[i-1][j]
			}
		}
	}
	return store[n][k]
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
