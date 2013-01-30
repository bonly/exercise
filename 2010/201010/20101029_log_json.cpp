/**
多线程中使用property_tr-mtee 需要
-mt
#define BOOST_SPIRIT_THREADSAFE
*/
#include <boost/property_tree/ptree.hpp>
#include <boost/property_tree/json_parser.hpp>
#include <iostream>
#include <string>

int makeson(){
  namespace js=boost::property_tree;
  js::ptree pt;
  pt.put("ip","192.168.8.1");
  pt.put<int>("port",1234);
  pt.put<int>("player_id", 12346);
  pt.put<int>("cmd_id", 11334);
  pt.put<int>("sub_cmd_id", 123);
  pt.put("last_msg_time","20120104 14:23:23");
  pt.put<int>("ret", 0);
  pt.put<int>("message_count", 3);

  std::ostringstream ss;
  js::write_json(ss, pt);

  std::clog << ss.str() << std::endl;
}

int main(){
  makeson();
  return 0;
}

