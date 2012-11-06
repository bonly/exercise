/**
 * @file 20100702_pyobj.cpp
 *
 * @author bonly
 * @date 2012-10-10 bonly Created
 */

#include<boost/python.hpp>
using namespace boost::python;

/**
def f(x, y):
     if (y == 'foo'):
         x[3:7] = 'bar'
     else:
         x.items += y(3, x)
     return x

def getfunc():
   return f;
 */
object f(object x, object y)
{
    if (y == "foo")
        x.slice(3,7) = "bar";
    else
        x.attr("items") += y(3,x);
    return x;
}

object getfunc()
{
    return object(f);
}
int main()
{
    getfunc();
    return 0;
}



