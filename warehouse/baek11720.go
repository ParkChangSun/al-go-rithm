package baeksolution

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Solution No.11720
func Solution() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	firstline, _ := reader.ReadString('\n')
	secondline, _ := reader.ReadString('\n')

	splitnums := strings.Split(secondline, "")

	length, _ := strconv.Atoi(strings.Trim(firstline, "\n"))
	sum := 0
	for i := 0; i < length; i++ {
		temp, _ := strconv.Atoi(splitnums[i])
		sum += temp
	}

	answer := strconv.Itoa(sum)
	writer.WriteString(answer)
	writer.Flush()
}
