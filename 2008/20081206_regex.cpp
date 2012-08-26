#define BOOST_REGEX_NO_LIB 
#include<iostream>
#include<boost/regex.hpp>
#include<string>
int main()
{
      boost::regex reg4("[^13579]");
      std::string s="0123456789";
      boost::sregex_iterator it(s.begin(),s.end(),reg4);
      boost::sregex_iterator end;
      while (it!=end)
     std::cout << *it++;
}
