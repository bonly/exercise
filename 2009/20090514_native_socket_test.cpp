#define _WIN32_WINNT 0x0501
#define __USE_W32_SOCKETS

#if !defined(BOOST_ALL_NO_LIB)
#define BOOST_ALL_NO_LIB 1
#endif // !defined(BOOST_ALL_NO_LIB)

#define BOOST_TEST_MODULE native_socket_test
#include <boost/test/included/unit_test.hpp>

#include <boost/asio.hpp>

BOOST_AUTO_TEST_CASE (nati)
{
  int native_socket1 = ::socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP);
  boost::asio::io_service ios;
  boost::asio::ip::udp::socket socket2 (ios, boost::asio::ip::udp::v4(), native_socket1);
  BOOST_CHECK (native_socket1 == socket2.native());

}

BOOST_AUTO_TEST_CASE (nati_tcp)
{
  int native_socket1 = ::socket(AF_INET, SOCK_DGRAM, IPPROTO_UDP);
  boost::asio::io_service ios;
  boost::asio::ip::tcp::socket socket2 (ios, boost::asio::ip::tcp::v4(), native_socket1);
  BOOST_CHECK (native_socket1 == socket2.native());

}

/*
aCC -AA +DD64 utest_native_socket.cpp -L ~/boost_1_37_0/stage/lib/ -l boost_system-mt-1_37  -o nati
hejb@ocstest2:~/try$ ./nati --report_level=detailed
Running 2 test cases...

Test suite "native_socket_test" passed with:
  2 assertions out of 2 passed
  2 test cases out of 2 passed

  Test case "nati" passed with:
    1 assertion out of 1 passed

  Test case "nati_tcp" passed with:
    1 assertion out of 1 passed
*/

