package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var N int
var W [][]int
var memo [][]int
var finish int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &N)
	W = make([][]int, N)
	for i := 0; i < N; i++ {
		W[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Fscan(reader, &W[i][j])
		}
	}
	var mask = 1 << N
	finish = mask - 1
	memo = make([][]int, N)
	for i := 0; i < N; i++ {
		memo[i] = make([]int, mask)
	}

	fmt.Fprintln(writer, route(0, 1))
}

func route(current int, visited int) int {
	if memo[current][visited] != 0 {
		return memo[current][visited]
	}
	if visited == finish {
		if W[current][0] == 0 {
			return math.MaxInt32 / 2
		}
		return W[current][0]
	}

	ans := math.MaxInt32 / 2
	for i, bit := 0, 1; i < N; i, bit = i+1, bit<<1 {
		if visited&bit == bit || W[current][i] == 0 {
			continue
		}
		ans = min(ans, route(i, visited|bit)+W[current][i])
	}

	memo[current][visited] = ans
	return memo[current][visited]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// 0번부터 시작해서 모든 도시를 다 거치고 나면 그게 정답이다.
// 0번에서 시작했기 때문에 visited는 1을 준다.
// 이 함수는 현재 상태에서 가질 수 있는 비용을 계산한다.

// visited가 다 차면 정답이다.

// inf+inf 가 되는 경우가 있다. ???
// 모든 지점을 순회한 후이더라도 마지막 지점에서 시작점으로 가는 길이 항상 있는 것이 아니다.
