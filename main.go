package main

import (
	"bufio"
	"fmt"
	"os"
)

var std *bufio.ReadWriter = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

func main() {
	defer std.Flush()

	var n int
	fmt.Fscanln(std, &n)

	// solution
}
