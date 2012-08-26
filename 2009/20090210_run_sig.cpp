//============================================================================
// Name        : io_srv.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
//ʹ����io_service.run()��cygwin���޷��յ��źţ���linux�¿������յ��ź�
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
	io.run ();//��workʱһֱ���У����᷵��
        //io.run_one (); //��workҲʱֻ����һ�Σ�����
	//do
	//{
	  //io.poll();//������׼���õ�
	  //io.poll_one();//ֻ����һ��׼���õ�
	//}while(true);
	return 0;
}



