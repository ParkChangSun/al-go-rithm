package baeksolution

/*
1
Implement the Quicksort
Count and show the number of comparison and exchange operations
2
compare the following three ways of choice of pivot.
2.1 The first element is the pivot as shown in the textbook.
2.2 A randomly chosen element between low and high is the pivot.
2.3 The pivot is chosen as the median number among the first, the last, and the mid elements between low and high.
(For example, if the first, the mid, and the last element are 10, 7, 20, respectively, the first element 10 is the median value. In Korean, 중간값.)
4. No need to show me the source code listing. I am interested in the final result of your experiments.
*/

// var arr = make([]int, 0)
// var count = 0

//Solution0 No.0
// func quick() {
// 	writer := bufio.NewWriter(os.Stdout)
// 	defer writer.Flush()

// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < 100; i++ {
// 		arr = append(arr, rand.Intn(101))
// 	}

// 	// qs(0, len(arr)-1)
// 	// qs2(0, len(arr)-1)
// 	qs3(0, len(arr)-1)
// 	fmt.Fprint(writer, arr)
// 	fmt.Fprintln(writer, count)
// }

// func qs(low, high int) {
// 	if low < high {
// 		pivot := partition1(low, high)
// 		qs(low, pivot-1)
// 		qs(pivot+1, high)
// 		count += 2
// 	}
// 	count++
// }

// func qs2(low, high int) {
// 	if low < high {
// 		pivot := partition2(low, high)
// 		qs2(low, pivot-1)
// 		qs2(pivot+1, high)
// 	}
// }

// func qs3(low, high int) {
// 	if low < high {
// 		pivot := partition3(low, high)
// 		qs3(low, pivot-1)
// 		qs3(pivot+1, high)
// 	}
// }

// func partition1(low, high int) int {
// 	pivot := arr[low]
// 	i := low
// 	count += 2
// 	for j := low + 1; j <= high; j++ {
// 		if arr[j] < pivot {
// 			i++
// 			arr[i], arr[j] = arr[j], arr[i]
// 			count += 3
// 		}
// 		count++
// 	}
// 	count++
// 	arr[i], arr[low] = arr[low], arr[i]
// 	count += 2
// 	return i
// }

// func partition2(low, high int) int {
// 	a := high - low - 1
// 	if a > 0 {
// 		pivotPoint := low + rand.Intn(a)
// 		arr[pivotPoint], arr[low] = arr[low], arr[pivotPoint]
// 		count += 5
// 	}
// 	count += 2

// 	return partition1(low, high)
// }

// func partition3(low, high int) int {
// 	mid := (low + high) / 2
// 	l := []int{arr[low], arr[mid], arr[high]}
// 	sort.Ints(l)
// 	if l[1] == arr[low] {
// 		count++
// 	} else if l[1] == arr[mid] {
// 		arr[low], arr[mid] = arr[mid], arr[low]
// 		count += 3
// 	} else {
// 		arr[low], arr[high] = arr[high], arr[low]
// 		count += 3
// 	}

// 	count += 3
// 	return partition1(low, high)
// }
