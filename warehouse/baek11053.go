package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

//Solution11053 No.11053
func Solution11053() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	sequence := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d", &sequence[i])
	}

	store := make([]int, n+1)
	for i := 1; i <= n; i++ {
		store[i] = 1
	}
	longPart(1, sequence, store)

	fmt.Fprint(writer, maxArr(store))
}

// s[n] = n을 포함하는 가장 긴 수열
// s[n] = n보다 작은 수가 최대값인 수열 중 가장긴수열 + 1

func longPart(n int, seq, store []int) {
	if n == len(seq) {
		return
	}
	for i := 1; i < n; i++ {
		if seq[n] > seq[i] && store[n] <= store[i] {
			store[n] = store[i] + 1
		}
	}
	longPart(n+1, seq, store)
}

// func maxArr(arr []int) (ans int) {
// 	for _, v := range arr {
// 		if v > ans {
// 			ans = v
// 		}
// 	}
// 	return
// }
