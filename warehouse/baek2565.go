package baeksolution

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

//Solution2565 No.2565
func Solution2565() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	input := make([][]int, n)
	for i := 0; i < n; i++ {
		input[i] = make([]int, 2)
		fmt.Fscanf(reader, "%d %d\n", &input[i][0], &input[i][1])
	}
	sort.SliceStable(input, func(i, j int) bool {
		return input[i][0] < input[j][0]
	})
	seq := make([]int, n+1)
	for i, v := range input {
		seq[i+1] = v[1]
	}

	store := make([]int, n+1)
	for i := 1; i <= n; i++ {
		store[i] = 1
	}
	lis(1, seq, store)

	fmt.Fprint(writer, n-maxArr(store))
}

// 8 2 9 1 4 6 7 10
// x   x x

func lis(n int, seq, store []int) {
	if n == len(seq) {
		return
	}
	for i := 1; i <= n; i++ {
		if seq[n] > seq[i] && store[n] <= store[i] {
			store[n] = store[i] + 1
		}
	}
	lis(n+1, seq, store)
}

// func maxArr(arr []int) (ans int) {
// 	for _, v := range arr {
// 		if v > ans {
// 			ans = v
// 		}
// 	}
// 	return
// }
