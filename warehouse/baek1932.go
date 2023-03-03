package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

//Solution1932 No.1932
func Solution1932() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	pyramid := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		a := make([]int, i)
		b := make([]interface{}, i)
		for j := 0; j < i; j++ {
			b[j] = &a[j]
		}
		fmt.Fscanln(reader, b...)
		pyramid[i] = a
	}
	store := []int{0}
	fmt.Print(maxArr(calcPyramid(pyramid, store, 1)))
}

func calcPyramid(pyramid [][]int, store []int, floor int) (newStore []int) {
	var v int
	for i := 0; i < floor; i++ {
		if i-1 < 0 {
			v = store[i]
		} else if i >= len(store) {
			v = store[i-1]
		} else {
			v = max(store[i-1], store[i])
		}
		newStore = append(newStore, pyramid[floor][i]+v)
	}
	if floor == len(pyramid)-1 {
		return
	}
	return calcPyramid(pyramid, newStore, floor+1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxArr(arr []int) (ans int) {
	for _, v := range arr {
		if ans < v {
			ans = v
		}
	}
	return
}
