package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX = 15000

var w, m int
var weights []int
var marbles []int
var memo [][]bool
var checked [][]bool

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &w)
	weights = make([]int, w+1)
	d1 := make([]interface{}, w)
	for i := range d1 {
		d1[i] = &weights[i+1]
	}
	fmt.Fscanln(reader, d1...)

	fmt.Fscanln(reader, &m)
	marbles = make([]int, m+1)
	d2 := make([]interface{}, m)
	for i := range d2 {
		d2[i] = &marbles[i+1]
	}
	fmt.Fscanln(reader, d2...)

	memo = make([][]bool, w+1)
	checked = make([][]bool, w+1)
	for i := 0; i <= w; i++ {
		memo[i] = make([]bool, MAX+1)
		checked[i] = make([]bool, MAX+1)
	}

	for i := 1; i <= m; i++ {
		if td(w, marbles[i]) {
			fmt.Fprint(writer, "Y")
		} else {
			fmt.Fprint(writer, "N")
		}
		fmt.Fprint(writer, " ")
	}
}

func td(wn, mw int) bool {
	if mw > MAX || wn <= 0 {
		return false
	}
	if weights[wn] == mw {
		return true
	}
	if checked[wn][mw] {
		return memo[wn][mw]
	}

	if td(wn-1, mw+weights[wn]) || td(wn-1, abs(mw-weights[wn])) || td(wn-1, mw) {
		memo[wn][mw] = true
	}
	checked[wn][mw] = true
	return memo[wn][mw]
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

// 구하려는 변수는 양쪽 팔의 차
// 그럼 변수는 각각 팔의 무게?

// 냅색 변수 남은가방무게 물건번호

// 한쪽무게 다른쪽무게 ?
// 둘수 있는 것은 n번째 구슬을
// 왼쪽, 오른쪽, 두지 않기
// 이렇게 나눌수 있다

// 현재 무게 차에 더하거나 빼거나 그냥두거나 ?

// 현재무게차 추번호

// 아니면 모든 추에 대해 더하기빼기를 다 해놓고
// 구슬들을 그 안에서 찾을수도 있을듯?

// 0번째 구슬에서 무게차 0 부터 시작

// for 무게차 0~40000?
// 1번째 구슬 w1의 0~k~40000 -> 0번째 구슬의 k-n1 or k+n1 or k
// 2번째 구슬 w2의 0~k~40000 -> 1번째 구슬의 k-n2 or k+n2 or k

// n 번째 구슬 wn의 0~k~40000 -> n-1 번째 구슬의 k-n or k+n or k
// if [n-1][k-n] or [n-1][k+n] or [n-1][k] then true

// if weiht==weight true

// 1번째 구슬에서 무게차 0+n1 빼기는 제외
// 2번째 구슬에서 무게차 n1+n2 또는 n1-n2
// 메모 값은 부울타입?

// 추로 만들 수 있는 무게의 최대값은 15000이다
// 40000은 그냥 예외값도 상정한 범위였다
