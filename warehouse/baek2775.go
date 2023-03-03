package baeksolution

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Solution2775 .
func Solution2775() {
	//0층 1 2            3          4.. b
	//1층 1 1+2          1+2+3        ...  (1~b)SigmaN b^2/2 + b/2
	//2층 1 1 + 1+2      1+ 1+2 1+2+3 (1~b)Sigma(1~b)SigmaN
	//3층 1 1 + 1+1+2    1 + 1+1+2 + 1+1+2+1+2+3

	//f층 h호 = f-1층 시그마
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	casesStr, _ := reader.ReadString('\n')
	casesTrim := strings.Trim(casesStr, "\n")
	cases, _ := strconv.Atoi(casesTrim)

	for i := 0; i < cases; i++ {
		floorStr, _ := reader.ReadString('\n')
		hoStr, _ := reader.ReadString('\n')
		floorTrim := strings.Trim(floorStr, "\n")
		hoTrim := strings.Trim(hoStr, "\n")
		floor, _ := strconv.Atoi(floorTrim)
		ho, _ := strconv.Atoi(hoTrim)

		ans := calc(floor, ho)
		ansStr := strconv.Itoa(ans)
		writer.WriteString(ansStr)
		writer.WriteString("\n")

	}
	writer.Flush()
}

func calc(floor int, ho int) (ans int) {
	if floor == 0 {
		ans = ho
	} else {
		for i := 1; i <= ho; i++ {
			ans += calc(floor-1, i)
		}
	}
	return
}
