#include <iostream>
#include <fstream>

using namespace std;

int
main ()
{
	ostream *oldtie;
	ofstream ofs;
	ofs.open ("text.txt");
	cout << "tie example: "<<endl;
	*cin.tie() << "this is inserted into cout \n";//cin由库默认tie了cout

	oldtie = cin.tie(&ofs); //更改tie同时把返回的旧绑定ostream保存在oldtie;
	ostream *mystr = cin.tie();
	*cin.tie() << "this is inserted into the file\n";
	*mystr << "this for mystr\n";
	//*oldtie << "this is inserted into prevstr\n";//此时旧的ostream还没有归属会异常
	cin.tie (oldtie);//把旧的cout的ostream恢复到cin中
	ofs.close();

  return 0;
}

