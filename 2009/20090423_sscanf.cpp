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
	string sp("����:100��;");

	char ty[20];
	bzero(ty,20);

	char num[20];
	bzero(num,20);

	sscanf (sp.c_str(),"%[^:]:%[^;];",ty,num);

	char buf[255];
  sprintf(buf,"�𾴵��û���"
              "�½�������%s��%s��"
              "%s��%s��%s����Ч��"
              "лл����ʹ�ã�",
              num,ty,"2009","10","20"
              );

  cout << buf << endl;
	return 0;
}
