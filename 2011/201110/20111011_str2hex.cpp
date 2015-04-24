#include <boost/lexical_cast.hpp>
#include <iostream>

using namespace std;
using namespace boost;

int main(){
  try{
     unsigned int x = lexical_cast<unsigned int>("0x0bad");
     cout << x << endl;
  }catch(bad_lexical_cast &){
     clog << "cast error" << endl;
  }
}

/*
'?' 字符ASCII=63
8进制表示  0开头  0123  '\077'
16进制表示 0x(大小写不分)开头 0x123   字符'\0x3f'

8和16进制都只能用于表示无符号正整数,如加了符号,则不被视为8/16进制数
*/

