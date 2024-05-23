#include <iostream>

using namespace std;

void binSearch(int tracer[], pair<int, int> list[], int low, int high, int val, int curi)
{
    if (high <= low)
    {
        if (list[tracer[high]].first >= val)
        {
            tracer[high] = curi;
            list[curi].second = tracer[high - 1];
        }
        else
        {
            tracer[high + 1] = curi;
            list[curi].second = tracer[high];
        }
        return;
    }

    int mid = (low + high) / 2;
    if (val <= list[tracer[mid]].first)
    {
        binSearch(tracer, list, low, mid - 1, val, curi);
    }
    else
    {
        binSearch(tracer, list, mid + 1, high, val, curi);
    }
}

int main(int argc, char const *argv[])
{
    int n;
    cin >> n;
    pair<int, int> l[n + 1];
    int tracer[n + 1];
    fill_n(tracer, n + 1, 0);
    int tindex = 0;

    for (size_t i = 1; i <= n; i++)
    {
        cin >> l[i].first;

        if (l[tracer[tindex]].first < l[i].first)
        {
            l[i].second = tracer[tindex];
            tindex++;
            tracer[tindex] = i;
        }
        else
        {
            binSearch(tracer, l, 1, tindex, l[i].first, i);
        }
    }

    int p[tindex];
    for (int i = tracer[tindex], j = tindex - 1; i != 0; i = l[i].second, j--)
    {
        p[j] = l[i].first;
    }
    cout << tindex + 1 << endl;
    for (auto &&i : p)
    {
        cout << i << " ";
    }

    return 0;
}
