package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

//Solution0 No.1747
func Solution1747() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	if n <= 2 {
		fmt.Fprint(writer, 2)
		return
	}
	if n%2 == 0 {
		n++
	}
	for {
		if isPrime(n) && isPalin(fmt.Sprint(n)) {
			break
		}
		n += 2
	}

	fmt.Fprint(writer, n)
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPalin(nstr string) bool {
	for i := 0; i < len(nstr)/2; i++ {
		if nstr[i] != nstr[len(nstr)-i-1] {
			return false
		}
	}
	return true
}
