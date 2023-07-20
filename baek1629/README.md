# 문제

https://www.acmicpc.net/problem/1629

자연수 A를 B번 곱한 수를 알고 싶다. 단 구하려는 수가 매우 커질 수 있으므로 이를 C로 나눈 나머지를 구하는 프로그램을 작성하시오.

# 내 풀이

bob 암호학 자료 참고

# 풀이

## Modular Exponentiation

### Modular Distributive Property

- $(a + b) \mod n = [(a \mod n) + (b \mod n)] \pmod n$
- $ab \mod n = [(a \mod n)(b \mod n)] \pmod n$
- if $a \equiv b , c \equiv d$ then $ac \equiv bd \pmod m$

### Decimal to Binary

n = n/2, write the remainder until n == 0

reverse the remainders.

### Modular Exponentiation R2L

$n = \mathit{(\mathit{b}_{k}\mathit{b}_{k-1}...\mathit{b}_{0})}_{2}$

$a^n = a^0 * a^1 * a^2 * a^4 * ... * a^\mathit{2^\mathit{bk}}$

$a^n \mod b = a^0 \mod b * a^1 \mod b * a^2 \mod b * a^4 \mod b \ ... \ a^\mathit{2^\mathit{bk}} \mod b$

```
ans := 1
A := a
while n!=0 do
    if n % 2 == 1 then ans = (ans * A) % c

    A = (A * A) % c
    n = n / 2
```

2진수를 먼저 계산해서 저장할 필요 없이, n % 2 마다 맨 오른쪽 (작은쪽) 에 들어가는 나머지가 계산된다. 따라서 R2L 방식으로 처리해야 한다.

이후 n / 2 계산하여 다음 루프로 넘어간다.

A 초기값이 a 이다. -> $a^{2^0} = a$

## 구현
$a^b \mod c$
```go
acc := int64(1)
A := int64(a)
C := int64(c)
for b != 0 {
    if b%2 == 1 {
        acc = (acc * A) % C
    }
    A = (A * A) % C
    b = b / 2
}
```
int32 * int32 값은 항상 int64 값 범위 안이다.

A, acc 값은 항상 C 이하로 유지되며 오버플로우가 발생하지 않는다.

## 에러

없음

# 메모

없음