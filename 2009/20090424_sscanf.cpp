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
	string sp("����:100��;���ʱ��:300��;");

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
		present.append("��");
		present.append(ty);
		present.append("��");
	}

	char buf[255];
  sprintf(buf,"�𾴵��û���"
              "�½�������%s"
              "%s��%s��%s����Ч��"
              "лл����ʹ�ã�",
              present.c_str(),"2009","10","20"
              );

  cout << buf << endl;
	return 0;
}

