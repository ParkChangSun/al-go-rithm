#include <iostream>

using namespace std;

int box[1 << 21];
int offset = (1 << 20) - 1;

void addValLeaf(int index, int val)
{
    for (int i = index; i > 0; i /= 2)
    {
        box[i] += val;
    }
}

int main(int argc, char const *argv[])
{
    cin.tie(NULL);
    cout.tie(NULL);
    ios_base::sync_with_stdio(false);

    int n, a, b, c;
    cin >> n;
    for (size_t i = 0; i < n; i++)
    {
        cin >> a;
        if (a == 1)
        {
            cin >> b;
            int index = 1;
            for (size_t t = 0; t < 20; t++)
            {
                index *= 2;
                if (box[index] < b)
                {
                    b -= box[index];
                    index++;
                }
            }
            addValLeaf(index, -1);
            cout << index - offset << '\n';
        }
        else
        {
            cin >> b >> c;
            addValLeaf(b + offset, c);
        }
    }

    return 0;
}
