package baeksolution

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Solution2869 a
func Solution2869() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	input, _ := reader.ReadString('\n')
	inputs := strings.Trim(input, "\n")
	nums := strings.Split(inputs, " ")
	a, _ := strconv.Atoi(nums[0])
	b, _ := strconv.Atoi(nums[1])
	v, _ := strconv.Atoi(nums[2])

	ans := 0
	if (v-a)%(a-b) != 0 {
		ans = (v-a)/(a-b) + 1
	} else {
		ans = (v - a) / (a - b)
	}
	ans++

	writer.WriteString(strconv.Itoa(ans))
	writer.Flush()
}
