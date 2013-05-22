#include <boost/asio.hpp>
using namespace boost::asio;

int main(int argc, char **argv)
{
   io_service io;
   io.run();
   return 0;
}

