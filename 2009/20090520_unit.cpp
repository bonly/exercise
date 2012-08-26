//============================================================================
// Name        : set_union.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
#include <vector>
#include <iterator>
#include <algorithm>
#include <boost/foreach.hpp>
#include <boost/bind.hpp>
#include <boost/operators.hpp>
using namespace std;
using namespace boost;
//#include <c++/bits/stl_function.h> //可通过这个报错的头文件检查正确的函数定义

//set_union 并集
//set_intersection 交集
//set_difference  减集
//set_symmetric_difference 各相减一次的集

#define BOOST_TEST_MODULE
#include <boost/test/included/unit_test.hpp>

BOOST_AUTO_TEST_CASE(right_intersection)
{
   int first[] = {5,10,15,20,25};
   int second[] = {50,40,30,20,10};
   vector<int> v(10);
   vector<int>::iterator it;

   sort(first,first+5);   //5  10 15 20 25
   sort(second,second+5); //10 20 30 40 50

   it = set_intersection (first, first+5, second, second+5, v.begin());

   cout << "intersection has " << int (it-v.begin()) << " elements.\n";
   BOOST_FOREACH(int i, v)
   {
  	 cout << i << "\t";
   }
}

BOOST_AUTO_TEST_CASE(right_diff)
{
   int first[] = {5,10,15,20,25};
   int second[] = {50,40,30,20,10};
   vector<int> v(10);
   vector<int>::iterator it;

   sort(first,first+5);   //5  10 15 20 25
   sort(second,second+5); //10 20 30 40 50

   it = set_difference (first, first+5, second, second+5, v.begin());

   cout << "\nintersection has " << int (it-v.begin()) << " elements.\n";
   BOOST_FOREACH(int i, v)
   {
  	 cout << i << "\t";
   }
}

BOOST_AUTO_TEST_CASE(right_diff1)
{
   int first[] = {5,10,15,20,25};
   int second[] = {50,40,30,20,10};
   //vector<int> v(10);
   vector<int> v;
   insert_iterator<vector<int> > it(v, v.begin());

   sort(first,first+5);   //5  10 15 20 25
   sort(second,second+5); //10 20 30 40 50

   it = set_difference (second, second+5, first, first+5, //v.begin());
  		 inserter(v ,v.begin()));

   cout << "\nright_diff1 result: ";
   BOOST_FOREACH(int i, v)
   {
  	 cout << i << "\t";
   }
}

BOOST_AUTO_TEST_CASE(right_diff2)
{
   int first[] = {5,10,15,20,25};
   int second[] = {50,40,30,20,10};
   //vector<int> v(10);
   vector<int> v;
   insert_iterator<vector<int> > it(v, v.begin());

   sort(first,first+5);   //5  10 15 20 25  //如果不先排列,结果不正确
   sort(second,second+5); //10 20 30 40 50

   it = set_symmetric_difference (second, second+5, first, first+5, //v.begin());
  		 inserter(v ,v.begin()));

   cout << "\nright_diff2 symmetric result: ";
   BOOST_FOREACH(int i, v)
   {
  	 cout << i << "\t";
   }
}

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
		friend bool operator<(const boostdata &d1, const boostdata &d2) //boost需要这个
		{ return d1.key<d2.key;}
		friend bool operator==(const boostdata &d1, const boostdata &d2)
		{ return d1.key==d2.key && d1.value==d2.value;}
};

struct bmydata : public boostdata
{

};

struct byourdata : public boostdata
{

};

