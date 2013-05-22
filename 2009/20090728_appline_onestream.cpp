#include <fstream>
using namespace std;

int main()
{
  fstream fs;
  fs.open("/tmp/tek.txt",ios::out|ios::binary);
  fs << "line 333333333\n";
  fs << "444444\n";
  //fs.flush();
  //fs.seekg(ios::beg);
  fs.close();

  fs.open("/tmp/tek.txt",ios::in);
  ofstream is;
  is.open("/tmp/tekret.txt",ios::out);
  is << "over write 111111\n";
  is << "write 2222 end\n";

  char c=0;
  while(fs.get(c).good())
  {
     is << c;
  }
  fs.close();
  is.close();
  return 0;
}

/*
写完后关闭文件再重新读入并补到另一文件中
flush()了再读不起作用,需close开open 来读
int get();
istream& get ( char& c );
istream& get ( char* s, streamsize n );
istream& get ( char* s, streamsize n, char delim );
istream& get ( streambuf& sb);
istream& get ( streambuf& sb, char delim );
*/
