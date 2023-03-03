package baeksolution

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Solution15649 .
func Solution15649() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)
	backTrack(n, m, 0, []int{}, writer)
}

// just make func contains()
// or exclude numbers that are used
func backTrack(n, m, place int, currSol []int, writer *bufio.Writer) {
	if m == place {
		fmt.Fprintln(writer, printArr(currSol))
	}
	for i := 1; i <= n; i++ {
		if !contains(currSol, i) {
			backTrack(n, m, place+1, append(currSol, i), writer)
		}
	}
}

func contains(arr []int, target int) bool {
	for _, v := range arr {
		if target == v {
			return true
		}
	}
	return false
}

func printArr(arr []int) string {
	return strings.Trim(fmt.Sprint(arr), "[]")
}
