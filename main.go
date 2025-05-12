package main

import (
	"bufio"
	"os"
)

func main() {
	std := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer std.Flush()

	// solution
}
