package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

//Solution9251 No.9251
func Solution9251() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b string
	fmt.Fscan(reader, &a, &b)

	fmt.Fprint(writer, lcs(a, b))
}

// ACAYKP
// CAPCAK
// if common char found
// l(a[1...n-1], b[1...m-1]) = l(a[1...n-2], b[1...m-2]) + 1
// if common char not found
// l(a[1...n-1], b[1...m-1]) = max l(a[1...n-1], b[1...m-2]) l(a[1...n-2], b[1...m-1])

func lcs(a, b string) int {
	store := make([][]int, len(a)+1)
	for i := 0; i <= len(a); i++ {
		store[i] = make([]int, len(b)+1)
	}
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				store[i][j] = store[i-1][j-1] + 1
			} else {
				store[i][j] = max(store[i-1][j], store[i][j-1])
			}
		}
	}
	return store[len(a)][len(b)]
}

// func max(a, b int) (ans int) {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
