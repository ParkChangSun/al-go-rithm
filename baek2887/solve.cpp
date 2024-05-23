#include <iostream>
#include <algorithm>
#include <cstring>
#include <cstdlib>

using namespace std;

typedef struct
{
    int n, x, y, z;
} Planet;
bool sortx(Planet l1, Planet l2)
{
    return l1.x < l2.x;
};
bool sorty(Planet l1, Planet l2)
{
    return l1.y < l2.y;
};
bool sortz(Planet l1, Planet l2)
{
    return l1.z < l2.z;
};

typedef struct
{
    int v1, v2, w;
} Tunnel;
bool sortw(Tunnel t1, Tunnel t2)
{
    return t1.w < t2.w;
}

Planet planets[100000];
Tunnel tunnels[300000];

int parent[100000];
int findParent(int a)
{
    if (parent[a] == a)
    {
        return a;
    }
    return parent[a] = findParent(parent[a]);
}

int main(int argc, char const *argv[])
{
    int n;
    cin >> n;
    for (size_t i = 0; i < n; i++)
    {
        cin >> planets[i].x >> planets[i].y >> planets[i].z;
        planets[i].n = i;
    }

    int c = 0;
    sort(planets, planets + n, sortx);
    for (size_t i = 0; i < n - 1; i++)
    {
        Tunnel &t = tunnels[c];
        t.w = planets[i + 1].x - planets[i].x;
        t.v1 = planets[i].n;
        t.v2 = planets[i + 1].n;
        c++;
    }
    sort(planets, planets + n, sorty);
    for (size_t i = 0; i < n - 1; i++)
    {
        Tunnel &t = tunnels[c];
        t.w = planets[i + 1].y - planets[i].y;
        t.v1 = planets[i].n;
        t.v2 = planets[i + 1].n;
        c++;
    }
    sort(planets, planets + n, sortz);
    for (size_t i = 0; i < n - 1; i++)
    {
        Tunnel &t = tunnels[c];
        t.w = planets[i + 1].z - planets[i].z;
        t.v1 = planets[i].n;
        t.v2 = planets[i + 1].n;
        c++;
    }

    sort(tunnels, tunnels + c, sortw);

    for (size_t i = 0; i < n; i++)
    {
        parent[i] = i;
    }

    int ans = 0;
    for (int i = 0, tCount = 0; tCount < n - 1; i++)
    {
        Tunnel &t = tunnels[i];
        if (findParent(t.v1) != findParent(t.v2))
        {
            parent[parent[t.v1]] = parent[t.v2];
            ans += t.w;
            tCount++;
        }
    }

    cout << ans;

    return 0;
}
