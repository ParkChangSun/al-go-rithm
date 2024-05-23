#include <iostream>
#include <cstring>

using namespace std;

int memo[101][101][101];
bool c[101][101][101];
const int p = 1000000007;

int solve(int n, int l, int r)
{
    if (l == 0 || r == 0 || n == 0)
    {
        return 0;
    }

    if (!c[n][l][r])
    {
        int a = solve(n - 1, l, r);
        int case1 = 0;
        for (int i = 0; i < n - 2; i++)
        {
            case1 = (case1 + a) % p;
        }
        int case2 = (solve(n - 1, l - 1, r) + solve(n - 1, l, r - 1)) % p;
        memo[n][l][r] = (case1 + case2) % p;
        c[n][l][r] = true;
    }
    return memo[n][l][r];
}

int main(int argc, char const *argv[])
{
    int n, l, r;
    cin >> n >> l >> r;
    memo[1][1][1] = 1;
    c[1][1][1] = true;
    cout << solve(n, l, r) << endl;

    return 0;
}
