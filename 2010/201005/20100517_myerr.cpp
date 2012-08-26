/**
 *  @file 20100517_myerr.cpp
 *
 *  @date 2012-6-30
 *  @author Bonly
 *   @par 可以用宏定义来减少错误类型的定义,但定义的顺序必须注意和msg中的顺序一致，　msg中不能带＂，＂
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

#define ERR(E,C,M) #M,
const char* Msg[]=
{
#include "20100516_errcode.h"
    "Unknown error"
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
            if (e < 0 || e >= ERR_MAX)
                return Msg[ERR_MAX];
            return Msg[e];
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






