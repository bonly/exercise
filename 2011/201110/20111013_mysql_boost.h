#ifndef __MY_H_
#define __MY_H_

#include <my_global.h>  //这两个文件不需要在C中?
#include <mysql.h>

#ifdef __cplusplus
extern "C"{
#endif

my_bool myboost_init(UDF_INIT *initid, UDF_ARGS *args, char *message);

long long myboost(UDF_INIT *initid, UDF_ARGS *args, char *is_null, char *error);

#ifdef __cplusplus
}
#endif
#endif
