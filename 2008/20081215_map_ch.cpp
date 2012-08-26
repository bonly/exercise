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
     MapRetCode2Str.insert(make_pair(2001,"成功"));
     MapRetCode2Str.insert(make_pair(1001,"用户不存在 | 帐户不存在 | 帐本不存在"));
	   MapRetCode2Str.insert(make_pair(1002,"用户余额不够"));
	   MapRetCode2Str.insert(make_pair(1003,"指定交易序列号不存在"));
	   MapRetCode2Str.insert(make_pair(1004,"指定交易不成功"));
	   MapRetCode2Str.insert(make_pair(1005,"回滚失败"));
	   MapRetCode2Str.insert(make_pair(1006,"充值系统执行错误"));
	   MapRetCode2Str.insert(make_pair(1007,"交易序号重复出错"));
	   MapRetCode2Str.insert(make_pair(1008,"回滚成功但记录日志出错"));
	   MapRetCode2Str.insert(make_pair(1009,"指定交易号已被回滚，不能重复回滚"));
	   MapRetCode2Str.insert(make_pair(1010,"手机用户帐号1效"));
	   MapRetCode2Str.insert(make_pair(1011,"交易号与手机号码不匹配，不允许回滚"));
	   MapRetCode2Str.insert(make_pair(1012,"用户在保留期，不允许回滚"));
	   MapRetCode2Str.insert(make_pair(1013,"更新原充值记录错误类型失败（ErrorType），回滚用户余额失败"));
	}
	std::string operator()(int key)
	{
	     	MapRetCode2StrPtr p = MapRetCode2Str.find(key);
		if (p!=MapRetCode2Str.end())
		    return p->second;
		else
		    return std::string("未知错误类型");
	}
   private:
	std::map<int, std::string> MapRetCode2Str;
}mapRetCode;

int main()
{
  cout << mapRetCode(1099);
	return 0;
}

