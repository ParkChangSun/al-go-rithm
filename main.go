package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int

func main() {
	std := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer std.Flush()

	fmt.Fscanln(std, &n)

	// solution
}
