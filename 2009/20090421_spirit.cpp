/*
 * t_spirit.cpp
 *
 *  Created on: 2009-6-9
 *      Author: Bonly
 */
#include <iostream>
#include <string>
#include <vector>
#include <boost/spirit/include/classic_core.hpp>
#include <boost/lexical_cast.hpp>
using namespace BOOST_SPIRIT_CLASSIC_NS;
using namespace std;

int
main ()
{
	string sp ("3006012;999;30060121:3006012:6000:210203000,30060122:3006012:104857600:210203000,;");
  rule<> r = *(real_p >> ';' >> real_p >> ';'
             >> *(real_p >> ':' >> real_p >> ':' >> real_p >> ':' >> real_p >> ',')
             >> ';');
  bool ret = parse(sp.c_str(), r, space_p).full;
	if (ret) cerr << "³É¹¦\n";
	else cerr << "Ê§°Ü\n";

	int TARIFF_ID=0;
	int PRICE=0;
  char count[255];
  memset(count,0,255);

  sscanf(sp.c_str(),"%d;%d;%s;",&TARIFF_ID,&PRICE,count);
  string cnt(count);

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

    printf("COUNTER_TARIFF_ID:%d\tTARIFF_ID:%d\tFREE_VALUE:%d\tCOUNTER_TYPE_ID:%d\n",
    		    COUNTER_TARIFF_ID,TARIFF_ID,FREE_VALUE,COUNTER_TYPE_ID);
  }

	printf("TARIFF_ID: %d\n",TARIFF_ID);
	printf("PRICE: %d\n",PRICE);
	printf("count: %s\n",count);

	return 0;
}

