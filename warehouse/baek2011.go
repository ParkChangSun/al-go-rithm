package baeksolution

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//4 = 1 2 1 - 3
//5 = 1 2 1 1 - 4
//6 = 1 2 2 1
//7 = 1 2 2 1 1
//8 = 1 2 2 2 1
//9 = 1 2 3 2 1 - 5
//10 = 1 2 3 2 1 1
//11 = 1 2 3 2 2 1
//12 = 1 2 3 3 2 1 - 6
//13 = 1 2 3 3 2 1 1
//14 = 1 2 3 3 2 2 1
//15 = 1 2 3 3 3 2 1
//16 = 1 2 3 4 3 2 1 - 7
//17 = 1 2 3 4 3 2 1 1
//18
//19
//20 = 1 2 3 4 4 3 2 1 - 8
//21 = 1 2 3 4 4 3 2 1 1
//22
//23
//24
//25 = 1 2 3 4 5 4 3 2 1 - 9

//Solution1011 .
func Solution1011() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	caseNumStr, _ := reader.ReadString('\n')
	caseNumTrim := strings.Trim(caseNumStr, "\n")
	caseNum, _ := strconv.Atoi(caseNumTrim)
	for i := 0; i < caseNum; i++ {
		alpha(reader, writer)
	}
}

func alpha(reader *bufio.Reader, writer *bufio.Writer) {
	testStr, _ := reader.ReadString('\n')
	testTrim := strings.Trim(testStr, "\n")
	testSpl := strings.Split(testTrim, " ")
	a, _ := strconv.Atoi(testSpl[0])
	b, _ := strconv.Atoi(testSpl[1])
	distance := b - a

	var ans int
	for i := 1; ; i++ {
		if distance <= i*i {
			if (i*i)-distance < i {
				ans = (i * 2) - 1
			} else {
				ans = (i * 2) - 2
			}
			break
		}
	}

	ansStr := strconv.Itoa(ans)
	ansEnter := strings.Join([]string{ansStr, "\n"}, "")
	writer.WriteString(ansEnter)
	writer.Flush()
}
