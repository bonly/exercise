/*
 * Decorator.h
 *
 *  Created on: 2012-1-6
 *      Author: Bonly
 * @brief 装饰模式 decorator
 */

#ifndef DECORATOR_H_
#define DECORATOR_H_

template<typename TC>
class BASE
{
public:
	int x(){return TC::cx;} //这里使用不了子类的数据
};

template<typename TB>
class DEC : public TB
{
public:
	TB* base;
	//Decorator(TB *b):base(b){}
	void setBase(TB *b){base = b;}
};

class Child : public BASE<Child>
{
public:
	int cx;
	int cy;
};

template<typename TD, typename TB>
class MyOt : public TD
{
public:
	int mx;
	int my;
	void setBase(TB *b){TD::setBase(b);}
};

#endif /* DECORATOR_H_ */

/*
 * Decorator.cpp
 *
 *  Created on: 2012-1-6
 *      Author: Bonly
 */
//#include "Decorator.h"

#include <iostream>

int main()
{
	Child ch;
	ch.cx = 11;
	//MyOt<DEC<Child >, BASE<Child> > myt; ///相当于从Child<BASE>继承
    MyOt<DEC<BASE<Child> >, BASE<Child> > myt;  ///从BASE继承
	myt.setBase((Child*)&ch);

	//std::clog << myt.x() << std::endl; /// 没有实例的错误 error: object missing in reference to `Child::cx'
	//std::clog << myt.base->x() << std::endl;

    return 0;
}


