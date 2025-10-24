package main

import (
	"bufio"
	"fmt"
	"os"
)

var std = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

func main() {
	defer std.Flush()

	var n int
	fmt.Fscan(std, &n)
	type App struct {
		s, i, p int
	}
	apps := make([]App, n)
	for i := range n {
		fmt.Fscan(std, &apps[i].s, &apps[i].i, &apps[i].p)
	}

	dp := make([][]bool, 1001)
	for i := range dp {
		dp[i] = make([]bool, 1001)
	}

	maxClear := 0

	queue := []App{{1, 1, 0}}
	dp[1][1] = true

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		all := 2
		clear := 0

		for _, v := range apps {
			if state.i >= v.i || state.s >= v.s {
				all += v.p
				clear++
			}
		}

		if maxClear < clear {
			maxClear = clear
		}

		turnGains := all - state.s - state.i
		if all >= 2000 {
			maxClear = n
			break
		} else if state.s+turnGains > 1000 {
			state.i += state.s + turnGains - 1000
			for i := range 1001 {
				if state.i+i > 1000 && 1000-i < state.s {
					break
				}
				if !dp[1000-i][state.i+i] {
					queue = append(queue, App{1000 - i, state.i + i, 0})
					dp[1000-i][state.i+i] = true
				}
			}
		} else if state.i+turnGains > 1000 {
			state.s += state.i + turnGains - 1000
			for i := range 1001 {
				if state.s+i > 1000 && 1000-i < state.i {
					break
				}
				if !dp[state.s+i][1000-i] {
					queue = append(queue, App{state.s + i, 1000 - i, 0})
					dp[state.s+i][1000-i] = true
				}
			}
		} else {
			for i := range turnGains + 1 {
				if !dp[state.s+i][state.i+(turnGains-i)] {
					queue = append(queue, App{state.s + i, state.i + (turnGains - i), 0})
					dp[state.s+i][state.i+(turnGains-i)] = true
				}
			}
		}
	}

	fmt.Fprint(std, maxClear)
}
