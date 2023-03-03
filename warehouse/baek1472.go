package baeksolution

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Solution1472 .
func Solution1472() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	nStr := strconv.Itoa(n)
	arr := make([]byte, len(nStr))
	//baekjoon i/o error!
	// if answer array declared like this : arr := make([]byte, 10)
	// when input is under 10digits - string(arr) seems like correct
	// but there's zero values(null character) so incorrect in baekjoon
	for i := 0; i < len(nStr); i++ {
		v := nStr[i]
		if i == 0 {
			arr[0] = v
		} else {
			j := i - 1
			for ; j >= 0 && v > arr[j]; j-- {
				arr[j+1] = arr[j]
			}
			arr[j+1] = v
		}
	}
	ans := string(arr)
	fmt.Fprintln(writer, ans)
}
