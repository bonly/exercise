#include "r5log.h"
#include <iostream>
#include <boost/format.hpp>
using namespace boost;
using namespace std;

int
main ()
{
   if (SET_LOG_DIR("te_log")<0)
   {
     cout << format ("log path is not exist!\n");
   }
   SET_LOG_LEVEL(1, 1);
   SET_LOG_NAME_HEAD("INFO");

   R5_DEBUG (("DEBUG:%s" "test\n"));
   R5_DEBUG (("KO:" "file %s" "OK"));

   g_r5_log.flush();
   return 0;

}

