package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

// type counted []int

//Solution1463 No.1463
func Solution1463() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	countedArr := make([]int, n+1)
	ans := countToOne(n, countedArr)
	fmt.Print(ans)
}

// M[n] means answer of n
// M[n] = M[n/3]+1 or M[n/2]+1 or M[n-1]+1
//      = min(M[n/3], M[n/2], M[n-1]) + 1
// M[1] = 0
func countToOne(n int, arr []int) (a int) {
	if n == 1 {
		return 0
	}
	if arr[n] != 0 {
		return arr[n]
	}

	a = countToOne(n-1, arr)
	if n%3 == 0 {
		a = min(a, countToOne(n/3, arr))
	}
	if n%2 == 0 {
		a = min(a, countToOne(n/2, arr))
	}
	a++
	arr[n] = a
	return
}

// func min(a, b int) int {
// 	if a >= b {
// 		return b
// 	}
// 	return a
// }
