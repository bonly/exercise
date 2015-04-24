#define BOOST_SPIRIT_THREADSAFE
#include <boost/spirit/include/classic.hpp>
#include <boost/bind.hpp>
#include <string>
#include <iostream>
using namespace std;
using namespace BOOST_SPIRIT_CLASSIC_NS;
using namespace boost;

string setup("t_need(2011)(3011)");

struct mydata{
  int myint;
};

struct myop{
  myop(mydata &val):value(val){
  }
  void operator() (int val) const{
    value.myint = val;
  }
  
  mydata &value;
};

struct myfunc{
  myfunc(mydata &val):value(val){
  }
  
  void operator() (char const* first, char const* last) const{
    std::string str(first, last);
    std::cout << "processing: " << str << std::endl;
    std::cout << "value: " << value.myint << endl;
  }
  
  mydata &value;
};

int main(){
  mydata md;
  rule<> rl = str_p("t_need") >> *(("(" >> int_p[myop(md)] >> ")")[myfunc(md)]);
  parse (setup.c_str(), rl, space_p).full;
  return 0;
}
