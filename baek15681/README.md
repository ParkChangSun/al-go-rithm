# 트리와 쿼리

https://www.acmicpc.net/problem/15681

# 풀이

입력이 undirected edge이다. 비 재귀 풀이에서는 어떤 노드의 부모 노드가 무엇인지 알아야 하지만 재귀 풀이에서는 필요가 없다.

## 재귀 풀이

$\mathit{C}_{R}=1+\mathit{C}_{r1}+\mathit{C}_{r2} \ldots$

어떤 노드의 부모 노드를 다시 방문하지 않기 위해 방문한 노드들을 표시하는 visited배열을 사용한다.

## 스택 비 재귀 풀이

스택을 사용해 계산하려면 R of $\mathit{T}_{R}$ 을 마지막으로 방문해야 한다. 즉 자식 노드가 루트보다 먼저 pop되어야 한다.

checkParent 스택을 사용해 parent배열을 기록한다. 루트 노드를 방문으로 표시하고, 자식 노드들의 부모 노드를 루트 노드로 저장하고, 자식 노드의 서브트리들에 대해 부모 노드 연산을 한다. (인풋이 방향없는 엣지) 동시에 postOrder 스택을 사용해 후위 순회 하도록 스택에 push한다. 따라서 루트 노드를 먼저 넣는다.

postOrder 스택에서 pop하면서 해당 노드의 부모 노드에 값을 더한다.

# 내가 틀린 부분

쿼리의 개수가 10^5 라서 매번 계산할 수 없었다. 메모이제이션을 사용해야 했다.

재귀로 하면 재귀 초과될 줄 알았는데 그냥 재귀로도 되더라.

# 메모

없음