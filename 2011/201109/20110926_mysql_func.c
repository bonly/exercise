#include <string.h>
#include "20110926_mysql_func.h"

my_bool mtest_init(UDF_INIT *initid, UDF_ARGS *args, char *message)
{
  if(0 != args->arg_count){
     strncpy(message, "mtest has no arguments", strlen("mtest has no arguments") + 1);
     return 1;
  }

  initid->ptr = calloc(1, 1024);

  return 0;
}

char *mtest(UDF_INIT *initid, UDF_ARGS *args,
          char *result, unsigned long *length,
          char *is_null, char *error)
{
   char *ps = "mysql plugin string type test.";
   *length = strlen(ps);

   memcpy(initid->ptr, ps, *length + 1);

   return initid->ptr;
}

void mtest_deinit(UDF_INIT *initid)
{
   free(initid->ptr);
}

/*
gcc -I ~/mysql/include/mysql 20110926_mysql_func.c -fPIC -shared -o libmtest.so
CREATE FUNCTION mtest RETURNS STRING SONAME 'libmtest.so';
select mtest();
*/
