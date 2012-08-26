#include <iostream>
using namespace std;

static void fun(int a)
{
    cout << "a: " << a << endl;
}

static void fun(char* const  a)
{
  cout<< "a: " << a <<endl;
}

class A
{
  public:
      static void fc(int a)
      {
          cout << "a: " << a << endl;
      }
      static void fc(char * const a)
      {
          cout << "a: " << a << endl;
      }
      static long tm(long a)
      {
          cout << "cm: "<< a<<endl;
          return 34;
      }
      ///重载不能只通过返回值来区分
      static long long tm(long a) ///< cann't overload with only return value type
      {
          cout << "ll cm: " << a<< endl;
          return 341341LL;
      }
};

int main()
{
  fun(12);
  fun("this is a test");
  A a;
  a.fc(13);
  a.fc("another test");
  a.tm(134L);
  long long k = a.tm(344L);
  return 0;
}
