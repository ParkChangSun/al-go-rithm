package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var std = bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))

type Point struct {
	x, y       int
	a, b, c, d *Point
}

func main() {
	defer std.Flush()

	var n, k int
	var jumps string
	fmt.Fscan(std, &n, &k, &jumps)

	dirad := map[int][]*Point{}
	dirbc := map[int][]*Point{}

	var cur *Point = nil

	for i := 0; i < n; i++ {
		p := &Point{}
		fmt.Fscan(std, &p.x, &p.y)

		dirad[p.x-p.y] = append(dirad[p.x-p.y], p)
		dirbc[p.x+p.y] = append(dirbc[p.x+p.y], p)

		if i == 0 {
			cur = p
		}
	}

	for _, v := range dirad {
		slices.SortFunc(v, func(a *Point, b *Point) int { return a.x - b.x })
		for i := 0; i < len(v)-1; i++ {
			v[i].a = v[i+1]
			v[i+1].d = v[i]
		}
	}
	for _, v := range dirbc {
		slices.SortFunc(v, func(a *Point, b *Point) int { return a.x - b.x })
		for i := 0; i < len(v)-1; i++ {
			v[i].b = v[i+1]
			v[i+1].c = v[i]
		}
	}

	for _, v := range jumps {
		if v == 'A' && cur.a != nil {
			jump(cur)
			cur = cur.a
		}
		if v == 'B' && cur.b != nil {
			jump(cur)
			cur = cur.b
		}
		if v == 'C' && cur.c != nil {
			jump(cur)
			cur = cur.c
		}
		if v == 'D' && cur.d != nil {
			jump(cur)
			cur = cur.d
		}
	}

	fmt.Fprint(std, cur.x, cur.y)
}

func jump(p *Point) {
	if p.a != nil {
		p.a.d = p.d
	}
	if p.d != nil {
		p.d.a = p.a
	}
	if p.b != nil {
		p.b.c = p.c
	}
	if p.c != nil {
		p.c.b = p.b
	}
}
