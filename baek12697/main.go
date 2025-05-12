package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	R, C int
}

type CellData struct {
	D     int
	Prev  Point
	Walls []Point
}

const (
	EMPTY = '.'
	WALL  = '#'
	START = 'O'
	CAKE  = 'X'
)

var NONEPOINT = Point{-1, -1}

var r, c int
var gameMap []string
var start, cake Point

func main() {
	std := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer std.Flush()

	var n int
	fmt.Fscanln(std, &n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(std, &r, &c)
		gameMap = make([]string, r)
		for j := 0; j < r; j++ {
			fmt.Fscanln(std, &gameMap[j])
			if k := strings.Index(gameMap[j], "O"); k != -1 {
				start = Point{j, k}
			}
			if k := strings.Index(gameMap[j], "X"); k != -1 {
				cake = Point{j, k}
			}
		}

		if ans := solve(); ans != 0 {
			fmt.Fprintf(std, "Case #%d: %d\n", i+1, ans)
		} else {
			fmt.Fprintf(std, "Case #%d: THE CAKE IS A LIE\n", i+1)
		}
	}
}

func solve() int {
	queue := []Point{start}
	status := map[Point]CellData{start: {0, NONEPOINT, []Point{}}}

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		sight := []Point{NONEPOINT, NONEPOINT, NONEPOINT, NONEPOINT}
		for i := 1; i <= 15; i++ {
			if sight[0] == NONEPOINT && isWall(Point{cur.R + i, cur.C}) {
				sight[0] = Point{cur.R + i - 1, cur.C}
			}
			if sight[1] == NONEPOINT && isWall(Point{cur.R - i, cur.C}) {
				sight[1] = Point{cur.R - i + 1, cur.C}
			}
			if sight[2] == NONEPOINT && isWall(Point{cur.R, cur.C + i}) {
				sight[2] = Point{cur.R, cur.C + i - 1}
			}
			if sight[3] == NONEPOINT && isWall(Point{cur.R, cur.C - i}) {
				sight[3] = Point{cur.R, cur.C - i + 1}
			}
		}
		t := status[cur]
		t.Walls = append(t.Walls, sight...)
		status[cur] = t

		_, neighs := neighbors(cur)
		for _, n := range neighs {
			if next, visited := status[n]; !visited || next.D > status[cur].D+1 {
				queue = append(queue, n)
				status[n] = CellData{status[cur].D + 1, cur, status[cur].Walls}
			}
		}

		for portalCursor, count := cur, 0; portalCursor != NONEPOINT; portalCursor, count = status[portalCursor].Prev, count+1 {
			if w, _ := neighbors(portalCursor); w {
				for _, portal := range status[cur].Walls {
					if next, visited := status[portal]; !visited || next.D > status[cur].D+1+count {
						queue = append(queue, portal)
						status[portal] = CellData{status[cur].D + 1 + count, cur, []Point{}}
					}
				}
				break
			}
		}
	}

	return status[cake].D
}

func isWall(p Point) bool {
	return p.R >= r || p.C >= c || p.R <= -1 || p.C <= -1 || gameMap[p.R][p.C] == WALL
}

func neighbors(cur Point) (bool, []Point) {
	wall := false
	dir := []Point{{cur.R - 1, cur.C}, {cur.R, cur.C - 1}, {cur.R + 1, cur.C}, {cur.R, cur.C + 1}}
	neigh := []Point{}
	for _, v := range dir {
		if isWall(v) {
			wall = true
		} else {
			neigh = append(neigh, v)
		}
	}
	return wall, neigh
}
