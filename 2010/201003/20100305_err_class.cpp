/**
 * @file 20100305_err_class.cpp
 * @brief 用类模拟标准errno
 *
 * @author bonly
 * @date 2011-10-13 bonly created
 */

#include <iostream>
using namespace std;

#ifndef __ERROR_CODE__
#define __ERROR_CODE__

struct Error
{
    const char* what(int code)
    {
        return this->msg[code];
    }
    const char* what()
    {
        return this->msg[this->code];
    }

    Error* operator()(int cd)
    {
        code = cd;
        return this;
    }
    int code;
    const char* msg[];
};
Error *error = 0;

#endif //__ERROR_CODE__

struct DB_Code : public Error
{
   enum {sucess=0,faild,conlost,conrej};
   DB_Code()
   {
       msg[sucess]="sucess";
       msg[faild]="faild";
       msg[conlost]="db connect lost";
       msg[conrej]="db connect reject";
   }
}DB_Error;


int gen_err()
{
   error = DB_Error(DB_Code::conrej);
   return -1;
}


int main(int argc, char* argv[])
{
    error = DB_Error(DB_Code::sucess);
    cerr << "errno: " << error->what() << endl;

    cerr <<  "errno: " << DB_Error.what(DB_Code::sucess) << endl;
    cerr <<  "errno: " << DB_Error.what(DB_Code::conlost) << endl;

    if (0 != gen_err())
    {
        cerr <<  "errno: " << error->what() << endl;
    }

    return 0;
}
/**
 * 一个类似系统的errno
 */



