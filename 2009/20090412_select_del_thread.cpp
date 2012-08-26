#include "db.hpp"
#include <boost/format.hpp>
#include <boost/thread.hpp>
#include <boost/bind.hpp>
using namespace bonly;
using namespace boost;

template<typename T>void print_var(const char* k, T v)
{ cerr << format("%1% : %2%\n")%k%v;}

#define PRINT_VAR(k) print_var(#k, k)

struct CSMS
{
  int         sms_id;
  char        msg_type[3+1];
  int         send_type;
  int         urgent;
  char        msisdn[20+1];
  char        brand_id[2+1];
  char        f1[20+1];
  char        f2[20+1];
  char        f3[20+1];
  char        f4[20+1];
  char        f5[20+1];
  char        f6[20+1];
  char        f7[20+1];
  char        f8[20+1];
  char        f9[20+1];
  char        f10[1024+1];
  CSMS(){memset(this,0,sizeof(CSMS));}
  void print();
};

void CSMS::print()
{
  PRINT_VAR(sms_id);
  PRINT_VAR(msg_type);
  PRINT_VAR(send_type);
  PRINT_VAR(urgent);
  PRINT_VAR(msisdn);
  PRINT_VAR(brand_id);
  PRINT_VAR(f1);
  PRINT_VAR(f2);
  PRINT_VAR(f3);
  PRINT_VAR(f4);
  PRINT_VAR(f5);
  PRINT_VAR(f6);
  PRINT_VAR(f7);
  PRINT_VAR(f8);
  PRINT_VAR(f9);
  PRINT_VAR(f10); 
}

class Del
{
  public:
    Del()
    {
        del_db.Connect("DSN=bonly");
        del_db.Debug("sql_del.log");
        del = del_db.Prepare(
      " DELETE FROM SMS WHERE sms_id = :1 "
        );
        del_db.Commit();
    }
    int del_sms(int id)
    {
      try
      {
        del->setParam(1, id);
        del->Execute(); 
        del_db.Commit();
      }
      catch(std::exception &e)
      {
        cerr << "del err "<<e.what()<<endl;
        return -1;
      }
      return 0;
    }
    ~Del()
    {
       del->Drop();
       del_db.Disconnect(); 
    }
    
    CDB del_db;
    PCmd del;
};

int 
main ()
{
  CDB  db("DSN=bonly");
  db.Debug("sql.log");
  PCmd get = db.Prepare(
" SELECT rows :1 to :2"
" sms_id,msg_type,send_type,urgent,msisdn,brand_id, "
" f1,f2,f3,f4,f5,f6,f7,f8,f9,f10 "
" FROM SMS ORDER BY sms_id "
  );
  db.Commit();

  Del del;
  thread_group thr;
    
  int from = 1;
  int step = 5;
  while(true)
  {
    int to = from + step;      
    get->setParam(1, from);
    get->setParam(2, to);
    get->Execute();
    int row = 0;
    while(get->FetchNext()==0)
    {
        CSMS sms;
        get->getColumnNullable(1, &sms.sms_id);
        get->getColumnNullable(2,  sms.msg_type);
        get->getColumnNullable(3, &sms.send_type);
        get->getColumnNullable(4, &sms.urgent);
        get->getColumnNullable(5,  sms.msisdn);
        get->getColumnNullable(6,  sms.brand_id);
        get->getColumnNullable(7,  sms.f1);
        get->getColumnNullable(8,  sms.f2);
        get->getColumnNullable(9,  sms.f3);
        get->getColumnNullable(10, sms.f4);
        get->getColumnNullable(11, sms.f5);
        get->getColumnNullable(12, sms.f6);
        get->getColumnNullable(13, sms.f7);
        get->getColumnNullable(14, sms.f8);
        get->getColumnNullable(15, sms.f9);
        get->getColumnNullable(16, sms.f10);    
        cout << "\n\nRecode: \n";
        sms.print();
        
        //使用线程删除,会使select漏了记录
        thr.create_thread(bind(&Del::del_sms,ref(del),sms.sms_id));
        ++row;
    }
    from = from + row;
    
    get->Close();
    if (from < to)
      break;
  }
  get->Drop();
  db.Disconnect();

  thr.join_all();
  return 0;

}

/*
aCC -o sm db.cpp main.cpp -AA +DD64 -DTT_64BIT -I${TT_HOME}/include -L${TT_HOME}/lib -ltten -lttclasses -mt -L/home/hejb/boost_1_37_0/stage/lib -lboost_thread-mt-1_37
*/



