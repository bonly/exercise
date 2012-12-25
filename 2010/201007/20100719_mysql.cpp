#include <mysql.h>
#include <iostream>
#include <signal.h>
using namespace std;

struct DB
{
  int port;
  const char* server;
  const char* db;
  const char* user;
  const char* passwd;
  MYSQL connect;
};

DB g_db;

int ping()
{
   if (mysql_ping(&g_db.connect) == 0)
   {
       clog << "ping ok.." << endl;
   }
   else
   {
       clog << "ping error" << endl;
       return -1;
   }
   return 0;
}
int main()
{
   g_db.port = 3306;
   g_db.server = "127.0.0.1";
   g_db.db = "paladin";
   g_db.user = "bonly";
   g_db.passwd = "";
   mysql_init (&g_db.connect);
    if (0 == mysql_real_connect(&(g_db.connect), g_db.server, //连接数据库
                g_db.user, g_db.passwd, g_db.db, g_db.port, NULL, CLIENT_INTERACTIVE|CLIENT_FOUND_ROWS))

    {
        clog << "connect to mysql failed: " << mysql_error(&g_db.connect);
        return  -1;
    }
    bool auto_cnt = true;
    mysql_options(&g_db.connect, MYSQL_OPT_RECONNECT, &auto_cnt);
    int timeout = 1;
    mysql_options(&g_db.connect, MYSQL_OPT_READ_TIMEOUT, &timeout);
 
    clog << "mysql id: " << g_db.connect.thread_id << endl;

    if (0 != mysql_query(&g_db.connect, "set session interactive_timeout=1"))
    {
        clog << "exec error" << endl;
    }
    sigset_t wait_mask;
    sigemptyset(&wait_mask);
    sigaddset(&wait_mask, SIGINT);
    sigaddset(&wait_mask, SIGQUIT);
    sigaddset(&wait_mask, SIGTERM);
    sigaddset(&wait_mask, SIGUSR2);
    sigaddset(&wait_mask, SIGHUP);
    pthread_sigmask(SIG_BLOCK, &wait_mask, 0);
    int sig = 0;
    int ret = -1;
    while (-1 != (ret = sigwait(&wait_mask, &sig)))
    {
        clog << "Receive signal. " << sig << endl;
        if (sig == SIGUSR2)
        {
            clog  << "go to ping" << endl;
            ping();
            clog << "mysql id: " << g_db.connect.thread_id << endl;
            continue;
        }
        if (sig == SIGHUP)
        {
            clog  << "Receive reload config signal." << endl;
            continue;
        }
        if (sig == SIGTERM || sig == SIGQUIT || sig == SIGINT)
        {
            clog << "Receive stop signal, Exit." << endl;
            break;
        }
   }
   return 0;
}
