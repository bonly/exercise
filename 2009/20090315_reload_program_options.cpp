//============================================================================
// Name        : program_op.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
//只能用NEW对象的方式重载配置
//当配置中出程序没有的项目会异常
#include <boost/program_options.hpp>
#include <boost/format.hpp>
#include <boost/shared_ptr.hpp>
#include <fstream>
#include <iostream>
using namespace std;
using namespace boost;
namespace po = boost::program_options;
int main()
{
 try{
  po::options_description desc("Opt");
  desc.add_options()
    ("my_int,i",po::value<int>(),"my int");

  po::variables_map vm;
  shared_ptr<po::variables_map> _vm(new po::variables_map);

  ifstream ifs("intcfg.txt");
//  store(parse_config_file(ifs,desc),vm);
//  ifs.close();
//  notify(vm);

//  ifs.open("intcfg.txt");
  store(parse_config_file(ifs,desc),*_vm);
  ifs.close();
  notify(*_vm);

  //cerr << format("before load is: %d\n")%vm["my_int"].as<int>();
  //取值出来时失败
  cerr << format("before load _vm is: %d\n")%(((*_vm)["my_int"]).as<int>());

  //vm->clear();不能用clear,否则无论如何,后面的打印都异常
  _vm.reset();
  _vm = shared_ptr<po::variables_map>(new po::variables_map);
  ifstream ifs2("intcfg.txt");
  //store(parse_config_file(ifs2,desc),vm);
  store(parse_config_file(ifs2,desc),*_vm);
  ifs2.close();
  //notify(vm);
  notify(*_vm);

  //cerr<< format("after load is: %d\n")%vm["my_int"].as<int>();
  cerr<< format("after load is: %d\n")%(*_vm)["my_int"].as<int>();
 }
 catch(std::exception &e)
 {
	cerr << e.what();
 }
 return 0;
}

