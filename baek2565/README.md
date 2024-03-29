# 전깃줄

https://www.acmicpc.net/problem/2565

두 전봇대 A와 B 사이에 하나 둘씩 전깃줄을 추가하다 보니 전깃줄이 서로 교차하는 경우가 발생하였다. 합선의 위험이 있어 이들 중 몇 개의 전깃줄을 없애 전깃줄이 교차하지 않도록 만들려고 한다.

예를 들어, < 그림 1 >과 같이 전깃줄이 연결되어 있는 경우 A의 1번 위치와 B의 8번 위치를 잇는 전깃줄, A의 3번 위치와 B의 9번 위치를 잇는 전깃줄, A의 4번 위치와 B의 1번 위치를 잇는 전깃줄을 없애면 남아있는 모든 전깃줄이 서로 교차하지 않게 된다.

전깃줄이 전봇대에 연결되는 위치는 전봇대 위에서부터 차례대로 번호가 매겨진다. 전깃줄의 개수와 전깃줄들이 두 전봇대에 연결되는 위치의 번호가 주어질 때, 남아있는 모든 전깃줄이 서로 교차하지 않게 하기 위해 없애야 하는 전깃줄의 최소 개수를 구하는 프로그램을 작성하시오.

## 풀이

### 내 풀이

가장 긴 증가하는 수열 LIS라는 힌트가 있긴 해도 구조 자체가 LIS라는 것을 알 수 있었다.
A->B 전깃줄 중에서 B가 중간에 커지는 것은 상관 없는데 줄어드는 경우에는 무조건 다른 줄과 겹치기 때문이다.

문제를 풀 때는 LIS 알고리즘을 모르는 상태였다.
이전의 시퀀스를 사용해서 N에서의 시퀀스를 구할 수 있기 때문에 문제를 나누는 데 집중했다.

패턴을 생각해 보기 위해 이 수열을 사용했다.

[1 99 100 70 80 10 82 3 4 5 2 90 89]

1. N에서 가장 가까우면서 작은 수가 있는 자리의 수열에 N을 더하기
    
    82에서는 10이 가장 가까우면서 작은 수다. 10에서는 [1 10]이 가장 긴 수열이다.

2. N-1에서의 가장 긴 수열에 N을 더하기

    90에서 계산을 할 때 2를 사용하게 된다. 2에서는 [1 2]가 가장 긴 수열이다.

3. N보다 작은 수열 중 가장 긴 수열에 N을 더하기

    루프를 한 번 더 돌아야 하지만 이것밖에 방법이 없었다. 필연적으로 N보다 작지만 가장 긴 수열에 붙어야 N에서의 가장 긴 수열을 구할 수 있다.

### 구현
```go
func prune(n int) int {
	if memo[n] != 0 {
		return memo[n]
	}

	acc := 0
	for i := 1; i < n; i++ {
		if t := prune(i); l[i] < l[n] && acc < t {
			acc = t
		}
	}
	memo[n] = acc + 1
	return acc + 1
}
```
알고리즘을 생각해 내는 것이 어렵지 구현은 딱히 어려운 부분이 없었다. 루프를 사용하는 bottom-up 방식도 생각해 내기 어렵지 않다.

### 에러

답이 N번째 수를 포함하지 않는 경우가 일반적이기 때문에 메모의 최대값을 구해야 했다. 그러나 백준의 예제가 메모의 마지막에 답이 있었다. 이 문제를 찾느라 시간을 꽤 썼었다. 내가 구하고자 하는 답을 어디에서 찾아야 하는지 생각해 보는 절차가 필요할 것 같다.

```go
for i := 0; i < n; i++ {
    lines[i] = make([]int, 2)
    fmt.Fscanln(reader, &lines[i][0], &lines[i][1])
}
```
`var a,b int` 이렇게 선언하는 것보다 더 효율적이다.
