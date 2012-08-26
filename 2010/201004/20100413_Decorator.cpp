/*****************************************************
* 模式名：装饰模式
* 时间：2010.6.3
*                                       -by bonly
*****************************************************/

#include <iostream>
#include <string>
using namespace std;

//////////////////////////////////////////////////////////////////////////
class cPerson;		// 人物类 （Component类）
class cFinery;		// 服装类 （Decorator类）
class TShirt;		// T-Shirt类 （具体的装饰类）
class BigTrouser;	// 裤子类 （具体的装饰类）
class Sneaker;		// 球鞋类 （具体的装饰类）
class LeatherShoes;	// 皮鞋类 （具体的装饰类）
class Tie;			// 领带类 （具体的装饰类）
class Suit;			// 西装类 （具体的装饰类）
//////////////////////////////////////////////////////////////////////////

// 人物类
// Component类，所有类的基类
// 同时是添加职责或功能的对象
// 定义虚函数Show，执行最基本的操作
// 注意：Component类完全不知道Decorator类的存在，即其完全不依赖于Decorator类
class cPerson
{
private:
	std::string m_name; // 人名

public:
	cPerson() {} // 必须带默认构造器，因为本类要用做多态，而其他装饰类并不需实际创建cPerson对象
	cPerson(const std::string& name) : m_name(name) {}
	virtual ~cPerson() {}

	// 返回装饰后的字符串
	virtual std::string Show() { return std::string("装扮的" + m_name); }
};

// 服装类
// Decorator类，所有装饰类的父类
// 定义函数Decorate，连接到下一个装饰对象，用于生成一条装饰链
// 成员m_pComponent指向装饰链的下一个装饰对象，装饰链最终终止于Component类，即人物类
class cFinery : public cPerson
{
protected:
	// 用cPerson不用cFinery，是因为装饰链是终止于cPerson的，所以必须要用cPerson
	// 这也解释了为什么cFinery要从cPerson继承的原因，就是要实现这种多态
	// 而实际上cFinery和cPerson之间并无意义上的继承关系
	cPerson* m_pComponent;

public:
	cFinery() : m_pComponent(NULL) {}
	virtual ~cFinery() {}

	// 装饰函数，生成一条装饰链
	// 返回this是为了可以如下使用：a.Decorate(b.Decorate(c));
	cPerson* Decorate(cPerson* component) { m_pComponent = component; return this;}

	// 返回装饰链所装饰的整个字符串
	std::string Show()
	{
		if(m_pComponent)
			return m_pComponent->Show();

		return std::string("");
	}
};

// T-Shirt类 （具体的装饰类）
class TShirt : public cFinery
{
public:
	// 这里要连带调用父类的Show，相当于调用装饰链后面对象的Show
	// 即相当于语句m_pComponent->Show();
	std::string Show() { return std::string("大T恤 ") + cFinery::Show(); }
};

// 裤子类 （具体的装饰类）
class BigTrouser : public cFinery
{
public:
	std::string Show() { return std::string("垮裤 ") + cFinery::Show(); }
};

// 球鞋类 （具体的装饰类）
class Sneaker : public cFinery
{
public:
	std::string Show() { return std::string("破球鞋 ") + cFinery::Show(); }
};

// 皮鞋类 （具体的装饰类）
class LeatherShoes : public cFinery
{
public:
	std::string Show() { return std::string("皮鞋 ") + cFinery::Show(); }
};

// 领带类 （具体的装饰类）
class Tie : public cFinery
{
public:
	std::string Show() { return std::string("领带 ") + cFinery::Show(); }
};

// 西装类 （具体的装饰类）
class Suit : public cFinery
{
public:
	std::string Show() { return std::string("西装 ") + cFinery::Show(); }
};

// 测试的main函数
int main()
	{
	using namespace std;

	cPerson xc(string("小菜"));

	Sneaker pqx;
	BigTrouser kk;
	TShirt dtx;

	dtx.Decorate(kk.Decorate(pqx.Decorate(&xc)));
	cout<<endl<<"第一种装扮："<<dtx.Show()<<endl;

	LeatherShoes px;
	Tie ld;
	Suit xz;

	xz.Decorate(ld.Decorate(px.Decorate(&xc)));
	cout<<endl<<"第二种装扮："<<xz.Show()<<endl;

	Sneaker pqx2;
	LeatherShoes px2;
	BigTrouser kk2;
	Tie ld2;
	ld2.Decorate(kk2.Decorate(px2.Decorate(pqx2.Decorate(&xc))));
	cout<<endl<<"第三种装扮："<<ld2.Show()<<endl;

	return 0;
}
