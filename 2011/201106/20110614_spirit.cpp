///定义regex 源代码只作为头文件包含
#define BOOST_SPIRIT_NO_REGEX_LIB

#include <boost/regex.hpp>
#include <boost/spirit.hpp>
#include <boost/spirit/home/classic/actor.hpp>
using namespace boost::spirit;

#include <iostream>
#include <string>

const std::string input = "This Hello World program using Spirit counts the number of Hello World occurrences in the input";

int main (){
  int count = 0;
  parse (input.c_str(),
         *(str_p("Hello World") [ increment_a(count) ]
           |
           anychar_p)
        );
        
  //str_p 和 anychar_p 都是 Spirit 中预定义的解析器 
  //—— str_p 匹配它所提供的字符串（在此为 Hello World）并成功调用 increment_a 例程将计数加 1。
  //anychar_p 是另一个预定义解析器，它可以匹配任何字符。
  std::cout << count << std::endl;
  return 0;
}