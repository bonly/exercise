#ifndef __CRONT_H__
#define __CRONT_H__
#include <my_global.h>
#include <mysql.h>

#ifdef __cplusplus
extern "C"{
#endif
my_bool cront_init(UDF_INIT *initid, UDF_ARGS *args, char *message);

long long cront(UDF_INIT *initid, UDF_ARGS *args, 
               char *is_null, char *error);

#ifdef __cplusplus
}
#endif
#endif
