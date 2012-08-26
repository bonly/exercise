//============================================================================
// Name        : mysql.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
#include <cstdio>
#include <iostream>
#include <mysql++.h>
using namespace std;

char const * db = "mysql";
char const * server = "localhost";
char const * user = "root";
char const * pass = "";

int simple_select ()
{
   mysqlpp::Connection conn(false);
   if (conn.connect(db, server, user, pass))
   {
      mysqlpp::Query query = conn.query("select cl_1,cl_2,cl_7 from testdb.testmk where cl_1='Linux'");
      //query.store();

      if (mysqlpp::StoreQueryResult res = query.store())
      {
         //cout << "records: " << endl;
         //for (size_t i =0; i < res.num_rows(); ++i)
         //   cout << '\t' << res[i][0] << '\t' << res[i][1] << endl;
      }
      else
      {
         cerr << "Failed to get item list: " << query.error() << endl;
         return 1;
      }

      cout << "finish\n";
      return 0;
   }
   else
   {
      cerr << "DB connection failed: " << conn.error() << endl;
      return 1;
   }
   return 0;
}

int select_from_merge ()
{
   mysqlpp::Connection conn(false);
   if (conn.connect(db, server, user, pass))
   {
      mysqlpp::Query query = conn.query("select cl_1,cl_2,cl_7 from testdb.testmg where cl_1='Linux'");
      //query.store();

      if (mysqlpp::StoreQueryResult res = query.store())
      {
         //cout << "records: " << endl;
         //for (size_t i =0; i < res.num_rows(); ++i)
         //   cout << '\t' << res[i][0] << '\t' << res[i][1] << endl;
      }
      else
      {
         cerr << "Failed to get item list: " << query.error() << endl;
         return 1;
      }

      cout << "from merge finish\n";
      return 0;
   }
   else
   {
      cerr << "DB connection failed: " << conn.error() << endl;
      return 1;
   }
   return 0;
}

int get_rand()
{
   /*
   for (int i = 0 ; i<=100 ; ++i)
   {
      cerr << rand()%500000 << endl;
   }
   */
   return rand()%500000;
}

int select_from_merge_by_ind ()
{
   mysqlpp::Connection conn(false);
   if (conn.connect(db, server, user, pass))
   {
      char sql[255]="";
      for (int i = 0; i<500000; ++i)
      {
         sprintf(sql,"select * from testdb.testmg where id=%d",get_rand());
         mysqlpp::Query query = conn.query(sql);
         query.store();
         /*
         if (mysqlpp::StoreQueryResult res = query.store())
         {
            //cout << "records: " << endl;
            for (size_t i =0; i < res.num_rows(); ++i)
               cout << '\t' << res[i][0] << '\t' << res[i][1] << '\t' << res[i][2] << endl;
         }
         else
         {
            cerr << "Failed to get item list: " << query.error() << endl;
            return 1;
         }
         */
      }

      cout << "from merge finish\n";
      return 0;
   }
   else
   {
      cerr << "DB connection failed: " << conn.error() << endl;
      return 1;
   }
   return 0;
}




int main(int argc, char* argv[])
{
   switch( atoi(argv[1]))
   {
      case 0:
         return simple_select();
      case 1:
         return select_from_merge();
      case 2:
         srand(time(0));
         return select_from_merge_by_ind();
      default:
         break;
   }
   return 0;
}