BOOST_AUTO_TEST_CASE(right_intersection_super)
{

	 vector<boostdata> first;
	 vector<boostdata> second;

	 boostdata dt;
	 dt.key=1,dt.value=12,dt.memo=34;
   first.push_back(dt);
   dt.key=2,dt.value=13,dt.memo=35;
   first.push_back(dt);
   dt.key=5,dt.value=11,dt.memo=39;
   first.push_back(dt);
   dt.key=7,dt.value=21,dt.memo=19;
   first.push_back(dt);

   dt.key=1,dt.value=12,dt.memo=34;
   second.push_back(dt);
   dt.key=2,dt.value=13,dt.memo=15;
   second.push_back(dt);
   dt.key=4,dt.value=11,dt.memo=89;
   second.push_back(dt);
   dt.key=5,dt.value=31,dt.memo=69;
   second.push_back(dt);

   vector<boostdata> v;
   insert_iterator<vector<boostdata> > it(v, v.begin());

   sort(first.begin(),first.end(),less<boostdata>());   //5  10 15 20 25
   sort(second.begin(),second.end());                   //10 20 30 40 50

   it = set_intersection (first.begin(), first.end(), second.begin(), second.end(),
   		 inserter(v ,v.begin()));

   cout << "\nright_intersection_eq boost struct result: ";
   BOOST_FOREACH(boostdata i, v)
   {
  	 cout << i.key << ":" << i.value << "\t";
   }

   v.clear();
   it = set_intersection (second.begin(), second.end(), first.begin(), first.end(),
   		 inserter(v ,v.begin()));

   cout << "\nright_intersection_eq boost struct result: ";
   BOOST_FOREACH(boostdata i, v)
   {
  	 cout << i.key << ":" << i.value << "\t";
   }
}

BOOST_AUTO_TEST_CASE(right_diff_boost)
{

	 vector<boostdata> first;
	 vector<boostdata> second;

	 boostdata dt;
	 dt.key=1,dt.value=12,dt.memo=34;
   first.push_back(dt);
   dt.key=2,dt.value=13,dt.memo=35;
   first.push_back(dt);
   dt.key=5,dt.value=11,dt.memo=39;
   first.push_back(dt);
   dt.key=7,dt.value=21,dt.memo=19;
   first.push_back(dt);

   dt.key=1,dt.value=12,dt.memo=34;
   second.push_back(dt);
   dt.key=2,dt.value=13,dt.memo=15;
   second.push_back(dt);
   dt.key=4,dt.value=11,dt.memo=89;
   second.push_back(dt);
   dt.key=5,dt.value=31,dt.memo=69;
   second.push_back(dt);

   vector<boostdata> v;
   insert_iterator<vector<boostdata> > it(v, v.begin());

   sort(first.begin(),first.end(),less<boostdata>());   //5  10 15 20 25
   sort(second.begin(),second.end());                   //10 20 30 40 50

   it = set_symmetric_difference (first.begin(), first.end(), second.begin(), second.end(),
   		 inserter(v ,v.begin()));

   cout << "\nright_diff symmetric boost struct result: ";
   BOOST_FOREACH(boostdata i, v)
   {
  	 cout << i.key << "\t";
   }
}

BOOST_AUTO_TEST_CASE(right_diff3)
{

	 vector<mydata> first;
	 vector<mydata> second;

   first.push_back((mydata){1,12,34});
   first.push_back((mydata){2,13,35});
   first.push_back((mydata){5,11,39});
   first.push_back((mydata){7,21,19});

   second.push_back((mydata){1,12,34});
   second.push_back((mydata){2,13,15});
   second.push_back((mydata){4,11,89});
   second.push_back((mydata){5,31,69});

   vector<mydata> v;
   insert_iterator<vector<mydata> > it(v, v.begin());

   sort(first.begin(),first.end(),less<mydata>());   //5  10 15 20 25
   sort(second.begin(),second.end());                //10 20 30 40 50

   it = set_difference (first.begin(), first.end(), second.begin(), second.end(),
   		 inserter(v ,v.begin()));

   cout << "\nright_diff struct result: ";
   BOOST_FOREACH(mydata i, v)
   {
  	 cout << i.key << "\t";
   }
}

