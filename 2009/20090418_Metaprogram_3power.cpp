//============================================================================
// Name        :
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : ��Ԫ��̷�ʽд3��N�η�(7)
//============================================================================

#include <iostream>
using namespace std;

template<int N>
class Pow3
{
	public:
		enum {result = 3*Pow3<N-1>::result};
};
template<>
class Pow3<0>
{
	public:
		enum {result = 1};
};

int main()
{
  cout << "Pow3<7>: " << (Pow3<7>::result) << endl;

	return 0;
}

