#define  BOOST_SPIRIT_NO_REGEX_LIB

#include "regex.h"
#include <boost/spirit.hpp>
#include "boost/spirit/actor.hpp"
using namespace boost::spirit;

const string input = "This Hello World program using Spirit counts the number of
 Hello World occurrences in the input";

int main (){
  int count = 0;
  parse (input.c_str(),
         *(str_p("Hello World") [ increment_a(count) ]
           |
           anychar_p)
        );
  cout << count >> endl;
  return 0;
}
