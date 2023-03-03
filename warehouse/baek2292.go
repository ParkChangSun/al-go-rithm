package baeksolution

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Solution2292 correct
func Solution2292() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	input, _ := reader.ReadString('\n')
	num := strings.Trim(input, "\n")
	n, _ := strconv.Atoi(num)

	ansInt := getLineNum(n)
	ansStr := strconv.Itoa(ansInt)
	writer.WriteString(ansStr)
	writer.Flush()
}

func getLineNum(x int) (lineNum int) {
	//1 7 19 37
	max := 1
	lineNum = 1
	for ; ; lineNum++ {
		if x > max {
			max += lineNum * 6
		} else {
			return
		}
	}
}
