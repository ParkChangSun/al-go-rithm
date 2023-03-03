package baeksolution

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Solution2447 .
func Solution2447() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	nStr, _ := reader.ReadString('\n')
	nTrim := strings.Trim(nStr, "\n")
	n, _ := strconv.Atoi(nTrim)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			printStar(i, j, n, writer)
		}
		writer.WriteString("\n")
	}
	writer.Flush()
}

func printStar(rowNum int, colNum int, size int, writer *bufio.Writer) {
	//0,0
	// rownum%3==1

	//345 121314 212223
	//1 4 7
	// rownum%/3%3==1

	//9_17
	//3~5
	//rownum/9%3==1

	if size <= 0 {
		writer.WriteString("*")
		return
	}
	if (rowNum/size)%3 == 1 && (colNum/size)%3 == 1 {
		writer.WriteString(" ")
	} else {
		printStar(rowNum, colNum, size/3, writer)
	}
}
