package baeksolution

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Solution1193 ...
func Solution1193() {
	// an = n(n+1)/2
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	numStr, _ := reader.ReadString('\n')
	num := strings.Trim(numStr, "\n")
	numInt, _ := strconv.Atoi(num)

	var lineNum, gap int

	for i := 1; ; i++ {
		sum := i * (i + 1) / 2
		if numInt <= sum {
			lineNum = i
			gap = sum - numInt
			break
		}
	}

	var mom, son int
	if lineNum%2 == 1 {
		mom = lineNum - gap
		son = 1 + gap
	} else {
		mom = 1 + gap
		son = lineNum - gap
	}

	momStr := strconv.Itoa(mom)
	sonStr := strconv.Itoa(son)
	fraction := []string{sonStr, "/", momStr}

	writer.WriteString(strings.Join(fraction, ""))
	writer.Flush()
}
