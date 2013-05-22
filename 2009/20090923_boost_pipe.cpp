#include <stdio.h>
#include <boost/iostreams/device/file_descriptor.hpp>
#include <boost/iostreams/stream.hpp>
#include <fstream>
#include <iostream>
using namespace boost::iostreams ;
 
struct opipestream : stream< file_descriptor_sink >
{
  typedef stream< file_descriptor_sink > base ;
  explicit opipestream( const char* command )
    : base( fileno( pipe = popen( command, "w" ) ) ) {}
  ~opipestream() { close() ; pclose( pipe ) ; }
  private : FILE* pipe ;
};
 
int main()
{
  std::cout << "#includes in this file:" << std::flush ;
  std::ifstream fin( __FILE__ ) ;
  opipestream pout( "grep '^#include' | wc -l" ) ;
  pout << fin.rdbuf() ;
}
 
// link with libboost_iostreams  eg.
// g++ -Wall -std=c++98 -pedantic -I /usr/local/include -L /usr/local/lib -lboost_iostreams myfile.cc