#include <iostream>
#include <limits.h>
#include <algorithm>

using namespace std;

typedef struct
{
    int v1, v2, w;
} Edge;

bool compare(Edge a, Edge b)
{
    return a.w < b.w;
}

int find(int parent[], int a)
{
    if (parent[a] == a)
    {
        return a;
    }
    return parent[a] = find(parent, parent[a]);
}

int main(int argc, char const *argv[])
{
    int v, e;
    cin >> v >> e;

    Edge edges[e];
    for (size_t i = 0; i < e; i++)
    {
        Edge &a = edges[i];
        cin >> a.v1 >> a.v2 >> a.w;
    }
    sort(edges, edges + e, compare);

    int parent[v + 1];
    for (size_t i = 0; i <= v; i++)
    {
        parent[i] = i;
    }

    int val = 0;
    int count = 0;

    for (auto &&e : edges)
    {
        if (find(parent, e.v1) != find(parent, e.v2))
        {
            parent[parent[e.v2]] = parent[e.v1];
            val += e.w;
            if (++count == v - 1)
            {
                break;
            }
        }
    }

    cout << val;

    return 0;
}
