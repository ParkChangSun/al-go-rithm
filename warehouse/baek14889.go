package baeksolution

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

//1 2 3 4 - 5 6 7 8
//2 3 4 5 - 1 6 7 8
//3 4 5 6 - 1 2 7 8

//Solution14889 .
func Solution14889() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	s := make([][]int, n)
	for i := range s {
		s[i] = make([]int, n)
	}

	//how to scan matrix
	for _, v := range s {
		for i := range v {
			var temp int
			fmt.Fscanf(reader, "%d", &temp)
			v[i] = temp
		}
		fmt.Fscanln(reader)
	}

	matchMaker := newMatchMaker()
	fmt.Fprintln(writer, matchMaker(n, 0, []int{}, s))
}

func newMatchMaker() func(int, int, []int, [][]int) int {
	min := -1
	var combination func(n, ch int, curr []int, table [][]int) int
	combination = func(n, ch int, curr []int, table [][]int) int {
		if len(curr) == n/2 {
			val := matchMaking(curr, table)
			if min < 0 || min > val {
				min = val
			}
		} else {
			for i := ch + 1; i < n; i++ {
				if !contains(curr, i) {
					combination(n, i, append(curr, i), table)
				}
			}
		}
		return min
	}
	return combination
}

func matchMaking(match []int, table [][]int) int {
	var red, blue int
	for _, v1 := range match {
		for _, v2 := range match {
			red += table[v1][v2]
		}
	}
	opposite := []int{}
	for i := 0; i < 2*len(match); i++ {
		if !contains(match, i) {
			opposite = append(opposite, i)
		}
	}
	for _, v1 := range opposite {
		for _, v2 := range opposite {
			blue += table[v1][v2]
		}
	}
	return int(math.Abs(float64(red - blue)))
}

// func contains(arr []int, target int) bool {
// 	for _, v := range arr {
// 		if v == target {
// 			return true
// 		}
// 	}
// 	return false
// }
