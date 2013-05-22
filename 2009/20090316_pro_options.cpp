/*
 * profile_config.hpp
 *
 *  Created on: 2009-4-10
 *      Author: Bonly
 */

#ifndef __PROFILE_CONFIG_H__
#define __PROFILE_CONFIG_H__
#include <boost/shared_ptr.hpp>
#include <boost/program_options.hpp>
#include <iosfwd>
using namespace std;
using namespace boost;
namespace op=boost::program_options;

class Config
{
  public:
    int parse(int argc=0, char* argv[]=0);
    static Config& instance();
    int count(const char* key){return (*_cfg).count(key);}
    void help();

    template<typename RET>
    RET get(const char* key){return (*_cfg)[key].as<RET>();}

  private:
  	shared_ptr<program_options::variables_map>         _cfg;
  	program_options::options_description               _desc_cfg;
  	int argc;char** argv;
  	int clear();

  public:
    static Config* _config;
};

#endif //__PROFILE_CONFIG_H__


/*
 * profile_config.cpp
 *
 *  Created on: 2009-4-10
 *      Author: Bonly
 */

#include <fstream>
#include <iostream>
#include "profile_config.hpp"

Config* Config::_config = 0;
void Config::help()
{
	cout << _desc_cfg;
}
int Config::clear()
{
	_cfg.reset();
	_cfg = shared_ptr<program_options::variables_map>(new program_options::variables_map);
	return 0;
}

int Config::parse(int argc, char* argv[])
{
  try
  {
  	clear();
  	if(argc!=0 || argv!=0)
  	{
  	  this->argc=argc;
  	  this->argv=argv;
  	}

    _desc_cfg.add_options ()
      ("help,h", "about this")
      ("version,v", "version ")
      ("config-file,c", program_options::value<std::string>()->default_value("default_config"),"use config file")
      ("ip,s", program_options::value<std::string>()->default_value("127.0.0.1"),"service ip")
      ("port,p", program_options::value<int>()->default_value(12005),"service port")
      ("dsn,t", program_options::value<std::string>()->default_value("DSN=BFS"),"DSN name")
      ("host,h", program_options::value<std::string>()->default_value("real.sms.revenco.com"),"Origin Host")
      ("sleep,e",program_options::value<int>()->default_value(5),"sleep time")
      ("log_path,l",program_options::value<string>()->default_value("./"),"log file path")
      ("file_level,v",program_options::value<int>()->default_value(1),"log file level")
      ("file_term,m",program_options::value<int>()->default_value(1),"log term level")
      ("log_head,g",program_options::value<string>()->default_value("life"),"log file prefix")
       ;

    program_options::positional_options_description p;
    p.add("config-file", -1);
    store (
        program_options::command_line_parser(this->argc,this->argv).options(_desc_cfg).positional(p).run(),
        *_cfg);

    notify(*_cfg);

    if ((*_cfg).count("config-file"))
    {
      cout << "use config file " << (*_cfg)["config-file"].as<std::string>()
           << "\n";
    }

    ifstream ifs((*_cfg)["config-file"].as<std::string>().c_str());
    store(parse_config_file(ifs,_desc_cfg),*_cfg );
    ifs.close();
    notify(*_cfg);
  }
  catch(std::exception& e)
  {
    cout << e.what() << "\n";
    exit(EXIT_FAILURE);
  }
  return 0;
}

Config& Config::instance()
{
  if (Config::_config == 0)
  {
    Config::_config = new Config;
  }
  return *Config::_config;
}


#include "profile_config.hpp"
#include <boost/format.hpp>
#include <iostream>
using namespace boost;
using namespace std;
#ifndef OPENING_TIME
#define OPENING_TIME (__DATE__ __TIME__)
#endif

#ifndef TEST_TIME
#define TEST_TIME (__DATE__ __TIME__)
#endif

#ifndef PROGRAM_VERSION
#define PROGRAM_VERSION  (__DATE__ __TIME__)
#endif

#define PROGRAM_NAME "life"

int
main (int argc, char* argv[])
{
  Config::instance().parse(argc,argv);

  if(Config::instance().count("version"))
  {
  	cout << PROGRAM_VERSION << endl;
  }
  if(Config::instance().count("help"))
  {
  	Config::instance().help();
  }
	return 0;
}

