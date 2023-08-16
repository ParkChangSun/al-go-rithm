# 이항 계수와 쿼리

https://www.acmicpc.net/problem/13977

 
M개의 자연수 N과 정수 K가 주어졌을 때 이항 계수 $\binom{N}{K}$를 1,000,000,007로 나눈 나머지를 구하는 프로그램을 작성하시오.

# 내 풀이

내 풀이가 정답이었다.

# 풀이

백준 11401번과 동일한 풀이이다. 11401번을 잘 풀어놔서 바로 정답을 맞을 수 있었다.

다른 점은 N K 쌍이 여러 개이기 때문에 가장 큰 정수를 특정할 수 없어서 팩토리얼 메모 4_000_000개를 먼저 만들어 놔야 한다는 점 정도이다.

또는 최대값을 뽑아낼 수도 있다.

## 구현
```go
func factorial() {
	f[0] = 1
	for i := 1; i <= MAXLIM; i++ {
		f[i] = f[i-1] * int64(i) % P
	}
}

func inverse(a int64) int64 {
	ans := int64(1)
	for exp := P - 2; exp != 0; exp /= 2 {
		if exp%2 == 1 {
			ans = ans * a % P
		}
		a = a * a % P
	}
	return ans
}

func binomial(n, k int) int64 {
	return f[n] * inverse(f[k]*f[n-k]%P) % P
}

// main
for i := 0; i < n; i++ {
    fmt.Fscanln(reader, &v1, &v2)
    fmt.Fprintln(writer, binomial(v1, v2))
}
```

Modular exponentiation 반복문에서 exp변수 선언과 나누기 연산을 반복문 선언부에 넣을 수 있어서 나름 최적화를 해 봤다.

## 에러

없음

# 메모

없음