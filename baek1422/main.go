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

	var k, n int
	fmt.Fscan(std, &k, &n)
	nums := make([][]byte, k)
	for i := range k {
		fmt.Fscan(std, &nums[i])
	}

	repeat := slices.MaxFunc(nums, func(a, b []byte) int {
		if len(a) != len(b) {
			return len(a) - len(b)
		}

		for i := range len(a) {
			if a[i] != b[i] {
				return int(a[i]) - int(b[i])
			}
		}
		return 0
	})

	slices.SortFunc(nums, func(a, b []byte) int {
		for i := range len(a) + len(b) {
			ai, bi := i, i
			for ai >= len(a) {
				ai -= len(a)
			}
			for bi >= len(b) {
				bi -= len(b)
			}
			if a[ai] != b[bi] {
				return int(b[bi]) - int(a[ai])
			}
		}
		return 0
	})

	res := []byte{}
	repeatFlag := false
	for _, v := range nums {
		if string(v) == string(repeat) && !repeatFlag {
			repeatFlag = true
			for range n - k {
				res = append(res, v...)
			}
		}
		res = append(res, v...)
	}

	fmt.Fprint(std, string(res))
}
