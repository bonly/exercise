/**
 *  @file 20100518_myerr.cpp
 *
 *  @date 2012-6-30
 *  @author Bonly
 *  @par 想通过预定义错误信息来输出错误资料，失败
 */

#include <boost/system/error_code.hpp>
#include <boost/tr1/tr1/string>
#include <iostream>

#define ERROR
#define ERR(E,C,M) E=C,
enum errc
{
#include "20100516_errcode.h"
    ERR_MAX
};
#undef ERR

#define ERR(E,C,M) \
    const char* ER_##C=#M;
#include "20100516_errcode.h"
#undef ERR

#define N2S(N) N
#define getERR(X) ER_##N2S(X)

class myerr: public boost::system::error_category
{
        virtual const char* name() const
        {
            return "Bonly's Error";
        }
        virtual std::string message(int e) const
        {
            if (e < 0 || e >= ERR_MAX)
                return "Unknown error";
            return getERR(e);
        }

    public:

        static boost::system::error_category& myerr_cat()
        {
            static myerr instance;
            return instance;
        }
};
#undef ERROR

void test(int e, boost::system::error_code& ec)
{
    ec = boost::system::error_code(e, myerr::myerr_cat());
}
int main(int argc, char* argv[])
{
    boost::system::error_code ec;
    test(atoi(argv[1]),ec);
    //test(0,ec);
    std::clog << ec.message();

    return 0;
}










