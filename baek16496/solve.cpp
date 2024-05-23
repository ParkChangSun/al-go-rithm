#include <iostream>
#include <algorithm>
#include <string>

using namespace std;

bool comp(int a, int b)
{
    string as = to_string(a);
    string bs = to_string(b);
    return stoull(as + bs) > stoull(bs + as);
}

int l[1000];

int main(int argc, char const *argv[])
{
    int n;
    cin >> n;
    for (size_t i = 0; i < n; i++)
    {
        cin >> l[i];
    }

    sort(l, l + n, comp);

    string ans;
    bool zeroFlag = true;
    for (size_t i = 0; i < n; i++)
    {
        if (l[i] != 0)
        {
            zeroFlag = false;
        }

        ans = ans + to_string(l[i]);
    }

    if (zeroFlag)
    {
        cout << 0;
    }
    else
    {
        cout << ans;
    }

    return 0;
}
