# 이항 계수 3

https://www.acmicpc.net/problem/11401

자연수 N과 정수 K가 주어졌을 때 이항 계수 $\binom{N}{K}$를 1,000,000,007로 나눈 나머지를 구하는 프로그램을 작성하시오.

첫째 줄에 N과 K가 주어진다. (1 ≤ N ≤ 4,000,000, 0 ≤ K ≤ N)

# 내 풀이

파스칼의 삼각형을 사용해 $\binom{N}{K} = \binom{n-1}{k-1} + \binom{n-1}{k}$을 구하는 방식을 사용했다. 다이나믹 프로그래밍을 사용해 이항계수를 구하려면 N*K 크기의 메모를 이용해야 한다. 하지만 메모리 제한이 걸려 있어서 불가능했다.

이후 modular inverse를 사용해서 성공했다.

# 풀이

## Fermat's Little Theorem

$x$ is modular inverse of $a$. $x$ is $a^{-1}$.

$ax \equiv 1 \pmod{m}$

Let us assume that a is positive and not divisible by p.

$a,2a,3a,\ldots ,(p-1)a$

and reduce each one modulo p, the resulting sequence turns out to be a rearrangement of

$1,2,3,\ldots ,p-1$

Furthermore, the numbers a, 2a, ..., (p − 1)a must all be distinct after reducing them modulo p, because if

$ka\equiv ma{\pmod {p}}$ where k and m are one of 1, 2, ..., p − 1, then the cancellation law tells us that

$k\equiv m{\pmod {p}}$ k and m are between 1 and p − 1, they must be equal.

Therefore, if we multiply together the numbers in each sequence, the results must be identical modulo p:

$a^{p-1}(p-1)!\equiv (p-1)!{\pmod {p}}$

Finally, we may “cancel out” the numbers 1, 2, ..., p − 1 from both sides of this equation, obtaining

$a^{p-1} \equiv 1 \pmod p$

(Euler's Theorem $a^{\varphi (n)}\equiv 1{\pmod {n}}$)

구현에서는 $a^{p-2} \equiv a^{-1} \pmod p$를 사용한다.

### Binomial Coefficient

$\binom{n}{k} = n!/k!(n-k)!$

$\binom{n}{k} \equiv n! * (k!(n-k)!)^{-1} \pmod p$

## 구현
```go
func binomial(a, b int) int64 {
	return (f[a] * inverse((f[b]*f[a-b])%p)) % p
}

func factorial() {
	f[0] = 1
	for i := 1; i <= n; i++ {
		f[i] = (int64(i) * f[i-1]) % p
	}
}

func inverse(a int64) int64 {
	t := p - 2
	ans := int64(1)
	acc := int64(a)
	for t != 0 {
		if t%2 == 1 {
			ans = (ans * acc) % p
		}
		acc = (acc * acc) % p
		t /= 2
	}
	return ans
}
```

팩토리얼을 먼저 전부 계산해서 배열에 저장한 후 계산한다. 이 문제에서는 굳이 팩토리얼을 저장할 필요는 없었다.

역원 구하는 함수는 Modular Exponentiation을 사용한다.

## 에러

없음

# 메모

없음