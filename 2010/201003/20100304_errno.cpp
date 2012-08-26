#include <iostream>
using namespace std;

#ifndef __ERROR_CODE__
#define __ERROR_CODE__

template<typename Data>
struct Code
{
    const char* what(int code)
    {
        return data->msg[code];
    }
    const char* what()
    {
        return data->msg[this->code];
    }
    int code;
    Data* data;
};

struct Error : public Code<Error>
{
    const char* msg[];
    Error()
    {
        data = this;
    }
    Error* operator()(int cd)
    {
        code = cd;
        return this;
    }
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
