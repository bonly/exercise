#include <iostream>
#include <string>
using namespace std;
int
main()
{
	  string desc;
		string source("23:45:32;13:23:26;");
		string psrc(source);
		int pot = 0;
		while( -1 != (pot = psrc.find_first_of(';',0)))
		{
			string msg(psrc,0,pot);
			int first=0,second=0;
			string kind;
			string money;
			string amount;
			if (0!=(first = msg.find_first_of(':',0)))
			{
						kind=msg.substr(0,first);
						if (0!=(second = msg.find_first_of(':',first+1)))
						{
						  money  = msg.substr(first+1,second-first-1);
						  amount = msg.substr(second+1);
						}
						else
						{
							money  = "0";
							amount = "0";
						}
			}
			else
			{
						kind   = msg;
						money  = "0";
						amount = "0";
			}
			psrc.erase(0,pot+1);

//      char prod_desc[1024+1];
//      char prod_name[50+1];
//      try{app.get_prod_describe(lexical_cast<int>(kind),prod_desc,1025,prod_name,51);}
//      catch(bad_lexical_cast &e)
//      {
//        R5_WARN(("prod_id err:%s\n",kind.c_str()));
//        continue;
//      }
//      float nMoney = (float)(lexical_cast<int>(money))/(float)100;
//      char  sMoney[30];memset(sMoney,0,30);
//      sprintf(sMoney,"月租%.2f元，",nMoney);
      desc.append (kind);
      desc.append (money);
      desc.append (amount);
      cerr << kind << " 应收 " << money << " 实收 " << amount << endl;
    }
}

