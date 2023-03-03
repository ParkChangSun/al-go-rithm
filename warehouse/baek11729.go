package baeksolution

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

//Solution11729 .
func Solution11729() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)
	defer writer.Flush()

	fmt.Fprintln(writer, int(math.Pow(2, float64(n)))-1)
	hanoi(1, 3, 2, n, writer)
}

func hanoi(from, to, via, n int, writer *bufio.Writer) {
	if n == 1 {
		fmt.Fprintf(writer, "%d %d\n", from, to)
		return
	}
	hanoi(from, via, to, n-1, writer)
	fmt.Fprintf(writer, "%d %d\n", from, to)
	hanoi(via, to, from, n-1, writer)
}
