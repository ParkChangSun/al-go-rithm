package baeksolution

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Solution1712 .
func Solution1712() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	line, _ := reader.ReadString('\n')
	lineTrim := strings.Trim(line, "\n")
	inputString := strings.Split(lineTrim, " ")

	var inputs [3]int
	var ans int
	for i, v := range inputString {
		temp, _ := strconv.Atoi(v)
		inputs[i] = temp
	}
	if inputs[1] >= inputs[2] {
		ans = -1
	} else {
		ans = (inputs[0] / (inputs[2] - inputs[1])) + 1
	}

	ansStr := strconv.Itoa(ans)
	writer.WriteString(ansStr)
	writer.Flush()
}
