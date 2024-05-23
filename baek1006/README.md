# 습격자 초라기

https://www.acmicpc.net/problem/1006

# 풀이

## 문제 분할

만약 두 구역을 합칠 수 있다면 위 또는 아래의 두 구역을 합치거나 위아래로 합쳐야 한다.

원형 구조를 펴서 일자로 만들어 생각한다.

1번부터 k번 자리까지 배치한 팀 수 = 1번부터 k-1번 자리까지 배치한 팀 수 + k번 자리에 배치한 팀 수

위와 같은 접근 방식으로 시작한다.

## k번 자리를 배치하는 경우

1. $\mathit{U}_{k}$ 만 배치한다.

    $\mathit{A}_{k} = \text{min} \begin{cases} \mathit{B}_{k-1}+1 & \text{if }\mathit{U}_{k-1}+\mathit{U}_{k} \leq W\\ \mathit{C}_{k-1}+1\end{cases}$

2. $\mathit{L}_{k}$ 만 배치한다.

    $\mathit{B}_{k} = \text{min} \begin{cases} \mathit{A}_{k-1}+1 & \text{if }\mathit{L}_{k-1}+\mathit{L}_{k} \leq W\\ \mathit{C}_{k-1}+1\end{cases}$

3. $\mathit{U}_{k},\mathit{L}_{k}$ 모두 배치한다.

    $\mathit{C}_{k} = \text{min} \begin{cases} \mathit{A}_{k}+1 \\ \mathit{B}_{k}+1 \\ \mathit{C}_{k-1}+1 & \text{if } \mathit{U}_{k}+\mathit{L}_{k} \leq W \\ \mathit{C}_{k-1}+2 \\ \mathit{C}_{k-2}+2 & \text{if } \mathit{U}_{k-1}+\mathit{U}_{k} \leq W \text{ and }\mathit{L}_{k-1}+\mathit{L}_{k} \leq W \end{cases}$

따라서 3개의 1차원 배열 메모가 필요하다.

'k번의 모든 칸이 채워지는 경우' 로 분할하지 않는 이유는 n번 자리를 계산할 때 이미 배치한 1번 자리에 간섭하기 때문이다.

EX>k번의 위쪽 칸은 k+1번과 합쳐서 배치하고, 아래 칸은 합치지 않고 배치한다.

## 1번과 n번 자리를 합치는 경우

원형 구조를 폈기 때문에 1번 자리와 n번 자리는 원래 붙어 있다.

1. 1번과 n번 자리를 합치지 않은 경우 단순히 계산하면 된다. 따라서 답은 $\mathit{C}_{n}$
2. $\mathit{U}_{1}, \mathit{U}_{n}$ 을 합친 경우 $\mathit{L}_{n}$ 까지만 배치하면 된다. 따라서 답은 $\mathit{B}_{n}$
3. $\mathit{L}_{1}, \mathit{L}_{n}$ 을 합친 경우 $\mathit{U}_{n}$ 까지만 배치하면 된다. 따라서 답은 $\mathit{A}_{n}$
4. $\mathit{U}_{1}, \mathit{U}_{n}$ 과 $\mathit{L}_{1}, \mathit{L}_{n}$ 을 모두 합친 경우 n-1번 구역까지만 배치하면 된다. 따라서 답은 $\mathit{C}_{n-1}$

위 네 가지 경우 중 가장 작은 값이 답이 된다.

# 내가 틀린 부분

처음에는 k번의 모든 칸을 채우는 경우 로 문제를 분할했다. 그리고 n번을 계산할 때 1번 자리에 간섭하는 것을 어떻게 처리할까를 계속 고민했다. 여기서 발상의 전환을 해서 위 풀이처럼 변형할 수 있었다면 힌트 없이 풀 수 있었을 것이다.

그동안 k번째 자리를 계산할 때 완벽하지 않은 경우, 즉 둘 중 한 칸만 채우는 경우로 생각해 본 적이 없었던 것 같다. 이 문제는 그걸 알려줬다.

# 메모

없음
