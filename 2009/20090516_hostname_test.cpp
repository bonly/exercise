/*
 *用boost_test_main就没有改变master cast 的名字,但也是生成了main函数的
 */
#define BOOST_TEST_MAIN native_socket_test
#include <boost/test/included/unit_test.hpp>

//#define _WIN32_WINNT 0x0501
#define __USE_W32_SOCKETS
#include <boost/asio.hpp>
#include <string>
#include <iostream>
using namespace std;
using namespace boost::asio;

BOOST_AUTO_TEST_CASE (hostname1)
{
  string hostname("bonly");
  string hn = boost::asio::ip::host_name();
  BOOST_CHECK(hostname==hn);
}

/*
 * mingw
 */

