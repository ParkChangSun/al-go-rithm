package main

import (
	"bufio"
	"fmt"
	"os"
)

var matrix = map[int][]int{}

var std *bufio.ReadWriter = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

func main() {
	defer std.Flush()

	var n, r, q int
	fmt.Fscanln(std, &n, &r, &q)

	// solution
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscanln(std, &u, &v)
		matrix[u] = append(matrix[u], v)
		matrix[v] = append(matrix[v], u)
	}

	ans := make([]int, n+1)

	visited := make([]bool, n+1)
	parent := make([]int, n+1)

	stack := []int{r}
	stacks := []int{r}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		visited[node] = true

		for _, v := range matrix[node] {
			if visited[v] {
				continue
			}
			stack = append(stack, v)
			stacks = append(stacks, v)
			parent[v] = node

		}
	}

	for len(stacks) != 0 {
		node := stacks[len(stacks)-1]
		stacks = stacks[:len(stacks)-1]

		ans[node]++
		ans[parent[node]] += ans[node]
	}

	for i := 0; i < q; i++ {
		var target int
		fmt.Fscanln(std, &target)
		fmt.Fprintln(std, ans[target])
	}
}
