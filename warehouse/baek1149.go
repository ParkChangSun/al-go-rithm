package baeksolution

import (
	"bufio"
	"fmt"
	"os"
)

type costVector struct {
	r, g, b int
}

type costObj struct {
	minCost  costVector
	costList []costVector
}

//Solution1149 No.1149
func Solution1149() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	costObj := costObj{costVector{0, 0, 0}, make([]costVector, n+1)}
	var r, g, b int
	for i := 1; i <= n; i++ {
		fmt.Fscanln(reader, &r, &g, &b)
		costObj.costList[i] = costVector{r, g, b}
	}

	ans := costObj.calcCost(1)
	fmt.Fprint(writer, min(ans.r, min(ans.g, ans.b)))
}

//i번째 집을 R로 칠할 때 총 비용은 i-1번째에서 B또는G로 칠한 두 경우 중 비용이 적은 쪽을 선택
//i번째 집을 b,g로 칠할 때도 다 구해서 세 가지를 비교한게 답
//1번째 집에서부터 최소값만 찾아서 더한다고 해서 답이 최소가 나오는 걸 보장할 수 없기 때문
//d=i번째를 rgb중 하나로 칠했을 때 1~i까지 칠하는 최소비용
//a=i번째 집을 rgb중 하나로 칠하는 비용
//d[i][r] = min(d[i-1][g], d[i-1][b]) + a[i][r]
func (costObj *costObj) calcCost(i int) costVector {
	costObj.minCost.sumCost(costObj.costList[i])
	if i == len(costObj.costList)-1 {
		return costObj.minCost
	}
	return costObj.calcCost(i + 1)
}

func (minCost *costVector) sumCost(i costVector) {
	var r, g, b int
	// if minCost.g >= minCost.b {
	// 	r = i.r + minCost.b
	// } else {
	// 	r = i.r + minCost.g
	// }
	r = i.r + min(minCost.g, minCost.b)
	// if minCost.r >= minCost.b {
	// 	g = i.g + minCost.b
	// } else {
	// 	g = i.g + minCost.r
	// }
	g = i.g + min(minCost.r, minCost.b)
	// if minCost.r >= minCost.g {
	// 	b = i.b + minCost.g
	// } else {
	// 	b = i.b + minCost.r
	// }
	b = i.b + min(minCost.r, minCost.g)
	minCost.r = r
	minCost.g = g
	minCost.b = b
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
