#include <iostream>
#include <boost/log/core.hpp>
#include <boost/log/trivial.hpp>
#include <boost/log/expressions.hpp>
namespace logging = boost::log;
using namespace std;
void SetFilter() {
  logging::core::get()->set_filter(logging::trivial::severity >= logging::trivial::info);
}
int main () {
  cout << "hello, world" << endl;
  SetFilter();
  BOOST_LOG_TRIVIAL(trace) << "A trace severity message";
  BOOST_LOG_TRIVIAL(debug) << "A debug severity message";
  BOOST_LOG_TRIVIAL(info) << "An informational severity message";
  BOOST_LOG_TRIVIAL(warning) << "A warning severity message";
  BOOST_LOG_TRIVIAL(error) << "An error severity message";
  BOOST_LOG_TRIVIAL(fatal) << "A fatal severity message";
}

//g++ 20110725_log.cpp -lboost_log -lboost_log_setup  -lboost_thread -lpthread -lboost_system -lboost_filesystem -DBOOST_LOG_DYN_LINK
