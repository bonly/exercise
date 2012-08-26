#include <iostream>
#include <cstdio>
#include <boost/lexical_cast.hpp>

int main()
{
  using namespace std;
  cout << "OK" << endl;
  cout << "sizeof (long long): " << sizeof(long long) << endl;
  cout << "sizeof (long): " << sizeof(long) << endl;
  cout << "sizeof (char): " << sizeof(char) << endl;
  //unsigned long long lg=13719360007;
  unsigned long long int lg=13719360007uLL;
  cout << "long long 13719360007 is: " <<  lg << endl;
  printf ("lg : %llu \n",lg);

  char ph[11];
  strncpy(ph,"13719350001",11);
  //ph[12]='\0';
  try
  {
    unsigned long long phone = boost::lexical_cast<unsigned long long>(ph);
    printf("cast ull is: %llu \n",phone);
  }
  catch(boost::bad_lexical_cast & e)
  {
    std::cerr << e.what();
  }

  return EXIT_SUCCESS;
}
