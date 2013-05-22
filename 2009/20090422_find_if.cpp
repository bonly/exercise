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
#include <boost/spirit/include/classic_core.hpp>
#include <boost/lexical_cast.hpp>
#include <boost/foreach.hpp>
using namespace BOOST_SPIRIT_CLASSIC_NS;
using namespace std;
struct Pos
{
	long key;
	int  line;
	bool operator()(const Pos& p)
	{
		return this->key==p.key;
	}
};

vector<Pos> vpos;

void getlind(string &sp,int lin)
{
	if (sp.empty()) return;
//  rule<> r = *(real_p >> ';' >> real_p >> ';'
//             >> *(real_p >> ':' >> real_p >> ':' >> real_p >> ':' >> real_p >> ',')
//             >> ';');
//  bool ret = parse(sp.c_str(), r, space_p).full;
//	if (ret) cerr << "成功\n";
//	else cerr << "失败\n";

	int TARIFF_ID=0;
	int PRICE=0;
  char count[255];
  memset(count,0,255);

  sscanf(sp.c_str(),"%d;%d;%s;",&TARIFF_ID,&PRICE,count);
  string cnt(count);

//printf("===BEGIN===\n");
  int pot = 0;
  while (-1 != (pot = cnt.find_first_of(',',0)))
  {
  	int COUNTER_TARIFF_ID=0;
  	int TARIFF_ID=0;
  	int FREE_VALUE=0;
  	int COUNTER_TYPE_ID=0;

    string msg(cnt,0,pot);
    sscanf(msg.c_str(),"%d:%d:%d:%d,;",&COUNTER_TARIFF_ID,&TARIFF_ID,&FREE_VALUE,&COUNTER_TYPE_ID);
    cnt.erase(0,pot+1);

//    printf("COUNTER_TARIFF_ID:%d\tTARIFF_ID:%d\tFREE_VALUE:%d\tCOUNTER_TYPE_ID:%d\n",
//    		    COUNTER_TARIFF_ID,TARIFF_ID,FREE_VALUE,COUNTER_TYPE_ID);
  }

//	printf("TARIFF_ID: %d\n",TARIFF_ID);
//	printf("PRICE: %d\n",PRICE);
//	printf("count: %s\n",count);
//	printf("=== END  ===\n");

	Pos pos;
	pos.key = TARIFF_ID;
	pos.line = lin;
	vpos.push_back(pos);
}
int
main ()
{
	char buf[255];
	ifstream fi("tr.log");

	do
	{
		bzero(buf,255);
		int lin = fi.tellg();
		//printf("pos is: %d\n",lin);
		fi.getline(buf,255);

		string sp (buf);
		//cout << sp << endl;

		getlind(sp, lin);
	}while(fi.good());

	//BOOST_FOREACH(Pos pos,vpos)
	//{
	//	cout << pos.key << "\t" <<pos.line << endl;
	//}

  while(true)
  {
		cout << "input tariff id for search: ";
		Pos ke;
		cin >> ke.key;
		if (ke.key==0)
		  break;

		vector<Pos>::iterator p = find_if(vpos.begin(),vpos.end(),ke); //用实例或类名，但类名中的比较数据没初始化，需自己解决（全局）
		if (p != vpos.end())
		{
			printf("key: %ld\t line: %d\n",p->key,p->line);
			ke.line = p->line;

			char tbuf[255];
			bzero(tbuf,255);

			//fi.seekg (0, ios::beg); //无需重置到文件头
			fi.clear();  //getline 到文件结尾会设置流状态标志位;在seekg之前调用infile.clear();把标志位去除即可
			fi.seekg(ke.line);
			fi.getline(tbuf,255);
			string val(tbuf);
			cout << val << endl;
		}
		else
			printf("not found\n");
  }

	return 0;
}

