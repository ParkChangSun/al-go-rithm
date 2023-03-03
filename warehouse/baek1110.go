package baeksolution

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Solution No.1110
func Solution1110() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, count int
	fmt.Fscanln(reader, &n)
	c := n

	for {
		next := cycle(c)
		count++
		if next == n {
			break
		}
		c = next
	}
	fmt.Fprintln(writer, count)
}

func cycle(n int) (res int) {
	var t1, t2 int
	numStr := fmt.Sprint(n)
	if n < 10 {
		t1 = 0
		t2 = n
	} else {
		splitNum := strings.Split(numStr, "")
		t1, _ = strconv.Atoi(splitNum[0])
		t2, _ = strconv.Atoi(splitNum[1])
	}
	product := t1 + t2
	var digit string
	if product >= 10 {
		digit = strings.Split(fmt.Sprint(product), "")[1]
	} else {
		digit = fmt.Sprint(product)
	}
	newNumStr := []string{fmt.Sprint(t2), digit}
	res, _ = strconv.Atoi(strings.Join(newNumStr, ""))
	return
}

// % 연산하면 나온다
