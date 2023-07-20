package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b, c int
	fmt.Fscanln(reader, &a, &b, &c)

	acc := int64(1)
	A := int64(a)
	C := int64(c)
	for b != 0 {
		if b%2 == 1 {
			acc = (acc * A) % C
		}
		A = (A * A) % C
		b = b / 2
	}
	fmt.Fprintln(writer, acc)
}

// func dtob(n int) [31]bool {
// 	l := [31]bool{}
// 	for i := 0; n != 0; i++ {
// 		fmt.Println(n % 2)
// 		l[i] = n%2 == 1
// 		n = n / 2
// 	}
// 	return l
// }

// 3 1 7 -> 3
// 3 2 7 - 2
// 3 3 7 - 6
// 3 4 7 - 4
// 3 5 7 - 5

// 2147483647

// r2l 이다
// 10 -> 2 변환시
// 제일 처음 나누기는 k0에 해당한다
