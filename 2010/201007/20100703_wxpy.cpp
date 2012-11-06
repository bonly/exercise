/**
 * @file 20100703_wxpy.cpp
 *
 * @author bonly
 * @date 2012-10-10 bonly Created
 */

#include<boost/python.hpp>
using namespace boost::python;

int main()
{
    class_<foo>("foo");
    return 0;
}


