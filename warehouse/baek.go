package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

//Solution0 No.0
func Solution0() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

}
