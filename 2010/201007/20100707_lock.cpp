#include <boost/thread.hpp>
#include <boost/shared_ptr.hpp>
#include <boost/make_shared.hpp>
#include <iostream>

struct Obj
{
   boost::shared_mutex _mutex;
   int num;
};
typedef boost::shared_ptr<Obj> obj_ptr;

Obj p_obj;
struct Want
{
   boost::unique_lock<boost::shared_mutex> lock;
   int myop;
};

void fun1()
{
   while(true)
   {
      //std::clog << "running " << std::endl;
      Want wt;
      wt.lock = boost::unique_lock<boost::shared_mutex>(p_obj._mutex,boost::defer_lock);
      if(wt.lock.try_lock() == true)
      {
             std::clog << "lock: " << boost::this_thread::get_id() << std::endl;
             boost::this_thread::sleep(boost::posix_time::milliseconds(5));
      }
      else
      {
             std::clog << "yield: " << boost::this_thread::get_id() << std::endl;
             boost::this_thread::yield();
      }
      if(wt.lock.owns_lock())
      {
             std::clog << "unlock: " << boost::this_thread::get_id() << std::endl;
             wt.lock.unlock();
      }
   }
}

int main()
{
   p_obj.num = 0;
   boost::thread_group thg;
   thg.create_thread(&fun1);
   thg.create_thread(&fun1);

   thg.join_all();
   return 0;
}

