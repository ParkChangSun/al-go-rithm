package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, k int
var coins, memo []int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n, &k)
	coins = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &coins[i])
	}
	memo = make([]int, k+1)
	memo[0] = 1

	fmt.Fprintln(writer, coin(k))
}

func coin(v int) int {
	for _, v := range coins {
		for i := 0; i <= k; i++ {
			if d := i - v; d >= 0 {
				memo[i] += memo[d]
			}
		}
	}
	return memo[k]
}

// 1 2 5

// 0 1 2 3 4
// 1 1 2 2 3

// 1 = 0
// 2 = 1 0
// 3 = 2 1 -> 2 1 (2>1이므로 1빼기)
// 4 = 3 2 -> 3 2 (1빼기)
// 5 = 4 3 0 -> (5>2) 4
// 6 = 5
// 7 = 6
// 8=7
// 9=8
// 10=10
// 10 1111111111 111111112 11111122 1111222 112222/ 22222/ 111115 11125 1225/ 55
// 9 111111111 11111112 1111122 111222 12222/ 11115 1125 225
// 8 11111111 1111112 111122 11222 2222 1115 125
// 7 1111111 111112 11122 1222 115 25
// 6 111111 11112 1122 222 15
// 5 11111 1112 122 5
// 4 1111 112 22
// 3 111 12
// 2 11 2
// 1 1
// 2단위로 2를 이용한 경우가 하나 는다
// 5단위로 마찬가지

// 겹치는 것을 어떻게 제거할 것인가?
// 겹치지 않게 계산하는 방법이 있는가?

// k-coin 의 경우를 모두 더한다? 아님 겹치는 경우가 있음

// coin(k) += coin(k-coin[i])

// 먼저 1동전을 사용해

// 메모의 어떤 값을 한번에 정답을 구하는 것이 아니라 여러 번 돌면서 답을 쌓아나가는 방식

// 탑다운이라면 변수는 가치, 그리고 사용한 동전의 가지수
// memo[3][10]
