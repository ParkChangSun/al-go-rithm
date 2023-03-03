package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

//Solution2579 No.2579
func Solution2579() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	stairs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &stairs[i])
	}

	if n == 1 {
		fmt.Fprintln(writer, stairs[1])
	} else if n == 2 {
		fmt.Fprintln(writer, stairs[1]+stairs[2])
	} else {
		a := []int{stairs[1], stairs[1], stairs[1] + stairs[2], stairs[2]}
		ans := climb(stairs, a, 3)
		fmt.Fprintln(writer, max(ans[2], ans[3]))
	}
}

//calculating n th stair
//n-2th stair and n-1th stair which is not from n-2th

func climb(stairs, state []int, floor int) []int {
	newState := []int{state[2], state[3]}
	newState = append(newState, state[3]+stairs[floor], max(state[0], state[1])+stairs[floor])
	if floor == len(stairs)-1 {
		return newState
	}
	return climb(stairs, newState, floor+1)
}

// func max(a, b int) int {
// 	if a >= b {
// 		return a
// 	}
// 	return b
// }
