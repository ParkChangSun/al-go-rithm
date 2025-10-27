package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var std = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

func main() {
	defer std.Flush()

	var n, targets []byte
	fmt.Fscan(std, &n, &targets)

	for _, target := range targets {
		hold := -1
		for i, number := range n {
			if target == number {
				if i == len(n)-1 {

					hold = i

					break
				}

				if n[i+1] < target {

					hold = i

				} else if n[i+1] == target {
					if hold == -1 {
						hold = i
					}
				} else {
					hold = i
					break
				}
			}
		}

		n = slices.Delete(n, hold, hold+1)
		// fmt.Println(string(target), string(n))
	}

	fmt.Fprint(std, string(n))
}
