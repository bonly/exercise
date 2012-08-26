/*
 * t_spirit.cpp
 *
 *  Created on: 2009-6-9
 *      Author: Bonly
 */
#include <iostream>
#include <string>
#include <vector>
#include <fstream>
#include <algorithm>
#include <boost/foreach.hpp>

using namespace std;

int
main()
{
	string sp("短信:100条;");

	char ty[20];
	bzero(ty,20);

	char num[20];
	bzero(num,20);

	sscanf (sp.c_str(),"%[^:]:%[^;];",ty,num);

	char buf[255];
  sprintf(buf,"尊敬的用户，"
              "月结赠送您%s的%s，"
              "%s年%s月%s日生效，"
              "谢谢您的使用！",
              num,ty,"2009","10","20"
              );

  cout << buf << endl;
	return 0;
}
