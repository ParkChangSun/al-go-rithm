#include <iostream>
#include <set>
#include <string>
#include <queue>

using namespace std;

string answer = string("123456780");
set<string> s;
queue<string> q;

void checkChild(string t, int i, int j)
{
    t[i] = t[j];
    t[j] = '0';
    if (s.insert(t).second)
    {
        q.push(t);
    }
}

int solve()
{
    for (int move = 0; !q.empty(); move++)
    {
        for (size_t n = 0, size = q.size(); n < size; n++)
        {
            string t = q.front();
            q.pop();
            if (t == answer)
            {
                return move;
            }

            int i = t.find('0');
            if (int j = i - 1; j >= 0 && j / 3 == i / 3)
            {
                checkChild(t, i, j);
            }
            if (int j = i + 1; j < 9 && j / 3 == i / 3)
            {
                checkChild(t, i, j);
            }
            if (int j = i - 3; j >= 0)
            {
                checkChild(t, i, j);
            }
            if (int j = i + 3; j < 9)
            {
                checkChild(t, i, j);
            }
        }
    }

    return -1;
}

int main(int argc, char const *argv[])
{
    string board;
    char c;
    for (size_t i = 0; i < 9; i++)
    {
        cin >> c;
        board.push_back(c);
    }
    q.push(board);

    cout << solve();

    return 0;
}
