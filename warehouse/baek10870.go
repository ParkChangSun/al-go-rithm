package baeksolution

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Solution10870 .
func Solution10870() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	nStr, _ := reader.ReadString('\n')
	nTrim := strings.Trim(nStr, "\n")
	n, _ := strconv.Atoi(nTrim)

	getFibo := fibo(n)
	ansStr := strconv.Itoa(getFibo(n))
	writer.WriteString(ansStr)
	writer.Flush()
}

func fibo(n int) func(int) int {
	fiboArray := make([]int, n+1)
	if cap(fiboArray) > 1 {
		fiboArray[1] = 1
	}
	var getFibo func(int) int
	getFibo = func(num int) int {
		if num == 0 {
			return 0
		} else if fiboArray[num] == 0 {
			return getFibo(num-1) + getFibo(num-2)
		} else {
			return fiboArray[num]
		}
	}
	return getFibo
}
