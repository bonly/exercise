#define _WIN32_WINNT 0x0501
#define BOOST_ALL_NO_LIB
#include <string>
#include <iostream>
#include <boost/date_time/posix_time/posix_time.hpp>
#include <boost/asio.hpp>
using namespace std;
using namespace boost;
using namespace boost::asio;
void print(const boost::system::error_code&)
{
	static int no=0;
	cerr << ++no << "\r"; //\r不起作为
	//printf("%d\r",no++);
}

void test()
{
	io_service io;
	io_service::work worker(io);
	deadline_timer t(io, posix_time::seconds(5));
	t.async_wait(print);
	//io.run ();//有work时一直运行，不会返回
	//io.run_one (); //有work也时只运行一次，返回
	do
	{
		io.poll();//运行已准备好的
		//io.poll_one();//只运行一次准备好的
	}while(true);
	return ;
}
#define BOOST_TEST_MODULE native_socket_test
#include <boost/test/included/unit_test.hpp>


BOOST_AUTO_TEST_CASE (hostname1)
{
	test();
}

/*
 *aCC -AA +DD64 po.cpp -L ~/boost_1_37_0/stage/lib/ -l boost_system-mt-1_37 -o pol
 */

