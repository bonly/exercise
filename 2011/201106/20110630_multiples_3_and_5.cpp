#include <iostream>

using namespace std;

int main()
{
    int tal1 = 0;
    int summa = 0;
    
    while (tal1 < 1000)
    {
        if (tal1 % 3 == 0 || tal1 % 5 == 0)
        {
            summa+=tal1;
            tal1+=1;
        }
        else
        {
            tal1+=1;
        }
    }
    cout << summa << "\n";
    //system("pause");    
    return 0;
}
