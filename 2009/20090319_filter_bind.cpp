//============================================================================
// Name        : search.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
#include <boost/shared_ptr.hpp>
#include <boost/assign/std/vector.hpp>
#include <boost/foreach.hpp>
#include <boost/format.hpp>
#include <boost/bind.hpp>
using namespace std;
using namespace boost;
using namespace boost::assign;

bool check(int i)
{
	return i%2 == 0;
}

void cp_int()
{
	vector<int> o_int;
	o_int +=1,2,3,4,5,6;

	BOOST_FOREACH(int i, o_int)
	{
		cout << format("%1% \n")%i;
	}


	vector<int> n1_int;
	remove_copy_if (o_int.begin(),o_int.end(),
			            inserter(n1_int,n1_int.begin()),
			            bind(&check,_1));

	copy (n1_int.begin(),n1_int.end(),ostream_iterator<int>(cout,"\t"));
}

template <typename T> void print_var(const char* k, T v)
{
  cout << k << " : " << v << endl;
}
#define PRINT_VAR(k) print_var(#k, k)
typedef int      SQLINTEGER;
//生命周期时间规则
struct LIFECYCLE_TIME_RULE
{
    char             BELONG_DISTRICT[6+1];
    SQLINTEGER       INITIAL_LIFECYCLE;
    SQLINTEGER       TARGET_LIFECYCLE;
    SQLINTEGER       SUBSCRIBER_CLASS;
    SQLINTEGER       REMIND_DAYS;
    SQLINTEGER       INITIAL_LIFECYCLE_MAXDAYS;
    LIFECYCLE_TIME_RULE()
    {memset(this,0,sizeof(LIFECYCLE_TIME_RULE));}
    void log();
};

void LIFECYCLE_TIME_RULE::log()
{
	cout << format("----BEGIN[LIFECYCLE_TIME_RULE]----\n");
	PRINT_VAR(BELONG_DISTRICT);
	PRINT_VAR(INITIAL_LIFECYCLE);
	PRINT_VAR(TARGET_LIFECYCLE);
	PRINT_VAR(SUBSCRIBER_CLASS);
	PRINT_VAR(REMIND_DAYS);
	PRINT_VAR(INITIAL_LIFECYCLE_MAXDAYS);
	cout << format("---- END[LIFECYCLE_TIME_RULE] ----\n");
}

bool belong_isnull(shared_ptr<LIFECYCLE_TIME_RULE> p)
{
	return strlen(p->BELONG_DISTRICT)==0;
}
struct ck
{
	bool operator()(shared_ptr<LIFECYCLE_TIME_RULE> p)
	{return strlen((p)->BELONG_DISTRICT)==0;}
};
void cp_rule()
{
  shared_ptr<LIFECYCLE_TIME_RULE> k1 (new LIFECYCLE_TIME_RULE);
  k1->REMIND_DAYS = 1;
  shared_ptr<LIFECYCLE_TIME_RULE> k2 (new LIFECYCLE_TIME_RULE);
  k2->REMIND_DAYS = 2;
  strcpy(k2->BELONG_DISTRICT,"020");
  shared_ptr<LIFECYCLE_TIME_RULE> k3 (new LIFECYCLE_TIME_RULE);
  k3->REMIND_DAYS = 3;

  vector<shared_ptr<LIFECYCLE_TIME_RULE> > o_rule;
  o_rule += k2,k3,k1;
  BOOST_FOREACH(shared_ptr<LIFECYCLE_TIME_RULE> p, o_rule)
  {
  	p->log();
  }

  vector<shared_ptr<LIFECYCLE_TIME_RULE> > n_rule;
  cout << "search for belong not null:\n";
  remove_copy_if (o_rule.begin(),o_rule.end(),
  		            inserter(n_rule,n_rule.begin()),
  		            //ck());
  		            bind(&belong_isnull,_1));
  BOOST_FOREACH(shared_ptr<LIFECYCLE_TIME_RULE> p, n_rule)
  {
  	p->log();
  }
}

int main()
{
  //cp_int();
	cp_rule();
	return 0;
}

