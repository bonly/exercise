//============================================================================
// Name        : try_log.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
#include <fstream>
#include "log.hpp"
using namespace std;
using namespace boost;
using namespace bonly;
int main()
{

  BONLY_LOG_INIT( (trace >> eol) ); //log format

  sink sink_cout(&std::cout);
  sink sink_file(new std::ofstream("./output.log"));
  sink_cout.attach_qualifier(bonly::log);
  sink_file.attach_qualifier(bonly::log);

  BONLY_LOG_ADD_OUTPUT_STREAM(sink_cout);
  BONLY_LOG_ADD_OUTPUT_STREAM(sink_file);

  BONLY_LOG_(1, "Hello World!");
  return 0;
}

