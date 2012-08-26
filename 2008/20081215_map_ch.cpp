//============================================================================
// Name        : cr.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
#include <map>
#include <string>
using namespace std;

class CMapRetCode
{
   public:
       typedef std::map<int, std::string>::iterator MapRetCode2StrPtr;
	CMapRetCode()
	{
     MapRetCode2Str.insert(make_pair(2001,"�ɹ�"));
     MapRetCode2Str.insert(make_pair(1001,"�û������� | �ʻ������� | �ʱ�������"));
	   MapRetCode2Str.insert(make_pair(1002,"�û�����"));
	   MapRetCode2Str.insert(make_pair(1003,"ָ���������кŲ�����"));
	   MapRetCode2Str.insert(make_pair(1004,"ָ�����ײ��ɹ�"));
	   MapRetCode2Str.insert(make_pair(1005,"�ع�ʧ��"));
	   MapRetCode2Str.insert(make_pair(1006,"��ֵϵͳִ�д���"));
	   MapRetCode2Str.insert(make_pair(1007,"��������ظ�����"));
	   MapRetCode2Str.insert(make_pair(1008,"�ع��ɹ�����¼��־����"));
	   MapRetCode2Str.insert(make_pair(1009,"ָ�����׺��ѱ��ع��������ظ��ع�"));
	   MapRetCode2Str.insert(make_pair(1010,"�ֻ��û��ʺ�1Ч"));
	   MapRetCode2Str.insert(make_pair(1011,"���׺����ֻ����벻ƥ�䣬������ع�"));
	   MapRetCode2Str.insert(make_pair(1012,"�û��ڱ����ڣ�������ع�"));
	   MapRetCode2Str.insert(make_pair(1013,"����ԭ��ֵ��¼��������ʧ�ܣ�ErrorType�����ع��û����ʧ��"));
	}
	std::string operator()(int key)
	{
	     	MapRetCode2StrPtr p = MapRetCode2Str.find(key);
		if (p!=MapRetCode2Str.end())
		    return p->second;
		else
		    return std::string("δ֪��������");
	}
   private:
	std::map<int, std::string> MapRetCode2Str;
}mapRetCode;

int main()
{
  cout << mapRetCode(1099);
	return 0;
}

