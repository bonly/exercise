//============================================================================
// Name        : grep_sql.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
#include <boost/filesystem.hpp>
#include <boost/format.hpp>
#include <iostream>
using namespace boost;
using namespace std;
namespace fs = boost::filesystem;

int grep(fs::path fullpath)
{
  if(fs::is_directory(fullpath))
  {
    fs::directory_iterator end_iter;
  }
}

int main(int argc, char **argv)
{
  fs::path fullpath = fs::system_complete(fs::path(argv[1]));
  clog << "full path is: " << fullpath.directory_string() << endl;

  if(!fs::exists(fullpath))
  {
    cerr <<  format("%s does not exists!\n") % fullpath.directory_string();
    return -1;
  }
  return 0;
}
