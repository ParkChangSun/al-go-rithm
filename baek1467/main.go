package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
)

var std = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

func main() {
	defer std.Flush()

	var numbers, targets []byte
	fmt.Fscan(std, &numbers, &targets)

	reduced := map[byte]int{}
	for _, v := range numbers {
		reduced[v]++
	}
	for _, v := range targets {
		reduced[v]--
		if reduced[v] == 0 {
			delete(reduced, v)
		}
	}

	res := []byte{}

	for len(reduced) > 0 {
		for i := byte('9'); i > byte('0'); i-- {
			if _, o := reduced[i]; !o {
				continue
			}
			targetIndex := slices.Index(numbers, i)

			contains := true
			temp := maps.Clone(reduced)
			for _, v := range numbers[targetIndex:] {
				temp[v]--
			}
			for _, v := range temp {
				if v > 0 {
					contains = false
					break
				}
			}
			if !contains {
				continue
			}

			res = append(res, i)
			numbers = numbers[targetIndex+1:]
			reduced[i]--
			if reduced[i] == 0 {
				delete(reduced, i)
			}
			break
		}
	}

	fmt.Fprint(std, string(res))
}
