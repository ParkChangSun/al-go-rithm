package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

type store struct {
	inc, dec int
}

//Solution11054 No.11054
func Solution11054() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	seq := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(reader, "%d", &seq[i])
	}

	lengthStore := make([]store, n+1)
	for i := 1; i <= n; i++ {
		lengthStore[i].inc = 1
		lengthStore[i].dec = 1
	}
	longBitonic(1, seq, lengthStore)
	fmt.Fprint(writer, maxStore(lengthStore))
	fmt.Fprint(writer, lengthStore)
}

// 앞뒤로 LIS 구해서 더하는 방식이 정석임

// s[n] = seq[n]을 포함하는 가장 긴 수열 길이
// s[n] = increasing seq[n] > seq[i]
// 				최대값 없는 경우 seq[n]보다 작은 수가 마지막인 수열 + 1
// 				decreasing seq[n] < seq[i]
// 				최대값이 정해지는 경우 seq[n]보다 큰 수가 마지막인 수열 + 1
// 				최대값 있는 경우 seq[n]보다 큰 수가 마지막인 수열 + 1
// 				중 가장 긴 것

func longBitonic(n int, seq []int, store []store) {
	if n == len(seq) {
		return
	}
	for i := 1; i < n; i++ {
		if seq[n] > seq[i] {
			if store[n].inc <= store[i].inc {
				store[n].inc = store[i].inc + 1
			}
		} else if seq[n] < seq[i] {
			if store[n].dec <= store[i].inc {
				store[n].dec = store[i].inc + 1
			} else if store[n].dec <= store[i].dec {
				store[n].dec = store[i].dec + 1
			}
		}
	}
	longBitonic(n+1, seq, store)
}

func maxStore(arr []store) (max int) {
	for _, v := range arr {
		if v.inc > max {
			max = v.inc
		}
		if v.dec > max {
			max = v.dec
		}
	}
	return
}
