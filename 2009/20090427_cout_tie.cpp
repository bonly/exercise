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
	*cin.tie() << "this is inserted into cout \n";//cin�ɿ�Ĭ��tie��cout

	oldtie = cin.tie(&ofs); //����tieͬʱ�ѷ��صľɰ�ostream������oldtie;
	ostream *mystr = cin.tie();
	*cin.tie() << "this is inserted into the file\n";
	*mystr << "this for mystr\n";
	//*oldtie << "this is inserted into prevstr\n";//��ʱ�ɵ�ostream��û�й������쳣
	cin.tie (oldtie);//�Ѿɵ�cout��ostream�ָ���cin��
	ofs.close();

  return 0;
}

