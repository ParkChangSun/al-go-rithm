package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

//Solution2750 .
func Solution2750() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int

	fmt.Fscanf(reader, "%d\n", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d\n", &arr[i])
	}
	for i := 1; i < n; i++ {
		j := i - 1
		temp := arr[i]
		for ; j >= 0 && arr[j] > temp; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
	}
	for _, v := range arr {
		fmt.Fprintf(writer, "%d\n", v)
	}
}
