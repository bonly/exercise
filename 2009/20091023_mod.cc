/*
 * main.hpp
 *
 *  Created on: 2010-9-6
 *      Author: bonly
 */

#ifndef MAIN_HPP_
#define MAIN_HPP_
struct Module
{
    int version;
    char mod_name[20];
    int (*init_mod)(int*);
    int (*run_mod)(int*);
    int (*end_mod)(int*);
};


//可使用configure生成的头文件中定义要包括的模块
enum NUM {Min=-1,MyMod,Max};
extern Module myMod;
Module* Mod[]=
{
  &myMod
};

#endif /* MAIN_HPP_ */

/*
 * pMain.cc
 *
 *  Created on: 2010-9-9
 *      Author: bonly
 */

//模块定义文件
#include "main.hpp"
#include <iostream>
using namespace std;


int myMod_init(int *i)
{
  clog << "my mod init\n";
  return 0;
}

Module myMod=
{
   1,
   "mytest",
   &myMod_init,
   NULL,
   NULL
};

//主程序文件
int main()
{
  for (int i=Min+1; i<Max; ++i)
  {
    Mod[i]->init_mod(&i);
  }
  return 0;
}
