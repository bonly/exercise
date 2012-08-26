/*****************************************************
* ģʽ����װ��ģʽ
* ʱ�䣺2010.6.3
*                                       -by bonly
*****************************************************/

#include <iostream>
#include <string>
using namespace std;

//////////////////////////////////////////////////////////////////////////
class cPerson;		// ������ ��Component�ࣩ
class cFinery;		// ��װ�� ��Decorator�ࣩ
class TShirt;		// T-Shirt�� �������װ���ࣩ
class BigTrouser;	// ������ �������װ���ࣩ
class Sneaker;		// ��Ь�� �������װ���ࣩ
class LeatherShoes;	// ƤЬ�� �������װ���ࣩ
class Tie;			// ����� �������װ���ࣩ
class Suit;			// ��װ�� �������װ���ࣩ
//////////////////////////////////////////////////////////////////////////

// ������
// Component�࣬������Ļ���
// ͬʱ�����ְ����ܵĶ���
// �����麯��Show��ִ��������Ĳ���
// ע�⣺Component����ȫ��֪��Decorator��Ĵ��ڣ�������ȫ��������Decorator��
class cPerson
{
private:
	std::string m_name; // ����

public:
	cPerson() {} // �����Ĭ�Ϲ���������Ϊ����Ҫ������̬��������װ���ಢ����ʵ�ʴ���cPerson����
	cPerson(const std::string& name) : m_name(name) {}
	virtual ~cPerson() {}

	// ����װ�κ���ַ���
	virtual std::string Show() { return std::string("װ���" + m_name); }
};

// ��װ��
// Decorator�࣬����װ����ĸ���
// ���庯��Decorate�����ӵ���һ��װ�ζ�����������һ��װ����
// ��Աm_pComponentָ��װ��������һ��װ�ζ���װ����������ֹ��Component�࣬��������
class cFinery : public cPerson
{
protected:
	// ��cPerson����cFinery������Ϊװ��������ֹ��cPerson�ģ����Ա���Ҫ��cPerson
	// ��Ҳ������ΪʲôcFineryҪ��cPerson�̳е�ԭ�򣬾���Ҫʵ�����ֶ�̬
	// ��ʵ����cFinery��cPerson֮�䲢�������ϵļ̳й�ϵ
	cPerson* m_pComponent;

public:
	cFinery() : m_pComponent(NULL) {}
	virtual ~cFinery() {}

	// װ�κ���������һ��װ����
	// ����this��Ϊ�˿�������ʹ�ã�a.Decorate(b.Decorate(c));
	cPerson* Decorate(cPerson* component) { m_pComponent = component; return this;}

	// ����װ������װ�ε������ַ���
	std::string Show()
	{
		if(m_pComponent)
			return m_pComponent->Show();

		return std::string("");
	}
};

// T-Shirt�� �������װ���ࣩ
class TShirt : public cFinery
{
public:
	// ����Ҫ�������ø����Show���൱�ڵ���װ������������Show
	// ���൱�����m_pComponent->Show();
	std::string Show() { return std::string("��T�� ") + cFinery::Show(); }
};

// ������ �������װ���ࣩ
class BigTrouser : public cFinery
{
public:
	std::string Show() { return std::string("��� ") + cFinery::Show(); }
};

// ��Ь�� �������װ���ࣩ
class Sneaker : public cFinery
{
public:
	std::string Show() { return std::string("����Ь ") + cFinery::Show(); }
};

// ƤЬ�� �������װ���ࣩ
class LeatherShoes : public cFinery
{
public:
	std::string Show() { return std::string("ƤЬ ") + cFinery::Show(); }
};

// ����� �������װ���ࣩ
class Tie : public cFinery
{
public:
	std::string Show() { return std::string("��� ") + cFinery::Show(); }
};

// ��װ�� �������װ���ࣩ
class Suit : public cFinery
{
public:
	std::string Show() { return std::string("��װ ") + cFinery::Show(); }
};

// ���Ե�main����
int main()
	{
	using namespace std;

	cPerson xc(string("С��"));

	Sneaker pqx;
	BigTrouser kk;
	TShirt dtx;

	dtx.Decorate(kk.Decorate(pqx.Decorate(&xc)));
	cout<<endl<<"��һ��װ�磺"<<dtx.Show()<<endl;

	LeatherShoes px;
	Tie ld;
	Suit xz;

	xz.Decorate(ld.Decorate(px.Decorate(&xc)));
	cout<<endl<<"�ڶ���װ�磺"<<xz.Show()<<endl;

	Sneaker pqx2;
	LeatherShoes px2;
	BigTrouser kk2;
	Tie ld2;
	ld2.Decorate(kk2.Decorate(px2.Decorate(pqx2.Decorate(&xc))));
	cout<<endl<<"������װ�磺"<<ld2.Show()<<endl;

	return 0;
}
