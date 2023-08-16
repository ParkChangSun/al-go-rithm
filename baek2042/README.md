# 구간 합 구하기

https://www.acmicpc.net/problem/2042

어떤 N개의 수가 주어져 있다. 그런데 중간에 수의 변경이 빈번히 일어나고 그 중간에 어떤 부분의 합을 구하려 한다. 만약에 1,2,3,4,5 라는 수가 있고, 3번째 수를 6으로 바꾸고 2번째부터 5번째까지 합을 구하라고 한다면 17을 출력하면 되는 것이다. 그리고 그 상태에서 다섯 번째 수를 2로 바꾸고 3번째부터 5번째까지 합을 구하라고 한다면 12가 될 것이다.

# 내 풀이

100% 내 자신의 힘으로 했는지는 모르겠다. 세그먼트 트리 구조에 대해 조금 공부한 뒤 풀어서 해결했다.

# 풀이

# Segment Tree

Segment Tree is a basically a binary tree used for storing the intervals or segments. Each node in the Segment Tree represents an interval.

- The root of Segment Tree will represent the whole array 
- Each leaf in the Segment Tree will represent a single element 
- The internal nodes in the Segment Tree represents the union of elementary intervals

Since a Segment Tree is a binary tree, a simple linear array can be used to represent the Segment Tree. Before building the Segment Tree, one must figure what needs to be stored in the Segment Tree's node?

For example, if the question is to find the sum of all the elements in an array from indices, then at each node (except leaf nodes) the sum of its children nodes is stored.

## Build

- Start from the root and recurse on the left and the right child until a leaf node is reached.
- From the leaves, recurse back to the root and update all the nodes in the path.
- Each nodes stores interval that are the union of elementary intervals.
- Each leaf nodes stores each element of array.

Since Segment Tree is a binary tree, for a node $N$, $2N$ and $2N+1$ will represent the left node and right node. Tree's node index $N$ starts with 1 at the root node.

## Query

To query on a given range, check 3 conditions.

- Range represented by a node is completely inside the given range
- Range represented by a node is completely outside the given range
- Range represented by a node is partially inside and partially outside the given range

## Update

To update an element, look at the interval in which the element is present and recurse accordingly on the left or the right child.

## 구현

### 내 구현
```go
func generateTree(h, t, index int) int {
	if h == t {
		tree[index] = arr[h]
		return tree[index]
	}

	b := (h + t) / 2
	tree[index] = generateTree(h, b, index*2) + generateTree(b+1, t, index*2+1)
	return tree[index]
}

func replace(index, head, tail, treeIndex, num int) {
	if head == tail {
		tree[treeIndex] = num
		arr[index] = num
		return
	}

	diff := num - arr[index]
	tree[treeIndex] += diff
	b := (head + tail) / 2

	if index <= b {
		replace(index, head, b, treeIndex*2, num)
	} else {
		replace(index, b+1, tail, treeIndex*2+1, num)
	}
}

func calc(from, to, head, tail, index int) int {
	if from == head && to == tail {
		return tree[index]
	}

	// 탐색 중인 범위 안에 더하고자 하는 숫자가 없는 경우
	if to < head || tail < from {
		return 0
	}

	ap, dui := 0, 0
	b := (head + tail) / 2

	if to < b {
		// 뒤를 계산할 필요가 없는 경우
		// 더하고자 하는 범위의 끝이 중간값보다 작다
		ap = calc(from, to, head, b, index*2)
		dui = 0
	} else if b < from {

		// 앞을 계산할 필요가 없는 경우
		// 더하고자 하는 범위의 시작이 중간값보다 크다
		ap = 0
		dui = calc(from, to, b+1, tail, index*2+1)
	} else {
		ap = calc(from, b, head, b, index*2)
		dui = calc(b+1, to, b+1, tail, index*2+1)
	}

	return ap + dui
}
```

### 정석 구현

```c
void build(int node, int start, int end)
{
    if(start == end)
    {
        // Leaf node will have a single element
        tree[node] = A[start];
    }
    else
    {
        int mid = (start + end) / 2;
        // Recurse on the left child
        build(2*node, start, mid);
        // Recurse on the right child
        build(2*node+1, mid+1, end);
        // Internal node will have the sum of both of its children
        tree[node] = tree[2*node] + tree[2*node+1];
    }
}

void update(int node, int start, int end, int idx, int val)
{
    if(start == end)
    {
        // Leaf node
        A[idx] += val;
        tree[node] += val;
    }
    else
    {
        int mid = (start + end) / 2;
        if(start <= idx and idx <= mid)
        {
            // If idx is in the left child, recurse on the left child
            update(2*node, start, mid, idx, val);
        }
        else
        {
            // if idx is in the right child, recurse on the right child
            update(2*node+1, mid+1, end, idx, val);
        }
        // Internal node will have the sum of both of its children
        tree[node] = tree[2*node] + tree[2*node+1];
    }
}

int query(int node, int start, int end, int l, int r)
{
    if(r < start or end < l)
    {
        // range represented by a node is completely outside the given range
        return 0;
    }
    if(l <= start and end <= r)
    {
        // range represented by a node is completely inside the given range
        return tree[node];
    }
    // range represented by a node is partially inside and partially outside the given range
    int mid = (start + end) / 2;
    int p1 = query(2*node, start, mid, l, r);
    int p2 = query(2*node+1, mid+1, end, l, r);
    return (p1 + p2);
}
```

### 차이점

내 구현에서는 update에 원래의 값과 바꿀 값의 차이를 미리 그 노드에 계산한 후 재귀함수를 실행하는데, 정석 구현에서는 재귀가 끝나고 난 후 계산을 다시 실행한다. 이 경우가 곱셈이나 다른 연산에 사용할 수도 있는 이 방식이 더 좋은 것 같다. 이 문제는 그냥 더하기 연산이라 상관없었다.

query에서 탐색 범위 시작과 끝의 중간이 찾고자 하는 범위의 앞에 있는지, 안에 있는지, 뒤에 있는지를 굳이 생각해서 조건을 나눴는데, 그럴 필요는 없었다. 이미 나뉘는 앞부분과 뒷부분 중에 없으면 0을 리턴하는 조건도 이미 만들어 놨는데 나누는 것이 소용이 없었지만, 풀 때는 그냥 놔뒀었다.

## 에러

없음

# 메모

없음