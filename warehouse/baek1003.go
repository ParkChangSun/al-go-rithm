package baeksolution

import (
	"bufio"
	"os"
)

// type countVertex struct {
// 	zeroCount, oneCount int
// }

// type fiboArray struct {
// 	arr []countVertex
// }

//Solution1003 .
func Solution1003() {
	// reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// var n, c int
	// fmt.Fscanln(reader, &n)

	// null := countVertex{0, 0}
	// for i := 0; i < n; i++ {
	// 	fmt.Fscanln(reader, &c)
	// 	f := fiboArray{make([]countVertex, c+1)}
	// 	zAns, oAns := fibo(c, &f, null)
	// 	fmt.Fprintln(writer, zAns, oAns)
	// }
}

// func fibo(n int, f *fiboArray, null countVertex) (int, int) {
// 	if n == 0 {
// 		return 1, 0
// 	} else if n == 1 {
// 		return 0, 1
// 	} else {
// 		//if doesnt exist
// 		if f.arr[n] != null {
// 			return f.arr[n].zeroCount, f.arr[n].oneCount
// 		}
// 		z1, o1 := fibo(n-1, f, null)
// 		z2, o2 := fibo(n-2, f, null)
// 		z, o := z1+z2, o1+o2
// 		f.arr[n] = countVertex{z, o}
// 		return z, o
// 	}
// }
