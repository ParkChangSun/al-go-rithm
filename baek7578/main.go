package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var n int
var tree []int

func main() {
	std := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer std.Flush()

	fmt.Fscanln(std, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(std, &a[i])
	}
	index := make([]int, 1000001)
	for i := 1; i <= n; i++ {
		var b int
		fmt.Fscan(std, &b)
		index[b] = i
	}

	treeSize := 1 << int(math.Ceil(math.Log2(float64(n)))+1)
	tree = make([]int, treeSize+1)
	ans := 0
	for _, v := range a {
		ans += query(1, 1, n, index[v], treeSize)
		update(1, 1, n, 1, index[v])
	}
	fmt.Fprintln(std, ans)
}

func update(treeIndex, left, right, value, arrIndex int) {
	if left == right {
		tree[treeIndex] = value
		return
	}

	mid := (left + right) / 2
	if arrIndex <= mid {
		update(treeIndex*2, left, mid, value, arrIndex)
	} else {
		update(treeIndex*2+1, mid+1, right, value, arrIndex)
	}
	tree[treeIndex] = tree[treeIndex*2] + tree[treeIndex*2+1]
}

func query(index, left, right, from, to int) int {
	if right < from || to < left {
		return 0
	}

	if from <= left && right <= to {
		return tree[index]
	}

	mid := (left + right) / 2
	return query(index*2, left, mid, from, to) + query(index*2+1, mid+1, right, from, to)
}
