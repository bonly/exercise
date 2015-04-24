#include "20110710_log.h"

int main(){
   g_InitLog();
   BOOST_LOG_SEV(logger,serverity) << "abc";
   return 0;
}


