//  Boost general library logging_test_hello_world.cpp header file  ----------//

//  (C) Copyright Jean-Daniel Michaud 2007. Permission to copy, use, modify, 
//  sell and distribute this software is granted provided this copyright notice 
//  appears in all copies. This software is provided "as is" without express or 
//  implied warranty, and with no claim as to its suitability for any purpose.

//  See http://www.boost.org/LICENSE_1_0.txt for licensing.
//  See http://code.google.com/p/loglite/ for library home page.

#include <iostream>
#include <fstream>
#include <logging.hpp>

int main()
{
  BOOST_LOG_INIT( (boost::logging::trace >> boost::logging::eol) ); //log format
   
  boost::logging::sink sink_cout(&std::cout);
  boost::logging::sink sink_file(new std::ofstream("./output.log"));
  sink_cout.attach_qualifier(boost::logging::log);
  sink_file.attach_qualifier(boost::logging::log);
  
  BOOST_LOG_ADD_OUTPUT_STREAM(sink_cout);
  BOOST_LOG_ADD_OUTPUT_STREAM(sink_file);

  BOOST_LOG_(1, "Hello World!");
  return 0;
}
