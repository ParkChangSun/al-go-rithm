package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

//Solution10872 .
func Solution10782() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	fmt.Fprintf(writer, "%d", fac(n))
	writer.Flush()
}

func fac(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fac(n-1)
}
