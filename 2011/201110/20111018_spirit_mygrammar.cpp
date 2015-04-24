/*
不通过编译
*/
//#define BOOST_SPIRIT_THREADSAFE
#include <boost/spirit/include/classic.hpp>
#include <boost/bind.hpp>
#include <string>
#include <iostream>
using namespace std;
using namespace BOOST_SPIRIT_CLASSIC_NS;
using namespace boost;

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

struct skill : public grammar<skill>{
  skill(mydata &val):value(val){
  }
  
  template<typename ScannerT>
  struct definition{
    definition(skill const& self) { //在此定义语法
      team_need =   str_p("t_need") >> ("(" >> int_p[myop(self.value)] >> ")")[myfunc(self.value)];
      attack_need = str_p("a_need") >> ("(" >> int_p[myop(self.value)] >> ")")[myfunc(self.value)];
      defend_need = str_p("d_need") >> ("(" >> int_p[myop(self.value)] >> ")")[myfunc(self.value)];
      an_condition = team_need | attack_need | defend_need;
      condition_list = an_condition >> *(an_condition);
    }
    
    rule<ScannerT> team_need, attack_need, defend_need, an_condition, condition_list;
    
    rule<ScannerT> const& start() const {
      return condition_list;
    }
  };
  
  mydata &value;
};
    
string str_test = "t_need(2011)t_need(3011)";

int main(){
  mydata myd;
  skill en(myd);
  int status = parse(str_test.c_str(), en, space_p).hit;
  cout << status << endl;
  return 0;
}
