package baeksolution

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Solution10250 .
func Solution10250() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	caseNumStr, _ := reader.ReadString('\n')
	caseNumTrim := strings.Trim(caseNumStr, "\n")
	caseNum, _ := strconv.Atoi(caseNumTrim)

	for i := 0; i < caseNum; i++ {
		calcRoom(reader, writer)
	}
}

func calcRoom(reader *bufio.Reader, writer *bufio.Writer) {
	caseArgsStr, _ := reader.ReadString('\n')
	caseArgsTrim := strings.Trim(caseArgsStr, "\n")
	caseArgs := strings.Split(caseArgsTrim, " ")
	var argsNum [3]int
	for i, v := range caseArgs {
		temp, _ := strconv.Atoi(v)
		argsNum[i] = temp
	}

	floor := argsNum[2] % argsNum[0]
	ho := argsNum[2] / argsNum[0]

	if floor == 0 {
		floor = argsNum[0]
	} else {
		ho++
	}

	floorStr := strconv.Itoa(floor)
	hoStr := strconv.Itoa(ho)
	if ho < 10 {
		hoStr = strings.Join([]string{"0", hoStr}, "")
	}
	ans := strings.Join([]string{floorStr, hoStr, "\n"}, "")

	writer.WriteString(ans)
	writer.Flush()
}
