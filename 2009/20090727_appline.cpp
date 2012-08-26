#include <fstream>
using namespace std;

int main()
{
  fstream of;
  of.open("/tmp/tek.txt",ios::out);
  of << "line 333333333\n";
  of << "444444\n";
  //of.seekg(0,ios::beg);
  of.close();

  ifstream in;
  in.open("/tmp/tek.txt",ios::in);
  
  of.open("/tmp/tekret.txt",ios::out);
  of << "over write 111111\n";
  of << "write 2222 end\n";

  char c=0;
  while(in.get(c).good())
  {
     of << c;
  }
  in.close();
  of.close();
  return 0;
}

/*
写完后关闭文件再重新读入并补到另一文件中
int get();
istream& get ( char& c );
istream& get ( char* s, streamsize n );
istream& get ( char* s, streamsize n, char delim );
istream& get ( streambuf& sb);
istream& get ( streambuf& sb, char delim );
*/
