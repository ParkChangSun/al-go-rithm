package main

import (
	"bufio"
	"fmt"
	"os"
)

const P = 1_000_000_007
const MAXLIM = 4_000_000

var n, v1, v2 int
var f [MAXLIM + 1]int64

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n)

	factorial()

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &v1, &v2)
		fmt.Fprintln(writer, binomial(v1, v2))
	}
}

func factorial() {
	f[0] = 1
	for i := 1; i <= MAXLIM; i++ {
		f[i] = f[i-1] * int64(i) % P
	}
}

func inverse(a int64) int64 {
	ans := int64(1)
	for exp := P - 2; exp != 0; exp /= 2 {
		if exp%2 == 1 {
			ans = ans * a % P
		}
		a = a * a % P
	}
	return ans
}

func binomial(n, k int) int64 {
	return f[n] * inverse(f[k]*f[n-k]%P) % P
}

// inverse for문 exp 선언이랑 나누기를 for첫줄안에 넣을수있다.
