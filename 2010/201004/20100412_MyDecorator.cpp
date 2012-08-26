//============================================================================
// Name        : MyDecorator.cpp
// Author      : 
// Version     :
// Copyright   : Your copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
using namespace std;
class Base
{
public:
	virtual int x()=0;
	virtual ~Base(){};
};

class Other
{
public:
	int ox;
};

template<typename T>
class OBJ1 : public T, public Base
{
public:
	T*  obj;
	virtual int x(){return obj->ox;}
};

template<typename T>
class OBJ2 : public Base
{
public:
	T* obj;
	virtual int x() {return obj->ox;}
};

int main()
{
	Other ot;
	ot.ox = 11;

	OBJ1<Other> myobj;
	myobj.obj = &ot;

	clog << "ox is: " << myobj.x() << endl;

	OBJ2<Other> myobj2;
	myobj2.obj = &ot;
	clog << "obj2's ox is: " << myobj2.x() << endl;
	return 0;
}
