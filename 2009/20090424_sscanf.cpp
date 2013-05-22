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
	string sp("短信:100条;免费时长:300秒;");

	string present;

	string cnt(sp);
	int pot = 0;
	while (-1 != (pot = cnt.find_first_of(';',0)))
	{
		char ty[20];
		bzero(ty,20);

		char num[20];
		bzero(num,20);

		string msg(cnt,0,pot);
		sscanf (cnt.c_str(),"%[^:]:%[^;];",ty,num);
		cnt.erase(0,pot+1);

		present.append(num);
		present.append("的");
		present.append(ty);
		present.append("，");
	}

	char buf[255];
  sprintf(buf,"尊敬的用户，"
              "月结赠送您%s"
              "%s年%s月%s日生效，"
              "谢谢您的使用！",
              present.c_str(),"2009","10","20"
              );

  cout << buf << endl;
	return 0;
}

