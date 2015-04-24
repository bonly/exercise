//#define BOOST_SPIRIT_THREADSAFE
#include <boost/spirit/include/classic.hpp>
#include <boost/bind.hpp>
#include <string>
#include <iostream>
using namespace std;
using namespace BOOST_SPIRIT_CLASSIC_NS;
using namespace boost;

struct my_grammar : public grammar<my_grammar>{
  /*
派生类必须有一个名为 definition（可以不修改此名）的嵌套的模板类/结构。definition 类有以下特性：
它是类型名为 ScannerT 的模板类。
语法规则在其构造函数中定义。构造函数被作为引用传递给实际的语法 self。
必须提供名为 start 的成员函数，它表示 start 规则。
  */
  template<typename ScannerT>
  struct definition{
    definition(my_grammar const& self) { //在此定义语法
      enum_specifier = enum_p >> '{' >> enum_list >> '}';
      enum_p = str_p("enum");
      enum_list = +id_p >> *(',' >> +id_p);
      id_p = range_p('a', 'z');    
    }
    
    rule<ScannerT> enum_specifier, enum_p, enum_list, id_p;
    rule<ScannerT> const& start() const {
      return enum_specifier;
    }
  };
};
    
    
string str_test = "enum{ah, bk}";

int main(){
  my_grammar en;
  int status = parse(str_test.c_str(), en, space_p).hit;
  cout << status << endl;
  return 0;
}
