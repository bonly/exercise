#include <string.h>
#include "20110928_tick.hpp"
#include "20110927_mysql_cront.h"

#ifdef __cplusplus
extern "C"{
#endif

my_bool cront_init(UDF_INIT *initid, UDF_ARGS *args, char *message) {
  return 0;
}

long long cront(UDF_INIT *initid, UDF_ARGS *args, 
               char *is_null, char *error){
   bus::Tick tk;
   tk.Parse(args->args[0]);
   return tk.TestNow();
}

#ifdef __cplusplus
}
#endif

/*
g++  20110927_mysql_cront.cpp -L. -ltick -I ~/mysql/include/mysql/ -fPIC -shared -o libbonly.so
*/
