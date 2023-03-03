package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

//Solution1912 No.1912
func Solution1912() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	seq := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d", &seq[i])
	}

	fmt.Fprint(writer, consum(seq, n))
}

func consum(seq []int, n int) int {
	max := seq[1]
	current := 0
	for i := 1; i <= n; i++ {
		current += seq[i]
		if max < current {
			max = current
		}
		if current < 0 {
			current = 0
		}
	}
	return max
}
