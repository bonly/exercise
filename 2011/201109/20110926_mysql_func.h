#ifndef MTEST_H_
#define MTEST_H_
#include <my_global.h>
#include <mysql.h>

my_bool mtest_init(UDF_INIT *initid, UDF_ARGS *args, char *message);

char *mtest(UDF_INIT *initid, UDF_ARGS *args,
          char *result, unsigned long *length,
          char *is_null, char *error);

void mtest_deinit(UDF_INIT *initid);

#endif
