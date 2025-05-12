package main

import (
	"bufio"
	"fmt"
	"os"
)

var m = []int{}
var s string
var p = [][]bool{}

func main() {
	std := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer std.Flush()

	m = make([]int, 2500)
	p = make([][]bool, 2500)
	for i := range p {
		p[i] = make([]bool, 2500)
	}

	fmt.Fscanln(std, &s)

	for i := range s {
		m[i] = 2500
		for j := 0; j <= i; j++ {
			if isP(j, i) {
				m[i] = getmin(j, i)
			}
		}
	}

	fmt.Fprintln(std, m[len(s)-1])
}

func getmin(j, i int) int {
	if j == 0 {
		return 1
	}
	a, b := m[i], m[j-1]+1
	if a > b {
		return b
	}
	return a
}

// 012345
func isP(j, i int) bool {
	for k := 0; k <= (i-j)/2; k++ {
		a, b := j+k, i-k
		if p[a][b] {
			p[j][i] = true
			return true
		} else if s[a] != s[b] {
			return false
		}
	}

	p[j][i] = true
	return true
}
