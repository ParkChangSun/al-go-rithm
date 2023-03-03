package baeksolution

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

//Dot dot
type Dot struct {
	x, y int
}

//Solution11650 .
func Solution11650() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscanln(reader, &n)
	dots := make([]Dot, 0)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)
		dots = append(dots, Dot{a, b})
	}
	// sort.SliceStable(dots, func(i, j int) bool {
	// 	return dots[i].x < dots[j].x
	// })
	sort.SliceStable(dots, func(i, j int) bool {
		if dots[i].x == dots[j].x {
			return dots[i].y < dots[j].y
		}
		return dots[i].x < dots[j].x
	})
	for _, v := range dots {
		fmt.Fprintln(writer, v.x, v.y)
	}
}

// 4
// 2 2
// 1 1
// 2 -2
// 1 -1
