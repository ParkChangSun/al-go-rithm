package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n)

	// solution
}
