#define BOOST_TEST_MODULE
#include <boost/test/included/unit_test.hpp>
#include <boost/shared_ptr.hpp>
#include <boost/operators.hpp>
#include <boost/foreach.hpp>
#include <boost/bind.hpp>
#include <vector>
using namespace std;
using namespace boost;

struct mydata
{
		int key, value, memo;
		bool operator<(const mydata &d2) const //注意两个const必写,因要与调用的定义一致,旧的less<>需要这个
		{return key<d2.key;}
};

struct boostdata : public less_than_comparable<boostdata>
        ,equality_comparable<boostdata>
{
		int key, value, memo;
		friend bool less_then(const shared_ptr<boostdata> d1, const shared_ptr<boostdata> d2); //带friend的当作本类外的函数
		//{ return d1->key<d2->key;}//本类外的friend函数不能定义在类中
		friend bool eq(const shared_ptr<boostdata> d1, const shared_ptr<boostdata> d2);
};
bool less_than(const shared_ptr<boostdata> d1, const shared_ptr<boostdata> d2)
{ return d1->key<d2->key;}
bool less_then(const shared_ptr<boostdata> d1, const shared_ptr<boostdata> d2) //
{ return d1->key<d2->key;}
bool eq(const shared_ptr<boostdata> d1, const shared_ptr<boostdata> d2)
{ return d1->key==d2->key && d1->value==d2->value;}

struct bmydata : public boostdata
{//父类的变量初始化不能写到初始化列表中
  bmydata(int k, int v, int m, int e):me(e){boostdata::key=k,boostdata::value=v,boostdata::memo=m;}
  int me;
};

struct byourdata : public boostdata
{
  byourdata(int k, int v, int m, int e):you(e){key=k,value=v,memo=m;}
  int you;
};

BOOST_AUTO_TEST_CASE(DEE)
{
   vector<shared_ptr<boostdata> > first;
   vector<shared_ptr<boostdata> > second;

   shared_ptr<boostdata> dt(new bmydata(1,12,34,89));
   //dt->key=1,dt->value=12,dt->memo=34;
   first.push_back(dt);
   dt = shared_ptr<bmydata>(new bmydata(7,21,19,88));
   //dt->key=7,dt->value=21,dt->memo=19;
   first.push_back(dt);
   dt = shared_ptr<bmydata>(new bmydata(2,13,35,87));
   //dt->key=2,dt->value=13,dt->memo=35;
   first.push_back(dt);
   dt = shared_ptr<bmydata>(new bmydata(5,11,39,82));
   //dt->key=5,dt->value=11,dt->memo=39;
   first.push_back(dt);

   shared_ptr<byourdata> yd(new byourdata(1,12,34,72));
   //yd->key=1,yd->value=12,yd->memo=34;
   second.push_back(yd);
   yd = shared_ptr<byourdata>(new byourdata(4,11,89,73));
   //yd->key=4,yd->value=11,yd->memo=89;
   second.push_back(yd);
   yd = shared_ptr<byourdata>(new byourdata(5,31,69,71));
   //yd->key=5,yd->value=31,yd->memo=69;
   second.push_back(yd);
   yd = shared_ptr<byourdata>(new byourdata(2,13,15,70));
   //yd->key=2,yd->value=13,yd->memo=15;
   second.push_back(yd);

   vector<shared_ptr<boostdata> > v;
   insert_iterator<vector<shared_ptr<boostdata> > > it(v, v.begin());

   //sort(first.begin(),first.end(),less<shared_ptr<boostdata> >());   //1 2 5 7
   sort(first.begin(),first.end(),bind(&less_than,_1,_2));                                  //1 2 5 7
   sort(second.begin(),second.end(),bind(&less_then,_1,_2));                                //1 2 4 5

   cout << "\nfirst data after sort is:\n";
   BOOST_FOREACH(shared_ptr<boostdata> i, first)
   {
  	 cout << i->key << ":" << i->value << ":" << ((bmydata*)&(*i))->me << "\t";
   }

   cout << "\nsecond data after sort is:\n";
   BOOST_FOREACH(shared_ptr<boostdata> i, second)
   {
  	 cout << i->key << ":" << i->value << ":" << ((byourdata*)&(*i))->you << "\t";
   }
   it = set_intersection (first.begin(), first.end(), second.begin(), second.end(),
   		 inserter(v ,v.begin()),bind(&eq,_1,_2));

   cout << "\nright_intersection_eq boost struct result mine - your: ";
   BOOST_FOREACH(shared_ptr<boostdata> i, v)
   {
  	 cout << i->key << ":" << i->value << "\t";
   }

   v.clear();
   it = set_intersection (second.begin(), second.end(), first.begin(), first.end(),
   		 inserter(v ,v.begin()),bind(&eq,_1,_2));

   cout << "\nright_intersection_eq boost struct result your - mine: ";
   BOOST_FOREACH(shared_ptr<boostdata> i, v)
   {
  	 cout << i->key << ":" << i->value << "\t";
   }

}
/*
aCC -AA shared.cpp -o shar
HP-UX/mingw下正确,没有内存泄露
cygwin下内存泄露
如果用list,在使用sort时报缺少操作符一堆
*/


