//============================================================================
// Name        :
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : 用元编程方式写求平方根
//============================================================================

#include <iostream>
using namespace std;

template<int N, int LO=1, int HI=N>
class Sqrt
{
	public:
		enum { mid = (LO + HI +1)/2 }; //计算中点
		enum { result = (N<mid*mid) ? Sqrt<N,LO,mid-1>::result
                	                    : Sqrt<N,mid,HI>::result };
};

template<int N, int M>
class Sqrt<N,M,M>
{
	public:
		enum {result = M};
};

int main()
{
  cout << "Sqrt<9>: " << (Sqrt<9>::result) << endl;

	return 0;
}

