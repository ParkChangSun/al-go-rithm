package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var n int
var w, store, checked []int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n)

	w = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &w[i])
	}

	checked = make([]int, n+1)
	store = make([]int, n+1)
	store[1] = w[1]
	if n > 1 {
		store[2] = w[2] + w[1]
	}

	fmt.Fprintln(writer, drink(n))
}

func drink(n int) int {
	if n <= 2 || checked[n] == 1 {
		return store[n]
	}

	notdrinkn := drink(n - 1)
	drinkone := w[n] + drink(n-2)
	drinkboth := w[n] + w[n-1] + drink(n-3)

	l := []int{notdrinkn, drinkone, drinkboth}
	sort.Ints(l)
	store[n] = l[2]
	checked[n] = 1
	return l[2]
}

// 최대 연속 두개까지만 먹는다
// 이것을 안먹거나 010 110 100 drink(n-1)
// 이것만 먹거나 001 101 [n]+drink(n-2)
// 이거랑 이 전것을 먹거나 011 [n]+[n-1]+drink(n-3)
