# 가장 긴 증가하는 부분 수열 5

https://www.acmicpc.net/problem/14003

# 풀이

원 수열 $L=\mathit{a}_{1}, \ldots, \mathit{a}_{n}$ 위의 수를 차례로 사용해 가장 긴 부분증가수열의 길이를 추적하는 $K = \mathit{b}_{1}, \ldots, \mathit{b}_{len(K)}$ 을 저장한다.

$L$ 위의 어떤 수 $\mathit{a}_{i}$ 는 각각의 $\mathit{a}_{j}\ (j<i,\mathit{a}_{j}<\mathit{a}_{i})$ 를 가장 큰 수로 가지는 모든 부분증가수열에 포함될 수 있다. 즉 $\mathit{a}_{i}$ 는 어떤 $\mathit{a}_{j}$ 를 포함하는 $\mathit{a}_{i}$ 앞의 가장 긴 부분증가수열에 포함될 수 있다.

- $\mathit{a}_{i}>\mathit{b}_{len(K)}$ 인 경우 $K=\mathit{b}_{1},\mathit{b}_{2} \ldots \mathit{b}_{len(K)},\mathit{a}_{i}$ 로 갱신한다.

- $\mathit{b}_{1}\leq\mathit{a}_{i}\leq\mathit{b}_{len(K)}$ 인 경우 이진 탐색을 사용해 Lower Bound $\mathit{b}_{m}\ (\mathit{b}_{1} \ldots \mathit{b}_{m-1}<\mathit{a}_{i}<\mathit{b}_{m})$ 을 찾은 후 $K=\mathit{b}_{1}, \ldots ,\mathit{b}_{m-1},\mathit{a}_{i},\mathit{b}_{m+1} \ldots ,\mathit{b}_{len(K)}$ 로 갱신한다.

    - 따라서 모든 $\mathit{a}_{m}\ (i<m,\mathit{a}_{i}<\mathit{a}_{m})$ 를 트래킹할 수 있다.

    - $\mathit{a}_{i}$ 와 같은 수가 $K$ 에 이미 있는 경우에는 굳이 갱신할 필요는 없다.

$L$ 위의 어떤 수 $\mathit{a}_{i}$ 가 모든 $\mathit{a}_{j}\ (i>j)$ 보다 작다면 자신만을 갖는 수열 외의 어떠한 부분증가수열에도 포함되지 않는다.

- $\mathit{a}_{i}<\mathit{b}_{1}$ 인 경우 $K=\mathit{a}_{i},\mathit{b}_{2} \ldots \mathit{b}_{len(K)}$ 로 갱신한다. 따라서 $\mathit{a}_{i}<\mathit{a}_{j}<\mathit{b}_{1}$ 인 수들을 트래킹할 수 있다.

- 구현 시에는 위의 로직에 의해 이미 이 개념이 구현되어 있다.

## LIS 출력

$K$ 를 갱신하는 과정에서 가장 긴 부분증가수열을 보존할 수는 없다.

$\mathit{a}_{i}$ 를 이용해 $K$ 를 갱신할 때 $\mathit{a}_{i}$ 가 포함될 수 있는 가장 긴 부분증가수열에 포함되도록 하므로, 그 부분증가수열 내에서 $\mathit{a}_{i}$ 바로 앞에 있는 수가 무엇인지 기억해야 한다.

따라서 $\mathit{a}_{i}$ 가 가리키는 부분증가수열 내의 이전 수를 저장하는, 길이 n의 배열이 필요하다.

$\mathit{b}_{len(K)}$ 부터 시작해 그 수가 어떤 수를 가리키는지 추적하여 LIS를 출력한다.

# 내가 틀린 부분

$\mathit{a}_{i}=\mathit{b}_{j}$ 인 수를 사용해 갱신한 경우, $\mathit{b}_{j-1}$ 이 갱신됐을 수 있으므로 $\mathit{a}_{i}$ 가 가리키는 수를 새로 저장해야 한다.

이럴 필요는 없다. 어차피 같은 수이고, 이 수 앞의 부분증가수열의 길이가 같으므로, 이전 수열이나 새로 갱신된 수열이나 이 수가 있어야 할 자리는 똑같다.

코드에는 등호가 삽입되어 있긴 하다.

# 메모

없음
