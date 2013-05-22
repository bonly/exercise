/*用类成员函数代替静态或全局函数
 * 注意利用第一个void*参数
 *
 */
#include <boost/bind.hpp>
#include <stdio.h>
#include <stdlib.h>
#include <iostream>
#include "sqlite3.h"
using namespace boost;
using namespace std;
static int callback(void *NotUsed, int argc, char **argv, char **azColName){
  NotUsed=0;
  int i;
  for(i=0; i<argc; i++){
    printf("%s = %s\n", azColName[i], argv[i] ? argv[i]: "NULL");
  }
  printf("\n");
  return 0;
}
class Hack
{
	private:
		int x;
	public:
		int callback(void *NotUsed, int argc, char **argv, char **azColName)
		{
		  NotUsed=0;
		  int i;
		  for(i=0; i<argc; i++){
		    printf("%s = %s\n", azColName[i], argv[i] ? argv[i]: "NULL");
		  }
		  printf("\n");
		  return 0;
		}
};
struct A
{
	Hack* p;
	int(Hack::*pmf)(void*,int,char**,char**);
};

int inter(void* ptr,int argc,char** argv,char**az)
{
	A* pa=static_cast<A*>(ptr);
	Hack* ph=pa->p;
	int(Hack::*pmf)(void*,int,char**,char**)=pa->pmf;
	return (ph->*pmf)(ptr,argc,argv,az);
}

int use_org(sqlite3 *db,char* sql)
{
	const char*   tail;
	sqlite3_stmt  *stmt;
	int           error;
	error = sqlite3_prepare(db,sql,256,&stmt,&tail);
	if ( error != SQLITE_OK )
	{
			cerr        << "sqlite_compile error: "
									<< sqlite3_errmsg(db)
									<< " on query: "
									<< sql << endl;
			return -1;
	}
  int cols = sqlite3_column_count(stmt);
  while(true)
  {
  	error = sqlite3_step(stmt);
  	if (error==SQLITE_DONE||error==SQLITE_ERROR)
  		break;
  	for (int i=0; i<cols; ++i)
  	{
  		cout << sqlite3_column_text(stmt,i);
  	}
  }
  error = sqlite3_finalize(stmt);
  if (error!=SQLITE_OK)
  {
  	cerr << "sqlite_finalize error: "
  	     << sqlite3_errmsg(db)
  	     << " on query: "
  	     << sql << endl;
  	return -1;
  }
  return 0;
}
int main(int argc, char **argv)
{
  sqlite3 *db;
  char *zErrMsg = 0;
  int rc;
  if( argc!=3 ){
    fprintf(stderr, "Usage: %s DATABASE SQL-STATEMENT\n", argv[0]);
    exit(1);
  }
  rc = sqlite3_open(argv[1], &db);
  if( rc ){
    fprintf(stderr, "Can't open database: %s\n", sqlite3_errmsg(db));
    sqlite3_close(db);
    exit(1);
  }

  A a;
  Hack h;
  a.p=&h;
  a.pmf=&Hack::callback;
  rc = sqlite3_exec(db, argv[2], inter, &a, &zErrMsg);
  rc = sqlite3_exec(db, argv[2], callback, 0, &zErrMsg);
  if( rc!=SQLITE_OK ){
    fprintf(stderr, "SQL error: %s\n", zErrMsg);
  }

  use_org(db,argv[2]);

  sqlite3_close(db);
  return 0;
}

