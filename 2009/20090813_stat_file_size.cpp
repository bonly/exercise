#include <iostream>
using namespace std;


int main()
{
    struct stat file_stat;
    if (0 != stat("data.xml", &file_stat))
    {
      cerr << "XML file read error!\n";
      return -1;
    }
    return 0;
}
