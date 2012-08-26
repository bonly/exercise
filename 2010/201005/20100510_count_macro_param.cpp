#include <iostream>
using namespace std;

#define VA_NUM(...) VA_NUM_IMPL(__VA_ARGS__,13,12,11,10,9,8,7,6,5,4,3,2,1)
#define VA_NUM_IMPL(_1,_2,_3,_4,_5,_6,_7,_8,_9,_10,_11,_12,_13,N,...) N

int main()
{
 cout << VA_NUM(1,2,3,4,5,6,7,8,9,0) << endl;
 cout << VA_NUM(1,2) << endl;
 return 0;
}

