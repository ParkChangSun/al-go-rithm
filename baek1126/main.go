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

	var n int
	fmt.Fscan(std, &n)
	blocks := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(std, &blocks[i])
	}

	l := make([]int, 250001)
	for i := range l {
		l[i] = -1
	}
	l[0] = 0
	tl := make([]int, 250001)

	for _, v := range blocks {
		if v > 250000 {
			continue
		}

		for j := 0; j < 250001; j++ {
			temp := []int{l[j]}

			if j+v < 250000 && l[j+v] >= 0 {
				temp = append(temp, l[j+v])
			}
			if j-v >= 0 && l[j-v] >= 0 {
				temp = append(temp, l[j-v]+v)
			}
			if v-j >= 0 && l[v-j] >= 0 {
				temp = append(temp, l[v-j]+j)
			}

			tl[j] = slices.Max(temp)
		}

		copy(l, tl)
	}

	if l[0] == 0 {
		fmt.Fprint(std, -1)
	} else {
		fmt.Fprint(std, l[0])
	}
}
