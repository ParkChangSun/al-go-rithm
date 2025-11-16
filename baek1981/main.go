package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var std = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

var n int
var matrix [][]int

type Point struct {
	i, j int
}

func main() {
	defer std.Flush()

	fmt.Fscan(std, &n)
	matrix = make([][]int, n)
	numberSet := []int{}
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			fmt.Fscan(std, &matrix[i][j])
			numberSet = append(numberSet, matrix[i][j])
		}
	}

	slices.Sort(numberSet)
	numberSet = slices.Compact(numberSet)

	res := 200

	lowIndex, highIndex := 0, 0

	for lowIndex <= highIndex && highIndex < len(numberSet) {
		if search(numberSet[lowIndex], numberSet[highIndex]) {
			res = min(res, numberSet[highIndex]-numberSet[lowIndex])
			lowIndex++
		} else {
			highIndex++
		}
	}

	fmt.Fprint(std, res)
}

var neighborSet = []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func search(minBound, maxBound int) bool {
	if x := matrix[0][0]; x < minBound || maxBound < x {
		return false
	}

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	queue := []Point{{0, 0}}
	visited[0][0] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.i == n-1 && cur.j == n-1 {
			return true
		}

		for _, v := range neighborSet {
			ni, nj := cur.i+v.i, cur.j+v.j

			if ni < 0 || ni >= n || nj < 0 || nj >= n || visited[ni][nj] || matrix[ni][nj] < minBound || maxBound < matrix[ni][nj] {
				continue
			}

			queue = append(queue, Point{ni, nj})
			visited[ni][nj] = true
		}
	}

	return false
}

// 5
// 9 0 9 9 9
// 9 0 9 0 9
// 9 0 9 0 9
// 9 0 9 0 9
// 9 9 9 0 9

// 5
// 8 0 8 8 8
// 0 0 8 0 8
// 8 0 8 0 8
// 8 0 8 0 8
// 8 8 8 0 8

// 5
// 1 1 3 6 8
// 1 2 2 5 5
// 4 4 0 3 3
// 8 0 2 3 4
// 4 3 0 2 1

// 5
// 3 4 11 8 7
// 0 12 17 11 10
// 13 19 12 15 6
// 10 8 15 17 18
// 6 0 15 1 8
// answer= 14

// 5
// 5 5 5 0 0
// 5 0 5 0 0
// 5 0 4 0 0
// 5 0 5 5 8
// 5 5 7 0 5

// 5
// 4 2 3 4 5
// 5 4 3 2 1
// 1 2 4 3 5
// 5 4 3 2 3
// 1 2 3 4 2
// 2

// 사고 과정 기록해야 함
// 정렬 후 유니크
