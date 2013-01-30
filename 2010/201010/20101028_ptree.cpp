#include <boost/property_tree/ptree.hpp>
#include <boost/property_tree/json_parser.hpp>
#include <iostream>
#include <string>

int makeson(){
  namespace js=boost::property_tree;
  js::ptree pt;
  pt.put("upid","1001");

  std::ostringstream ss;
  js::write_json(ss, pt);

  std::clog << ss.str() << std::endl;
}

int other_json(){
    namespace js=boost::property_tree;
    js::ptree pt_1,pt_11,pt_12;

    pt_11.put("id","3445");
    pt_11.put<int>("age",29);
    pt_11.put("name","chen");    

    pt_12.push_back(make_pair("",pt_11));
    pt_12.push_back(make_pair("",pt_11));

    //replace or create child node "data"
    pt_1.put_child("data",pt_12);

    std::ostringstream os;
    write_json(os,pt_1);
    std::cout << os.str() << std::endl;
}

int main(){
  makeson();
  //other_json();
  return 0;
}

