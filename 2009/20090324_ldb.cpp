/*
 * SQL_DB.h
 *
 *  Created on: 2009-4-18
 *      Author: Bonly
 */

#ifndef __SQL_DB_HPP__
#define __SQL_DB_HPP__
#include <boost/shared_ptr.hpp>
#include <iostream>
#include "sqlite3.h"
namespace bonly
{
using namespace std;
using namespace boost;

class DB;
class Command
{
	public:
		friend  class DB;
		//void        db(DB* d){_db=d;}
		int           Perform();
		//int         Close();
		int           Drop();
		sqlite3_stmt* stmt()
		{return _stmt;}

		int           setParam(int i, int val);
		int           setParam(int, const char*, int n=-1, void(*p)(void*)=SQLITE_STATIC);

    void          getColumn(int i, int* var);
    void          getColumn(int i, const unsigned char* *var);

	private:
		sqlite3_stmt*  _stmt;
		DB*            _db;
};

typedef shared_ptr<Command> PCmd;

class DB
{
	public:
		DB();
		DB(const char* dsn);
		virtual ~DB();
		int      Connect(const char* dsn);
		int      Disconnect();
		PCmd     Prepare(const char* sql);
		//int    Commit();
		//int    Rollback();
		void     errmsg(const char* where);

	private:
		sqlite3*      _db;
};
}
#endif /* __SQL_DB_HPP__ */

/*
 * SQL_DB.cpp
 *
 *  Created on: 2009-4-18
 *      Author: Bonly
 */

#include "ldb.hpp"

namespace bonly
{
int
Command::Perform()
{
	int ret = sqlite3_step (_stmt);
	if (!(ret == SQLITE_ROW
			||ret == SQLITE_DONE
      ||ret == SQLITE_OK))
	{
		_db->errmsg("Command::Perform");
	}
  return ret;
}

int
Command::Drop()
{
	int ret = sqlite3_finalize(_stmt);
	if (ret != SQLITE_OK)
	{
		_db->errmsg("Command::Drop");
	}
	return ret;
}

int  Command::setParam(int i, int val)
{
	int ret = sqlite3_bind_int (_stmt,i,val);
	if (SQLITE_OK != ret)
	{
		_db->errmsg("Command::setParam");
	}
	return ret;
}

int  Command::setParam(int i, const char* val, int n, void(*p)(void*))
{
	int ret = sqlite3_bind_text (_stmt,i,val,n,p);
	if (SQLITE_OK != ret)
	{
		_db->errmsg("Command::setParam");
	}
	return ret;
}

void Command::getColumn(int i, int* var)
{
	*var = sqlite3_column_int(_stmt, i);
}

void Command::getColumn(int i, const unsigned char* *var)
{
  *var = sqlite3_column_text(_stmt, i);
}

/*
 * *******************************************************************
 */
DB::DB(){}
DB::DB(const char* dsn)
{	Connect (dsn);}
DB::~DB(){}

int
DB::Connect (const char* dsn)
{
	int ret = sqlite3_open (dsn, &_db);
	if (ret != SQLITE_OK)
	{
		errmsg("DB::Connect");
	}
	return ret;
}

int
DB::Disconnect()
{
	sqlite3_stmt *pStmt;
	while( (pStmt = sqlite3_next_stmt(_db, 0))!=0 )
	{
	    sqlite3_finalize(pStmt);
	}

	int ret = sqlite3_close (_db);
	if (ret != SQLITE_OK)
	{
		errmsg("DB::Disconnect");
	}
	return ret;
}

PCmd
DB::Prepare (const char* sql)
{
	const char* tail;
	PCmd cmd (new Command);
	cmd->_db=this;
	//param[3]=-1: 根据sql的第一个\000截止
	int ret = sqlite3_prepare(_db,sql,-1,&(cmd->_stmt),&tail);
	if (ret != SQLITE_OK)
	{
     errmsg("DB::Prepare");
	}
	return cmd;
}

void
DB::errmsg(const char* where)
{
	printf("When %s\n",where);
	printf("error code: %d\n"
			   "ext   code: %d\n"
			   "       msg: %s\n",
			    sqlite3_errcode(_db),
			    sqlite3_extended_errcode(_db),
			    sqlite3_errmsg(_db));
}


}


/*
 * try_sqlite.cpp
 *
 *  Created on: 2009-4-19
 *      Author: Bonly
 */
#include "ldb.hpp"
using namespace bonly;
int
main()
{
  DB db("testa");
  PCmd ta = db.Prepare("select count(*) from sqlite_master where tbl_name='myta'");
  int tbl=-1;
  if (ta->Perform()==SQLITE_ROW)
  	tbl=sqlite3_column_int(ta->stmt(),0);

  if(tbl<=0)
  {
		ta = db.Prepare("create table myta(k int, ca varchar(20))");
		if (ta->Perform()!=SQLITE_DONE)
		{
			db.errmsg("create");
			cerr << "create fail\n";
		}
  }
  ta = db.Prepare("insert into myta values(14,'another tests')");
  if (ta->Perform()!=SQLITE_DONE)
  {
  	db.errmsg("insert");
    cerr << "inser fail\n";
  }

  ta = db.Prepare("select * from myta where k=? and ca=?");
  ta->setParam(1,14);
  ta->setParam(2,"another tests");
  while (ta->Perform()==SQLITE_ROW)
  {
  	int d;
  	ta->getColumn(0, &d);
  	const unsigned char* buf;
  	ta->getColumn(1, &buf);
  	cout << d << "\t" << buf <<endl;
  }
  ta->Drop();
  db.Disconnect();
	return 0;
}

