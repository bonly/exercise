#include <iostream>
using namespace std;


int main()
{
    for (int i = 0; i < 10; ++i)
    {
        if (i == 8)
        {
            clog << i << endl;
        }
        else 
        {
            if (i == 9)
            {
                clog << "==9" << endl;
                break;
            }
            clog << "in esle" << endl;
        }
        clog << i << endl;
    }
    return 0;
}
