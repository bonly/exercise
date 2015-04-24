//#define BOOST_SPIRIT_THREADSAFE
#include <boost/spirit/include/classic.hpp>
#include <boost/bind.hpp>
#include <string>
#include <iostream>
using namespace std;
using namespace BOOST_SPIRIT_CLASSIC_NS;
using namespace boost;

string setup("t2011|d3011&a1400");

struct mydata{
  int myint;
};

struct myop{
  myop(mydata &val):value(val){
  }
  void operator() (int val) const{
    //clog << "chg val: " << val << endl;
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
  rule<> r_single = (str_p("t") >> int_p[myop(md)])[myfunc(md)] | 
                    (str_p("a") >> int_p[myop(md)])[myfunc(md)] |
                    (str_p("d") >> int_p[myop(md)])[myfunc(md)] ;
  rule<> r_op = ch_p('|')|ch_p('&');                   
  rule<> r_ajust =  r_single >> *(r_op >> r_single) ;
  int status = parse (setup.c_str(), r_ajust, space_p).full;
  clog << "status: " << status << endl;
  return 0;
}

/*
g++ 20111016_spirit.cpp  -l pthread -l boost_thread -l boost_system
*/
