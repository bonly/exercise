#include <fstream>
using namespace std;

int main()
{
  fstream of;
  of.open("/tmp/tek.txt",ios::out);
  of << "line 1 " << endl;
  of << "\n";
  of << "line 3\n";
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
