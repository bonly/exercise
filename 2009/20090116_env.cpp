//============================================================================
// Name        : try_process.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
#include <boost/format.hpp>
using namespace std;
using namespace boost;

int main(int argc, char* argv[])
{
  int i;
  cout << "Argument: \n";
  format k("Arg%d is: %s\n");
  for(i=0;i<=argc;++i)
  	cout << k%i%argv[i];
  printf("Evvironment: \n");
  k=format("environ%d: %s\n");
  for (i=0;environ[i]!=NULL;++i)
  	cout << k%i%environ[i];
  return 0;
}


