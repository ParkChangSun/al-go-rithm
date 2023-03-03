package baeksolution

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Solution0 No.0
func Solution1652() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	room := [][]string{}
	for i := 0; i < n; i++ {
		var tp string
		fmt.Fscanln(reader, &tp)
		room = append(room, strings.Split(tp, ""))
	}

	// i row
	// j col
	var rowCount, colCount, rowRes, colRes int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if room[i][j] == "." {
				rowCount++
			} else if room[i][j] == "X" {
				if rowCount > 1 {
					rowRes++
				}
				rowCount = 0
			}

			if room[j][i] == "." {
				colCount++
			} else if room[j][i] == "X" {
				if colCount > 1 {
					colRes++
				}
				colCount = 0
			}
		}
		if rowCount > 1 {
			rowRes++
		}
		if colCount > 1 {
			colRes++
		}
		rowCount = 0
		colCount = 0
	}
	fmt.Print(rowRes, colRes)
}
