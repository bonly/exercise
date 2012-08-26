/**
 * @file 20100526_thread.cpp
 * @brief
 *
 * @author bonly
 * @date 2012-7-11 bonly created
 */

#include <boost/thread.hpp>
#include <boost/bind.hpp>
#include <iostream>

void fun()
{
    std::clog << "hello " << std::endl;
}

int main()
{
    boost::thread th(fun);
    th.join();
    return 0;
}


