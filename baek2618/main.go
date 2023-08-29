package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var n, w int
var cases, memo [][]int

func main() {
	std := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer std.Flush()

	fmt.Fscanln(std, &n)
	fmt.Fscanln(std, &w)

	cases = make([][]int, w+1)
	for i := 1; i <= w; i++ {
		cases[i] = make([]int, 2)
		fmt.Fscanln(std, &cases[i][0], &cases[i][1])
	}

	memo = make([][]int, w+1)
	for i := 0; i <= w; i++ {
		memo[i] = make([]int, w+1)
		for j := 0; j <= w; j++ {
			memo[i][j] = -1
		}
	}

	fmt.Fprintln(std, solve(0, 0))

	track(std)
}

func solve(a, b int) int {
	if max(a, b) == w {
		memo[a][b] = 0
		return 0
	}

	if memo[a][b] != -1 {
		return memo[a][b]
	}

	caseNum := max(a, b) + 1

	aMove := solve(caseNum, b) + dist(a, caseNum)
	bMove := solve(a, caseNum) + dist(caseNum, b)
	memo[a][b] = min(aMove, bMove)
	return memo[a][b]
}

func track(std *bufio.ReadWriter) {
	aPos, bPos := 0, 0

	for i := 1; i <= w; i++ {
		if memo[aPos][bPos]-dist(aPos, i) == memo[i][bPos] {
			aPos = i
			fmt.Fprintln(std, 1)
			continue
		}

		if memo[aPos][bPos]-dist(i, bPos) == memo[aPos][i] {
			bPos = i
			fmt.Fprintln(std, 2)
			continue
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dist(n1, n2 int) int {
	p1 := cases[n1]
	p2 := cases[n2]
	if n1 == 0 {
		p1 = []int{1, 1}
	}
	if n2 == 0 {
		p2 = []int{n, n}
	}
	return int(math.Abs(float64(p1[0]-p2[0])) + math.Abs(float64(p1[1]-p2[1])))
}
