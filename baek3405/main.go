package main

import (
	"bufio"
	"fmt"
	"os"
)

var std *bufio.ReadWriter = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

var matrix [][]int
var sum []int
var checked []bool
var n int

func main() {
	defer std.Flush()

	var t int
	fmt.Fscanln(std, &t)

	for i := 0; i < t; i++ {
		fmt.Fscanln(std, &n)
		matrix = [][]int{}
		for i := 0; i < n; i++ {
			matrix = append(matrix, make([]int, n))
			for j := 0; j < n; j++ {
				fmt.Fscan(std, &matrix[i][j])
			}
		}

		sum = []int{}
		for _, v := range matrix {
			debtSum := 0
			for _, debt := range v {
				debtSum += debt
			}
			sum = append(sum, debtSum)
		}

		checked = make([]bool, 1<<n)

		solve(0)

		standing := false
		for i := 0; i < n; i++ {
			cur := ((1 << n) - 1) ^ (1 << i)
			if checked[cur] {
				fmt.Fprintf(std, "%d ", i+1)
				standing = true
			}
		}
		if !standing {
			fmt.Fprint(std, 0)
		}
		fmt.Fprintln(std)

		fmt.Fscanln(std)
	}
}

func solve(mask int) {
	for i, v := range sum {
		if ((1<<i)&mask) == 0 && v > 0 && !checked[mask+(1<<i)] {
			// i번째 왕국이 파산 가능
			for j := 0; j < n; j++ {
				sum[j] -= matrix[j][i]
			}
			solve(mask + (1 << i))
			for j := 0; j < n; j++ {
				sum[j] += matrix[j][i]
			}
		}
	}
	checked[mask] = true
}

// 1111 = 10000 - 1
// 1110 -> 0001
// 1101 -> 0010

// 7
// 5
// 0 -1 3 -2 1
// 1 0 1 4 -4
// -3 -1 0 4 -1
// 2 -4 -4 0 3
// -1 4 1 -3 0
// 5
// 0 -1 3 -2 1
// 1 0 1 4 -3
// -3 -1 0 -3 2
// 2 -4 3 0 3
// -1 3 -2 -3 0
// 1
// 0
// 2
// 0 0
// 0 0
// 2
// 0 1
// -1 0
// 3
// 0 -3 1
// 3 0 -2
// -1 2 0
// 3
// 0 1 0
// -1 0 1
// 0 -1 0
