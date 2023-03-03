package baeksolution

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

//Solution1181 .
func Solution1181() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscanln(reader, &n)

	m := make(map[string]struct{})
	for i := 0; i < n; i++ {
		var str string
		fmt.Fscanln(reader, &str)
		m[str] = struct{}{}
	}

	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if len(keys[i]) == len(keys[j]) {
			return keys[i] < keys[j]
		}
		return len(keys[i]) < len(keys[j])
	})

	for _, v := range keys {
		fmt.Fprintln(writer, v)
	}
}
