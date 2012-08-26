//============================================================================
// Name        : io_srv.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
//使用了io_service.run()在cygwin下无法收到信号，但linux下可正常收到信号
#define __USE_W32_SOCKETS
//lib,mswsock,ws2_32
#include <boost/asio.hpp>
#include <boost/date_time/posix_time/posix_time.hpp>
#include <iostream>
#include <csignal>
using namespace boost;
using namespace boost::asio;
using namespace std;
void stop(int signum,siginfo_t *info,void*)
{
	cerr << "recv a signal \n";
}

void reg_sig()
{
  struct sigaction act;
  sigemptyset(&act.sa_mask);
  act.sa_sigaction=stop;
  act.sa_flags=SA_SIGINFO;
  if(sigaction(SIGUSR2,&act,NULL)<0)
  	cerr << "registry err\n";
}

void print(const boost::system::error_code& /*e*/)
{
  std::cout << "Hello, world!\n";
}

int main()
{
	reg_sig();
	io_service io;
	io_service::work worker(io);
	deadline_timer t(io, posix_time::seconds(5));
	t.async_wait(print);
	io.run ();///有work时一直运行，不会返回
    /// 可以用 work保证run()不退出,然后只要有注册事件响应,程序就可以继续触发式工作
	return 0;
}

/*
  work是一个很小的辅助类，只支持构造函数和析构函数。（还有一个get_io_service返回所关联的io_service） 
  构造一个work时，outstanding_work_+1，使得io.run在完成所有异步消息后判断outstanding_work_时不会为0，因此会继续调用GetQueuedCompletionStatus并阻塞在这个函数上。
  而析构函数中将其-1，并判断其是否为0，如果是，则post退出消息给GetQueuedCompletionStatus让其退出。

  因此work如果析构，则io.run会在处理完所有消息之后正常退出。work如果不析构，则io.run会一直运行不退出。如果用户直接调用io.stop，则会让io.run立刻退出。

  特别注意的是，work提供了一个拷贝构造函数，因此可以直接在任意地方使用。对于一个io_service来说，有多少个work实例关 联，则outstanding_work_就+1了多少次，只有关联到同一个io_service的work全被析构之后，io.run才会在所有消息处 理结束之后正常退出。
 */

/*
g++ 20100108_asio_work.cpp -lboost_system
*/

