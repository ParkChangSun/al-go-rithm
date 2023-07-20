package main

import (
	"bufio"
	"fmt"
	"os"
)

const p int64 = 1_000_000_007

var (
	n, k int
	f    []int64
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n, &k)

	f = make([]int64, n+1)
	factorial()

	fmt.Fprintln(writer, binomial(n, k))
}

func binomial(a, b int) int64 {
	return (f[a] * inverse((f[b]*f[a-b])%p)) % p
}

func factorial() {
	f[0] = 1
	for i := 1; i <= n; i++ {
		f[i] = (int64(i) * f[i-1]) % p
	}
}

func inverse(a int64) int64 {
	t := p - 2
	ans := int64(1)
	acc := int64(a)
	for t != 0 {
		if t%2 == 1 {
			ans = (ans * acc) % p
		}
		acc = (acc * acc) % p
		t /= 2
	}
	return ans
}

// 팩토리얼 굳이 저장안해도 된다
// 메모리가 극한으로 적을때 방법
// ans:=1
// ans=ans*i
// return ans
