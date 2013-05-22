#include <iostream>
#include <boost/xpressive/xpressive.hpp>
using namespace boost::xpressive;
int main()
{
    std::string hello( "hello world!" );
    sregex rex = sregex::compile( "(\\w+) (\\w+)!" );
    smatch what;
    if( regex_match( hello, what, rex ) )
    {
        std::cout << what[0] << '\n'; // whole match 完整的匹配
        std::cout << what[1] << '\n'; // first capture 第一项匹配 
        std::cout << what[2] << '\n'; // second capture 第二项匹配
    }
    return 0;
}

