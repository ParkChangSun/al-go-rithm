package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var words []string
var n int

func main() {
	std := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer std.Flush()

	fmt.Fscanln(std, &n)
	words = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(std, &words[i])
	}

	ans := ""

	// for {
	// 	next, point := dictchar()
	// 	if next == -1 {
	// 		break
	// 	}
	// 	ans = fmt.Sprintf("%s%s", ans, words[next][0:point+1])
	// 	words[next] = words[next][point+1:]
	// }

	for {
		next := pqueue()
		if next == 0 {
			break
		}
		ans = fmt.Sprintf("%s%s", ans, string(next))
	}
	fmt.Fprint(std, ans)
}

// func dictchar() (int, int) {
// 	validIndexes := []bool{}
// 	for i := 0; i < n; i++ {
// 		validIndexes = append(validIndexes, true)
// 	}

// 	level := 0
// 	stop := false

// 	lowestPoint := 0
// 	lowestChar := byte(255)

// 	previousSmallest := []int{}

// 	for {
// 		hold := byte(255)
// 		smallest := []int{}
// 		for i, w := range words {
// 			if !validIndexes[i] || len(w) <= level {
// 				continue
// 			}

// 			if foreLetter := byte(w[level]); hold > foreLetter {
// 				hold = foreLetter
// 				smallest = []int{i}

// 			} else if hold == foreLetter {
// 				smallest = append(smallest, i)
// 			}
// 		}

// 		if len(smallest) == 1 {
// 			return smallest[0], lowestPoint
// 		} else if len(smallest) == 0 {
// 			if len(previousSmallest) == 0 {
// 				return -1, 0
// 			} else {
// 				return previousSmallest[0], lowestPoint
// 			}
// 		}

// 		if !stop && lowestChar >= hold {
// 			lowestChar = hold
// 			lowestPoint = level
// 		} else {
// 			stop = true
// 		}

// 		for i, _ := range validIndexes {
// 			if !slices.Contains(smallest, i) {
// 				validIndexes[i] = false
// 			}
// 		}

// 		previousSmallest = smallest
// 		level++
// 	}
// }

func pqueue() byte {
	slices.SortFunc(words, func(a string, b string) int {
		if b != "" && strings.HasPrefix(a, b) {
			return -1
		}
		if a != "" && strings.HasPrefix(b, a) {
			return 1
		}
		return strings.Compare(a, b)
	})
	if words[0] == "" {
		words = words[1:]
	}
	if len(words) == 0 {
		return 0
	}
	r := byte(words[0][0])
	words[0] = words[0][1:]
	return r
}
