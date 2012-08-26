#include <iostream>
#include <string>
#include <boost/format.hpp>
using namespace std;
using namespace boost;

void _hex(const char c)
{
	//cout << hex << (int)c << ' ';
	cout << format("%02X ")%(int)c;
}

int
main()
{
	string str ("test string");
	unsigned int i=0;
	for(const char* p = str.c_str(); i <= str.length(); ++p,++i)
	{
		_hex(*p);
	}
	return 0;
}

