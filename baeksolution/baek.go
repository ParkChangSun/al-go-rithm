package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

type country struct {
	i, g, s, b int
}

//Solution0 No.8979
func Solution() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	countries := []country{}
	var target country
	var in, g, s, b int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &in, &g, &s, &b)
		if in == k {
			target = country{in, g, s, b}
		}
		countries = append(countries, country{in, g, s, b})
	}

	rank := 1
	for _, v := range countries {
		if v.g > target.g {
			rank++
		}
	}
	for _, v := range countries {
		if v.g == target.g {
			if v.s > target.s {
				rank++
			}
		}
	}
	for _, v := range countries {
		if v.g == target.g && v.s == target.s {
			if v.b > target.b {
				rank++
			}
		}
	}
	fmt.Fprint(writer, rank)
}
