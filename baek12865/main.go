package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, k int
var stuff [][2]int
var memo [][]int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n, &k)
	stuff = make([][2]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &stuff[i][0], &stuff[i][1])
	}

	memo = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			memo[i][j] = -1
		}
	}

	fmt.Fprintln(writer, pack(n, k))
}

func pack(n, w int) int {
	if n == 0 {
		return 0
	}
	if memo[n][w] != -1 {
		return memo[n][w]
	}

	if stuff[n][0] > w {
		memo[n][w] = pack(n-1, w)
	} else {
		putn := stuff[n][1] + pack(n-1, w-stuff[n][0])
		notn := pack(n-1, w)
		if putn > notn {
			memo[n][w] = putn
		} else {
			memo[n][w] = notn
		}
	}
	return memo[n][w]
}

// 풀이
// n개의 물건들을 w만큼의 가방에 넣었을때 최대 밸류
// n을 넣고 w-n 가방의 최대 밸류 n+pack(n-1,w-n)
// n을 안넣고 w 가방의 최대 pack(n-1,w)

// 가방 무게가 t인 가방의 최대 밸류 + N물건 -> 가방 무게가 t+N인 가방의 최대 밸류
// 이런 식으로 문제를 나눴네
// 나는 N에서의 최대값이 가방의/그 N에서의 최대값인 줄로만 알았는데

// n이 1부터 시작하는거 실수
// init -1로 하는 팁

// N번째 물건에서의 최대값은
// N번째 물건을 넣을 수 있는 밸류가 가장 큰 지점에서 N번째 물건을 넣기? 100 1 같은 비효율적인 물건을 넣을수도
// 밸류가 가장 큰 지점에서 결정하는 것은 맞는데, 이 물건을 안넣을수도 있나?
// max weight 1000
// 10 1
// 20 2
// 800 4
// 900 0
// 30 3
// 800 4
// 200 100
// 100 50
// 200에서 최대 밸류는 1+2+3 인데 이 알고리즈대로면 1+2가 최대가 된다
// 모든 N에 대해서 전부 계산해서 비교해야 하나?

// 정렬해야 하나?
