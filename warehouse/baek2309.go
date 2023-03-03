package baeksolution

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

//Solution No.2309
func Solution2309() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	lil := make([]int, 9)
	sum := 0
	for i := 0; i < 9; i++ {
		fmt.Fscanln(reader, &lil[i])
		sum += lil[i]
	}
RES:
	for i := 0; i < 9; i++ {
		for j := i + 1; j < 9; j++ {
			if sum-100 == lil[i]+lil[j] {
				lil = append(lil[:i], lil[i+1:]...)
				lil = append(lil[:j-1], lil[j:]...)
				break RES
			}
		}
	}
	sort.SliceStable(lil, func(i, j int) bool {
		return lil[i] < lil[j]
	})
	for _, v := range lil {
		fmt.Fprintln(writer, v)
	}

}
