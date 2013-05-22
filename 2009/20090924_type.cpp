#include <iostream>
#include <string>
using namespace std;

template<typename T>
int foo(const T)
{
  return T.size();
}

template<typename T>
class fool
{
  typedef T::size_type result;
};


int
main()
{
  clog << "foo: " << foo("cccc") << endl;
  clog << "foo<string>::result = " << fool<string>::result << endl;
  return 0;

}