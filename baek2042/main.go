package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var n, m, k int
var arr, tree []int
var logLeaf, treeLen, upperTotal int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n, &m, &k)
	arr = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &arr[i])
	}

	logLeaf = 1 << int(math.Log2(float64(n)))
	upperTotal = 1<<int(math.Log2(float64(n))+1) - 1
	treeLen = (n-logLeaf)*2 + upperTotal

	tree = make([]int, 1<<(int(math.Ceil(math.Log2(float64(n))))+1))
	generateTree(1, n, 1)
	// fmt.Println(tree)

	for i := 0; i < m+k; i++ {
		var a, b, c int
		fmt.Fscanln(reader, &a, &b, &c)
		if a == 1 {
			replace(b, 1, n, 1, c)
			// fmt.Println(tree)
		} else {
			fmt.Fprintln(writer, calc(b, c, 1, n, 1))
		}
	}
}

func generateTree(h, t, index int) int {
	if h == t {
		tree[index] = arr[h]
		return tree[index]
	}

	b := (h + t) / 2
	tree[index] = generateTree(h, b, index*2) + generateTree(b+1, t, index*2+1)
	return tree[index]
}

func replace(index, head, tail, treeIndex, num int) {
	if head == tail {
		tree[treeIndex] = num
		arr[index] = num
		return
	}

	diff := num - arr[index]
	tree[treeIndex] += diff
	b := (head + tail) / 2

	if index <= b {
		replace(index, head, b, treeIndex*2, num)
	} else {
		replace(index, b+1, tail, treeIndex*2+1, num)
	}
}

// 네 값 모두 arr위의 값이다
// 초기값은12
// from to 는 더하고자 하는 범위
// head tail 은 1 len(arr)
func calc(from, to, head, tail, index int) int {
	if from == head && to == tail {
		return tree[index]
	}

	// 탐색 중인 범위 안에 더하고자 하는 숫자가 없는 경우
	if to < head || tail < from {
		return 0
	}

	ap, dui := 0, 0
	b := (head + tail) / 2

	if to < b {
		// 뒤를 계산할 필요가 없는 경우
		// 더하고자 하는 범위의 끝이 중간값보다 작다
		ap = calc(from, to, head, b, index*2)
		dui = 0
	} else if b < from {

		// 앞을 계산할 필요가 없는 경우
		// 더하고자 하는 범위의 시작이 중간값보다 크다
		ap = 0
		dui = calc(from, to, b+1, tail, index*2+1)
	} else {
		ap = calc(from, b, head, b, index*2)
		dui = calc(b+1, to, b+1, tail, index*2+1)
	}

	return ap + dui
}

// logleaf + n = len tree

// if logleaf*2 + index - 1 <= len tree then 1,2,3..
// else ...n

// logleaf 4
// [0 15 6 9 3 3 4 5 1 2 0 0 0 0 0 0]
//     1 2 3 4 5 6 7 8 9    treeindex
//             3 4 5 1 2    arrindex

// [0 6 3 3 1 2 0 0]
//        3 4 5
//        3 1 2

// logleaf 2
// uppertree 3
// treelengh 5

// replace 이분정복 활용 못한것
