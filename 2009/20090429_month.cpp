#include <iostream>
#include <string>
using namespace std;

struct MONTH
{
  char year[5];
  char month[3];
  MONTH(const char* data)
  {
    memset (this,0,sizeof(MONTH));
    char tmonth[3];
    memset (tmonth,0,3);

    strncpy (year,data,4);
    strncpy (tmonth,data+4,2);

    int tmon = atoi(tmonth);
    sprintf(month,"%d",tmon);
  }
};

int
main ()
{
	string tm("20090303");
	MONTH month(tm.c_str());
	cout << "month: "<<month.month<< "\n";
	return 0;
}

