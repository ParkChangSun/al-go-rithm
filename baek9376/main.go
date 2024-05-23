package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	WALL = '*'
	DOOR = '#'
	ROOM = '.'
	MAN  = '$'
)

func main() {
	std := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	defer std.Flush()

	var n int
	fmt.Fscanln(std, &n)

	for i := 0; i < n; i++ {
		p := Prison{}
		fmt.Fscanln(std, &p.H, &p.W)
		p.World = make([]string, p.H)
		for i := 0; i < p.H; i++ {
			fmt.Fscanln(std, &p.World[i])
		}
		var s strings.Builder
		for i := 0; i < p.W; i++ {
			fmt.Fprintf(&s, "%c", ROOM)
		}
		p.World = append([]string{s.String()}, append(p.World, s.String())...)
		for i, v := range p.World {
			p.World[i] = fmt.Sprintf("%c%s%c", ROOM, v, ROOM)
		}
		p.H += 2
		p.W += 2

		fmt.Fprintln(std, solve(p))
	}
}

type Point struct {
	X, Y int
}

func (p Point) setValue(arr *[][]int, v int) {
	(*arr)[p.X][p.Y] = v
}

func (p Point) getValue(arr [][]int) int {
	return arr[p.X][p.Y]
}

type Prison struct {
	World []string
	H, W  int
}

func (prison Prison) location(p Point) byte {
	return prison.World[p.X][p.Y]
}

func solve(p Prison) int {
	mans := []Point{{0, 0}}
	for i, v := range p.World {
		for j, l := range v {
			if l == MAN {
				mans = append(mans, Point{i, j})
			}
		}
	}

	res := make([][]int, p.H)
	for i := range res {
		res[i] = make([]int, p.W)
	}

	for _, v := range mans {
		for i, r := range bfs(v, p) {
			for j, c := range r {
				res[i][j] += c
			}
		}
	}

	hold := math.MaxInt
	for i, r := range res {
		for j := range r {
			loc := Point{i, j}
			if p.location(loc) == DOOR {
				loc.setValue(&res, loc.getValue(res)-2)
			}
			if c := loc.getValue(res); c >= 0 && hold > c {
				hold = c
			}
		}
	}

	return hold
}

func bfs(p Point, prison Prison) [][]int {
	dist := make([][]int, prison.H)
	for i := range dist {
		dist[i] = make([]int, prison.W)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	p.setValue(&dist, 0)
	deque := []Point{p}

	for len(deque) > 0 {
		cur := deque[0]
		deque = deque[1:]

		for _, v := range getNearPath(cur, prison) {
			if dist[v.X][v.Y] != -1 {
				continue
			}

			if l := prison.location(v); l == DOOR {
				v.setValue(&dist, cur.getValue(dist)+1)
				deque = append(deque, v)
			} else if l == ROOM || l == MAN {
				v.setValue(&dist, cur.getValue(dist))
				deque = append([]Point{v}, deque...)
			}
		}
	}

	return dist
}

func getNearPath(p Point, prison Prison) []Point {
	res := []Point{}
	if p.X-1 >= 0 {
		res = append(res, Point{p.X - 1, p.Y})
	}
	if p.X+1 < prison.H {
		res = append(res, Point{p.X + 1, p.Y})
	}
	if p.Y-1 >= 0 {
		res = append(res, Point{p.X, p.Y - 1})
	}
	if p.Y+1 < prison.W {
		res = append(res, Point{p.X, p.Y + 1})
	}
	return res
}

// func debug(a [][]int) {
// 	for _, v := range a {
// 		fmt.Printf("%2d", v)
// 		fmt.Println()
// 	}
// }

// 문 따놓은 상태를 어떻게 보존해야 하지?

// 문 따놓은 것만 표시해놓으면 될거같은데
// 두명 가는 길이 다른 경우
// 두명 가는 길이 같은 경우

// 먼저 한명을 방문한다
// 문 어떤 문을 땄는지

// 생각못한것에 쓸것
// 두명이 다른 공간에 있을 수도 있다 - 입구를 2개를 사용해야 한다
// 두명이 같은 공간에 있긴 하지만 입구를 2개를 사용하는 것이 효율적이다

// 죄수부터 시작해서 길찾기? 죄수를 start로 놓고 밖으로 나가는 것을 계산하기

// 리스트를 죄소1 죄수2 두개를 쓰면 만났다는것을 알수있다
// 만약 값이 -1이 아니라면 다른 죄수가 여기를 지나왔다는 것이다
// 리스트는 결국 x,y 형태니까
// 그럼 죄수1이 간 길에 값이 n이면 그걸 죄수2가 그대로 따라가면 되나?

// 일단 예제에서 보기에는 죄수1의 거리 + 죄수2의 거리 하니까 대체로 되는데
// 3예제에서 6+1 5+2 4+3 3+3 해서 6에서 만나는 것을 확인

// 두번째는 dist 값을 죄수1값에 +해서 적어야한다? 6+1을 큐에 넣고 5+2 큐에 넣고... 3+3큐에 넣고

// 경우의수 - 두 죄수가
// 아예 다른 공간에 있다. - 그냥 서치만 되면 됨
// 같은 공간에 있고 한 길로 둘 다 접근한다. - 갈림길이 맵 안에 있다.
// 같은 공간에 있지만 두 입구로 들어가 각각 접근한다. - 갈림길이 맵 밖이다.

// 그럼 갈림길을 어떻게 찾는가?
// 두 죄수 길찾기를 동시에 한다.
// 두 죄수 길찾기를 따로 한다.

//     #
// $#######$

// 그냥 죄수1에 죄수2를 더하는 것은 안 된다. 갈림길 이후의 문들도 다 더해버린다.
