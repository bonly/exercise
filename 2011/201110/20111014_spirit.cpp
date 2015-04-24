//#include <boost/spirit.hpp> //过时的
#include <boost/spirit/include/classic.hpp>
#include <boost/bind.hpp>
#include <string>
#include <iostream>
using namespace std;
using namespace BOOST_SPIRIT_CLASSIC_NS;
using namespace boost;

void func(double n){
  clog << "get a number: " << n << endl;
}

int main(){
    string str("1.44,3.45");
    rule<> r = real_p[&func] >> *(ch_p(',') >> real_p[boost::bind(&func,_1)]);  //不需要ch_p也行,因为幕后隐式生成的chlit对象
    parse (str.c_str(), r, space_p).full;
     
    //parse(str.c_str(), real_p[&func] >> *(',' >> real_p[boost::bind(&func,_1)]), space_p).full;
    return 0;
}
