#include <fstream>
using namespace std;

int main()
{
  fstream of;
  of.open("/tmp/tek.txt",ios::out);
  of << string(' ',255) << endl;
  of << string(' ',255) << endl;
  of << "line 333333333\n";
  of << "444444\n";
  of.seekg(0,ios::beg);
  of << "over write 111111\n";
  of << "write 2222 end\n";
  of.close();
  return 0;
}

/*
证明回到头会把内容复盖
*/
