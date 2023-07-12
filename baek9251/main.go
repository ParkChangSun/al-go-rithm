package main

import (
	"bufio"
	"fmt"
	"os"
)

var s1, s2 string
var memo [][]int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &s1)
	fmt.Fscanln(reader, &s2)
	memo = make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		memo[i] = make([]int, len(s2)+1)
		for j := 0; j <= len(s2); j++ {
			memo[i][j] = -1
		}
	}

	fmt.Fprintln(writer, lcs(len(s1), len(s2)))
}

func lcs(n, k int) int {
	if n == 0 || k == 0 {
		return 0
	}
	if memo[n][k] != -1 {
		return memo[n][k]
	}

	if s1[n-1] == s2[k-1] {
		memo[n][k] = lcs(n-1, k-1) + 1
	} else {
		ns1 := lcs(n-1, k)
		ns2 := lcs(n, k-1)
		if ns1 > ns2 {
			memo[n][k] = ns1
		} else {
			memo[n][k] = ns2
		}
	}

	return memo[n][k]
}

// s1 n 번째까지 s2 k 번째까지 2차원 배열?
// if s1[n]==s2[k] memo[n][k]=memo + 1
// else memo[n][k]=memo

// ACAYKP
// CAPCAK
