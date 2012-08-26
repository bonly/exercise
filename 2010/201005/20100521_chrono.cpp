/**
 *  @file 20100521_chrono.cpp
 *
 *  @date 2012-7-9
 *  @author Bonly
 */

#include <boost/chrono.hpp>
#include <cmath>

int main()
{
    boost::chrono::system_clock::time_point start = boost::chrono::system_clock::now();

    for ( long i = 0; i < 10000000; ++i )
    std::sqrt( 123.456L ); // burn some time

    boost::chrono::duration<double> sec = boost::chrono::system_clock::now() - start;
    std::cout << "took " << sec.count() << " seconds\n";
    return 0;
}