BOOST_AUTO_TEST_CASE(right_diff4)
{

	 vector<mydata> first;
	 vector<mydata> second;

   first.push_back((mydata){1,12,34});
   first.push_back((mydata){2,13,35});
   first.push_back((mydata){5,11,39});
   first.push_back((mydata){7,21,19});

   second.push_back((mydata){1,12,34});
   second.push_back((mydata){2,13,15});
   second.push_back((mydata){4,11,89});
   second.push_back((mydata){5,31,69});

   vector<mydata> v;
   insert_iterator<vector<mydata> > it(v, v.begin());

   sort(first.begin(),first.end(),less<mydata>());   //5  10 15 20 25
   sort(second.begin(),second.end());                //10 20 30 40 50

   it = set_symmetric_difference (first.begin(), first.end(), second.begin(), second.end(),
   		 inserter(v ,v.begin()));

   cout << "\nright_diff symmetric struct result: ";
   BOOST_FOREACH(mydata i, v)
   {
  	 cout << i.key << ":" << i.value << "\t";
   }
}

#include <boost/shared_ptr.hpp>
using namespace boost;
BOOST_AUTO_TEST_CASE(right_intersection_eq)
{
  vector<shared_ptr<boostdata> > first;
  vector<shared_ptr<boostdata> > second;

  shared_ptr<boostdata> dt(new bmydata);
  dt->key=1,dt->value=12,dt->memo=34;
   first.push_back(dt);
   dt = shared_ptr<bmydata>(new bmydata);
   dt->key=7,dt->value=21,dt->memo=19;
   first.push_back(dt);
   dt = shared_ptr<bmydata>(new bmydata);
   dt->key=2,dt->value=13,dt->memo=35;
   first.push_back(dt);
   dt = shared_ptr<bmydata>(new bmydata);
   dt->key=5,dt->value=11,dt->memo=39;
   first.push_back(dt);


   shared_ptr<byourdata> yd(new byourdata);
   yd->key=1,yd->value=12,yd->memo=34;
   second.push_back(yd);
   yd = shared_ptr<byourdata>(new byourdata);
   yd->key=2,yd->value=13,yd->memo=15;
   second.push_back(yd);
   yd = shared_ptr<byourdata>(new byourdata);
   yd->key=4,yd->value=11,yd->memo=89;
   second.push_back(yd);
   yd = shared_ptr<byourdata>(new byourdata);
   yd->key=5,yd->value=31,yd->memo=69;
   second.push_back(yd);

   vector<shared_ptr<boostdata> > v;
   insert_iterator<vector<shared_ptr<boostdata> > > it;

   sort(first.begin(),first.end(),less<shared_ptr<boostdata> >());   //1 2 5 7
   sort(second.begin(),second.end());                                //1 2 4 5

   cout << "\ndata is:\n";
   BOOST_FOREACH(shared_ptr<boostdata> i, first)
   {
  	 cout << i->key << ":" << i->value << "\t";
   }

   set_intersection (first.begin(), first.end(), second.begin(), second.end(),
   		 inserter(v ,v.begin()));

   cout << "\nright_intersection_eq boost struct result mine - your: ";
   BOOST_FOREACH(shared_ptr<boostdata> i, v)
   {
  	 cout << i->key << ":" << i->value << "\t";
   }

   v.clear();
   set_intersection (second.begin(), second.end(), first.begin(), first.end(),
   		 inserter(v ,v.begin()));

   cout << "\nright_intersection_eq boost struct result your - mine: ";
   BOOST_FOREACH(shared_ptr<boostdata> i, v)
   {
  	 cout << i->key << ":" << i->value << "\t";
   }
}
/*
在cygwin下有错,但单独在hp-ux上正确,以后测试最后一用例单独在cygwin下情况
*** No errors detected
     79 [main] set_union 1816 _cygtls::handle_exceptions: Exception: STATUS_ACCESS_VIOLATION
    638 [main] set_union 1816 open_stackdumpfile: Dumping stack trace to set_union.exe.stackdump
*/

