#include <boost/dynamic_bitset.hpp>
#include "20111013_mysql_boost.h"
#include <string>
using namespace boost;
using namespace std;

#ifdef __cplusplus
extern "C"{
#endif

my_bool myboost_init(UDF_INIT *initid, UDF_ARGS *args, char *message) {
  return 0;
}

long long myboost(UDF_INIT *initid, UDF_ARGS *args, char *is_null, char *error){
    dynamic_bitset<>  db(std::string(args->args[0]));
    return db.to_ulong();
}

#ifdef __cplusplus
}
#endif

/*
g++  20111013_mysql_boost.cpp -I ~/mysql/include/mysql/ -fPIC -shared -o libmyboost.so
*/
