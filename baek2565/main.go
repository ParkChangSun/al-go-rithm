package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var n int
var l, memo []int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n)

	l = make([]int, n+1)
	memo = make([]int, n+1)

	lines := make([][]int, n)
	for i := 0; i < n; i++ {
		lines[i] = make([]int, 2)
		fmt.Fscanln(reader, &lines[i][0], &lines[i][1])
	}
	sort.Slice(lines, func(i, j int) bool {
		return lines[i][0] < lines[j][0]
	})
	for i := range lines {
		l[i+1] = lines[i][1]
	}

	prune(n)
	sort.Ints(memo)
	fmt.Fprintln(writer, n-memo[n])
}

func prune(n int) int {
	if memo[n] != 0 {
		return memo[n]
	}

	acc := 0
	for i := 1; i < n; i++ {
		if t := prune(i); l[i] < l[n] && acc < t {
			acc = t
		}
	}
	memo[n] = acc + 1
	return acc + 1
}

// 리스트에서 중간에 감소하는 값이 없어야 함->증가수열
// N번째에서 가장 긴 증가하는 수열
// lis(n) n까지의 가장 긴 수열
// N보다 작은 수가 있는 자리에서 +1하는 것? 아님
// lis(N-1)+1? 아님
// N보다 작은 수 중 가장 시퀀스가 높은 수+1
// 1 99 100 70 80 10 82 3 4 5 2 90 89

// 최대값 찾아야 하느거
// 리스트 만들어 놓고 값 리스트에 넣는것
// l[i+1] = lines[i][1] 굳이 안해도 되지만 특히나 바텀업일때는
