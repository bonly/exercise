#include <boost/asio.hpp>
#include <boost/thread.hpp>
using namespace boost;



int
main()
{
  asio::io_service io;
  thread_group tg;
  asio::deadline_timer t(io, boost::posix_time::seconds(5));

  tg.create_thread(bind(&asio::io_service::run,&io));
  io.run();
  tg.join_all();
  return 0;
}
