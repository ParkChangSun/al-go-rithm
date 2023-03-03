package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

//Solution1904 No.1904
func Solution1904() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

}

//n=1 1
//n=2 00 11
//n=3 001 100 111
//n=4 0000 0011 1001 1100 1111
//n=5 00001 00100 00111 10000 10011 11001 11100 11111
