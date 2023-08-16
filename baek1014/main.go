package main

import (
	"bufio"
	"fmt"
	"os"
)

const X byte = 'x'

var k, l int
var room []string
var memo [][]int

func main() {
	std := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer std.Flush()

	var n int
	fmt.Fscanln(std, &n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(std, &k, &l)
		room = make([]string, k)
		for i := 0; i < k; i++ {
			fmt.Fscanln(std, &room[i])
		}
		memo = make([][]int, k)
		for i := 0; i < k; i++ {
			memo[i] = make([]int, 1<<l)
			for j := 0; j < 1<<l; j++ {
				memo[i][j] = -1
			}
		}

		fmt.Fprintln(std, solve(0, 0, 0, 0))
	}
}

func solve(uppermask int, line int, seatMask int, seatNum int) int {
	// fmt.Printf("%10b %d %10b %d\n", uppermask, line, seatMask, seatNum)
	if line == k {
		return 0
	}

	if seatNum >= l {
		// 이번 줄에 이렇게 앉힌 상태는 오직 아래에만 영향을 미친다
		// 자리 하나하나마다 저장하면 윗줄의 상태를 반영할 수 없다
		if memo[line][seatMask] == -1 {
			memo[line][seatMask] = solve(seatMask, line+1, 0, 0)
		}
		return memo[line][seatMask]
	}

	unavailable := room[line][seatNum] == X
	leftCheat := seatNum > 0 && ((uppermask>>(seatNum-1))&1) == 1
	rightCheat := seatNum < l-1 && ((uppermask>>(seatNum+1))&1) == 1
	if unavailable || leftCheat || rightCheat {
		return solve(uppermask, line, seatMask, seatNum+1)
	}

	noCase := solve(uppermask, line, seatMask, seatNum+1)
	sitCase := solve(uppermask, line, seatMask|(1<<seatNum), seatNum+2) + 1
	return max(noCase, sitCase)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 1
// 10 10
// ....x.....
// ..........
// ..........
// ..x.......
// ..........
// x...x.x...
// .........x
// ...x......
// ........x.
// .x...x....

// 가능한 변수는 - 몇번째 줄인지, 그 줄에 어떻게 학생들을 앉혔는지(비트마스크?)

// 첫번째 줄부터 시작
// 그 줄에 앉히는 것도 변수 맞춰서 해야하나?
// 이 줄의 이 자리에 앉혔을때가 더 큰가? 아니면 안 앉혔을때(아마 기존의 값)이 더 큰가?
// max(이 자리에 앉음, 기존의 값)

// 첫째 줄을 계산할 때 확실한 값이 나와야 하는가?
// 그것도 알 수 없다. 첫째 줄에서 최대값이라고 해서 정답인 것은 아니다.

// 100개를 마스킹할 수는 없다.

// 무슨 값을 저장해야 하는가?
// 앉을 수 있는 최대값인가?
// 그 줄의 최대값인가?

// 함수를 실행할 때 현재의 모든 비트마스크 10줄을 다 줘야 하나?
// 필요한 건 바로 윗줄이긴 하다.

// 아니 이번 라인에서 이렇게 앉혔다고 최대값을 적었는데, 다른 방법의 윗줄들이 더 많이 앉혔을수도 있는거 아님?
// 어떤 값을 저장하고, 어떤 값을 함수에서 들고 있는지를 봐야 하는데...

// 10101001000
// 10001001001

// 문제 분해를 n번째 줄에서부터 마지막 줄까지의 정답에 속하는, 앉을 수 있는 인원의 최대값이라고 할 수 있는가?

// 메모에 저장하는 거는 어떤 줄에 어떻게 앉혔을때니까 그 아랫줄부터만 영향을 미친다.
// 메모는 이게 맞는거 같기도 한데?

// 앉힌거랑 안앉힌거랑 순서에 차이가 있다?

// 17
// 2 3
// ...
// ...
// 2 3
// x.x
// xxx
// 2 3
// x.x
// x.x
// 10 10
// ....x.....
// ..........
// ..........
// ..x.......
// ..........
// x...x.x...
// .........x
// ...x......
// ........x.
// .x...x....
// 10 10
// xxxxxxxxxx
// xxxxxxxxxx
// xxxxxxxxxx
// xxxxxxxxxx
// xxxxxxxxxx
// xxxxxxxxxx
// xxxxxxxxxx
// xxxxxxxxxx
// xxxxxxxxxx
// xxxxxxxxxx
// 3 3
// xx.
// xxx
// x.x
// 3 5
// ..x..
// ..x..
// .xxx.
// 1 6
// ..x...
// 10 10
// .x.x...x..
// .x..x.....
// x.x.......
// .x.x......
// x...x.....
// .x.x...x..
// .x..x.....
// x.x.......
// .x.x......
// x...x.....
// 5 10
// .x.x...x..
// .x..x.....
// x.x.......
// .x.x......
// x...x.....
// 5 8
// .x...x..
// ..x.....
// x.......
// .x......
// ..x.....
// 5 7
// x...x..
// .x.....
// .......
// x......
// .x.....
// 4 3
// x.x
// x..
// .x.
// .xx
// 3 4
// .x.x
// .xx.
// x.x.
// 5 4
// x.x.
// .x.x
// x.x.
// x..x
// ..xx
// 10 10
// .x.x...x..
// .x..x.....
// x.xx..x...
// .x.x....x.
// x...x.x...
// .x.x...x..
// .x..x...x.
// x.x.......
// .x.x......
// x.........
// 10 10
// xxxxxxxxxx
// ..........
// ..........
// xxxxxxxxxx
// ..........
// xxxxxxxxxx
// .........x
// ...x......
// ........x.
// xxxxxxxxxx
