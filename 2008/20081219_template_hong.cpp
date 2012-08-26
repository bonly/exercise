#include <iostream>

using namespace std;

class Base
{
  public:
    virtual void who(){cout <<"Base\n";}
};

class Child
{
  public:
    virtual void who(){cout << "Child\n";}
};

template<typename input, typename cs>
void
fun(input a,cs myc)
{
  cout << "your input is: " << a << endl;
#define RETURN(sd) {sd->who();}
  RETURN(myc);
}


int 
main()
{
  Base b;
  Child c;
  fun("this is a test", &b); 
  fun("this is second test", &c); 
  return 0;
}

