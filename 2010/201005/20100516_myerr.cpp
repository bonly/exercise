/**
 *  @file 20100516_myerr.cpp
 *
 *  @date 2012-6-30
 *  @author Bonly
 *  @par 希可以用宏定义来减少错误类型的定义,实际试验失败
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

class myerr: public boost::system::error_category
{
        virtual const char* name() const
        {
            return "Bonly's Error";
        }
        virtual std::string message(int e) const
        {
        ///#C的值没转成功? @todo #M只能是一个词,不能是多个词组成的串
#define ERR(E,C,M) \
                if (e == (int)#C) return #M;
#include "20100516_errcode.h"
                return "Unknown error";
#undef ERR
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
    //test(atoi(argv[1]),ec);
    test(0,ec);
    std::clog << ec.message();

    return 0;
}




