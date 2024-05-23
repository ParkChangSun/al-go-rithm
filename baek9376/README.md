# 탈옥

https://www.acmicpc.net/problem/9376

# 내 풀이 유도 과정

- 구출자가 01bfs로 두 곳의 거리의 합을 구할 수 있나? 아닌 것 같다.
- 그렇다면 역으로 죄수 두 명이 문을 따면서 밖으로 나오는 경우를 구할 수 있을 것 같다. - 하지만 두 죄수가 완전히 분리된 공간에 있다면? 01bfs니까 감옥 밖에 한칸씩 확장하면 동일한 공간으로 볼 수 있겠다.
- 그렇다면 그 분기점을 어떻게 구하는가? 어떤 문으로 나가야 하는가? 단순히 두 명의 거리를 더하기만 하는 것으로는 겹치는 경로의 거리를 뺄 방법이 없다.

# 풀이

밖의 구출자가 최단 경로로 두 죄수을 구출하는 경우는 다음과 같다.

- 하나의 입구로 들어가 두 명을 구출하는 경우
    - 두 명이 분리되어 있어 구출 경로에 분기가 존재하는 경우
    - 일직선으로 두 명을 구출할 수 있는 경우
- 두 개의 입구로 들어가 따로따로 구출하는 경우

케이스 일반화

- 구출 경로에 분기점이 있다면 감옥 안에 존재한다. 구출 경로가 일직선이라면 죄수 중 하나의 위치가 분기점이다.
- 구출자가 감옥 밖을 자유로이 돌아다닐 수 있다면 구출 경로의 분기점이 감옥 밖에 존재하는 것으로 생각할 수 있다.

따라서 감옥 밖에 가중치가 0인 공간을 확장하면 동일한 알고리즘을 적용할 수 있다.

구출자가 분기점이 있는 구출 경로를 지나는 것은 구출자와 두 죄수가 분기점으로 모이는 것으로 생각할 수 있다. 즉 구출 거리는 분기점에서 세 명의 시작 위치의 거리의 합이다. 또한 분기점이 문인 경우에는 한 명만 열면 되므로 합에서 2를 뺀다.

## 구현
```go
// main
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
```

코드의 Prison 구조체는 Height, Width 값을 가지고, Point 구조체는 X, Y 값을 가진다. 하지만 실제로 사용될 때에는 Height 를 사용할 때 X 값을 사용하고 Width 를 사용할 때 Y 값을 사용한다. 다음에는 헷갈릴 수도 있으니 Point 도 마찬가지로 H, W 로 이름을 지정하는 것이 좋을 것 같다.

# 메모

이번 문제에서 생각해 내지 못한 것은 감옥 맵의 밖에 가중치가 0인 공간을 확장하지 못한 것, 그리고 죄수 두명의 bfs만 하려고 했지 죄수와 구출자까지 3명의 bfs를 구해 더하는 방법은 생각해 내지 못한 것이다. 

분기에서 만나거나 분기 경로를 계산하는 문제를 다시 만난다면 이 문제처럼 각 사람들로부터의 거리를 계산하는 아이디어를 떠올릴 수 있을 것이다.