package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

//Solution2156 No.2156
func Solution2156() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	wines := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &wines[i])
	}
	store := make([]int, n+1)
	fmt.Fprint(writer, drinkWine(n, wines, store))
}

func drinkWine(n int, wines, store []int) int {
	if n == 1 {
		return wines[1]
	} else if n == 2 {
		return wines[1] + wines[2]
	} else if n == 0 {
		return wines[0]
	}

	if store[n] == 0 {
		store[n] = max(max(wines[n]+wines[n-1]+drinkWine(n-3, wines, store), wines[n]+drinkWine(n-2, wines, store)), drinkWine(n-1, wines, store))
	}
	return store[n]
}

// s[n] = n번째 잔까지의 최대값
// s[n] = w[n] + s[n-3] + w[n-1]
//     or w[n] + s[n-2]
// 잔을 두번 건너뛰는 경우도 존재 - s[n-1]의 값이 커서 w[n]을 마시면 3연속 금지를 지키는 데 방해되는 경우
// s[n] = max(위의 식, s[n-1])

// func max(a, b int) int {
// 	if a >= b {
// 		return a
// 	}
// 	return b
// }
