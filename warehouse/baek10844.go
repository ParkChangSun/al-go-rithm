package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

const bigNum int = 1000000000

//Solution10844 No.10844
func Solution10844() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	store := make([][]int, n+1)
	for i := range store {
		store[i] = make([]int, 10)
	}
	for i := range store[1] {
		store[1][i] = 1
	}

	var ans int
	for i := 0; i < 10; i++ {
		ans += stairNum(n, i, store)
		ans %= bigNum
	}
	fmt.Fprintln(writer, ans)
}

func stairNum(n, lastDigit int, store [][]int) int {
	if n == 1 && lastDigit == 0 {
		return 0
	}

	if store[n][lastDigit] == 0 {
		if lastDigit == 0 {
			store[n][lastDigit] = stairNum(n-1, 1, store)
		} else if lastDigit == 9 {
			store[n][lastDigit] = stairNum(n-1, 8, store)
		} else {
			store[n][lastDigit] = (stairNum(n-1, lastDigit-1, store) + stairNum(n-1, lastDigit+1, store)) % bigNum
		}
	}

	return store[n][lastDigit]
}

// n자리 a마지막 = s(n,a)
// a!=0,9 : s(n,a) = s(n-1, a-1) + s(n-1, a+1)
// a==0,9 : s(n,a) = s(n-1, a+-1)
// a!=0 : s(1,a) : 1
// a==0 : s(1,a) : 0
// 005 -> 045 , 065
// 000 -> 010

//overflow
//store[n][a] is under bigNum apparently
//but sum of two store[][] values can be over bigNum
