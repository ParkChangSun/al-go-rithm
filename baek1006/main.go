package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var std *bufio.ReadWriter = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

var min int
var n, w int
var inner, outer []int

func main() {
	defer std.Flush()

	var t int
	fmt.Fscanln(std, &t)

	for i := 0; i < t; i++ {
		min = math.MaxInt

		fmt.Fscanln(std, &n, &w)
		inner = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(std, &inner[i])
		}
		fmt.Fscanln(std)
		outer = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(std, &outer[i])
		}
		fmt.Fscanln(std)

		if n == 1 {
			if inner[0]+outer[0] <= w {
				fmt.Fprintln(std, 1)
			} else {
				fmt.Fprintln(std, 2)
			}
			continue
		}

		solve()

		fmt.Fprintln(std, min)
	}
}

var flag bool

func solve() {
	// a: outer block
	// b: inner block
	// c: all block
	memoa := make([]int, n)
	memob := make([]int, n)
	memoc := make([]int, n)

	// case clean
	memoa[0] = 1
	memob[0] = 1
	if inner[0]+outer[0] <= w {
		memoc[0] = 1
	} else {
		memoc[0] = 2
	}
	flag = true

	calc(memoa, memob, memoc)
	min = getmin([]int{min, memoc[n-1]})
	// fmt.Println(min)

	flag = false

	// case outer merge
	if outer[0]+outer[n-1] <= w {
		memoa[0] = 1
		memob[0] = -1
		memoc[0] = 2

		calc(memoa, memob, memoc)
		min = getmin([]int{min, memob[n-1]})
	}
	// fmt.Println(min)

	// case inner merge
	if inner[0]+inner[n-1] <= w {
		memoa[0] = -1
		memob[0] = 1
		memoc[0] = 2

		calc(memoa, memob, memoc)
		min = getmin([]int{min, memoa[n-1]})
	}
	// fmt.Println(min)

	// case both merge
	if outer[0]+outer[n-1] <= w && inner[0]+inner[n-1] <= w {
		memoa[0] = -1
		memob[0] = -1
		memoc[0] = 2

		calc(memoa, memob, memoc)
		min = getmin([]int{min, memoc[n-2]})
	}
	// fmt.Println(min)
}

func calc(memoa, memob, memoc []int) {
	for i := 1; i < n; i++ {
		var temp []int

		temp = []int{}
		if memob[i-1] > 0 && outer[i]+outer[i-1] <= w {
			temp = append(temp, memob[i-1]+1)
		}
		temp = append(temp, memoc[i-1]+1)
		memoa[i] = getmin(temp)

		temp = []int{}
		if memoa[i-1] > 0 && inner[i]+inner[i-1] <= w {
			temp = append(temp, memoa[i-1]+1)
		}
		temp = append(temp, memoc[i-1]+1)
		memob[i] = getmin(temp)

		temp = []int{}
		if inner[i]+outer[i] <= w {
			temp = append(temp, memoc[i-1]+1)
		} else {
			temp = append(temp, memoc[i-1]+2)
		}
		if inner[i-1]+inner[i] <= w && outer[i-1]+outer[i] <= w {
			if i == 1 && flag {
				temp = append(temp, 2)
			} else if i > 1 {
				temp = append(temp, memoc[i-2]+2)
			}
		}
		temp = append(temp, memoa[i]+1)
		temp = append(temp, memob[i]+1)
		memoc[i] = getmin(temp)
	}
}

func getmin(l []int) int {
	res := math.MaxInt
	for _, v := range l {
		if res > v {
			res = v
		}
	}
	return res
}
