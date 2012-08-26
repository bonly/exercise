/**
 *  @file 20100515_errcode.cpp
 *
 *  @date 2012-6-30
 *  @author Bonly
 */
#include <boost/system/error_code.hpp>
#include <boost/tr1/tr1/string>
#include <iostream>
enum errc
{
    succ=0,
    faild=1,
    other=3
};
class myerr: public boost::system::error_category
{
        virtual const char* name() const
        {
            return "Bonly's Error";
        }
        virtual std::string message(int e) const
        {
            switch (e)
            {
                case succ:
                    return "success";
                default:
                    return "Unknown error";
            }
        }

    public:

        static boost::system::error_category& myerr_cat()
        {
            static myerr instance;
            return instance;
        }
};

void test(boost::system::error_code& ec)
{
    ec = boost::system::error_code(3, myerr::myerr_cat());
}
int main(int argc, char* argv[])
{
    boost::system::error_code ec;
    test(ec);
    std::clog << ec.message();

    return 0;
}

