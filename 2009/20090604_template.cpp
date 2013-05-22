#include <iostream>

using namespace std;

template< class A, class B>
int func(void)
{
	A a = 1;
	B b = 0;
	return a - b;
}

typedef int ( * func_type) (void );

class TT
{
public:
	TT( int ){}
};

void Func( const TT & tt )
{
}

class Checker
{
public:
	virtual bool CheckCond( ) const
	{
		return true;
	}
	virtual ~Checker(){}
};

template<class T, class S,template<class ,class > class FIELD >
class Checker_Impl
	: public Checker
{
public:
	enum OP_TYPE{ EQ, NE, LT, LE, GT, GE };
	Checker_Impl( OP_TYPE Op, const T&t )
		: m_Op( Op )
		, m_Value( t )
	{}
	bool CheckCond(  )
	{
		FIELD<T,S> AAA;
		return false;
	}
public:
	OP_TYPE  m_Op;
	const T& m_Value;
};

class Where
{
public:
	Where( const Checker& checker )
		: m_checker( checker )
	{}
private:
	const Checker &m_checker;
};

typedef struct
{
	int FIELD_NAME;
}MY_FIELD;

// T 表示数据类型
// O 表示所在的结构
template< class T, class S >
class FIELD_NAME							
{											
public:										
	FIELD_NAME()							
		: m_Value(0), m_OpType(ASSIGN)		
	{}										
	FIELD_NAME(const T &t)					
		: m_Value(t), m_OpType( ASSIGN )	
	{}										
public:										
	FIELD_NAME & operator+( const T &t )	
	{										
		m_Value = t;						
		m_OpType = ADD;						
		return *this;						
	}										
	FIELD_NAME & operator-( const T &t )	
	{										
		m_Value = t;						
		m_OpType = SUB;						
		return *this;						
	}																		
	Checker operator==( const T&t )			
	{
		typedef Checker_Impl<T, S, FIELD_NAME> XX;
		return Checker_Impl<T, S, FIELD_NAME>( Checker_Impl<T, S, FIELD_NAME>::EQ ,t );
	}										
private:									
	T m_Value;								
	enum { ASSIGN, ADD, SUB, MUL, DIV }m_OpType;
};


int main(void)
{
	typedef struct __a
	{
		func_type x;
	} AAA;
	AAA a,b,c;
	a.x=func<int,int>;
	b.x=func<long,long>;
	c.x=func<short,short>;
	cout << "1" << a.x() << endl;
	cout << "2" << b.x() << endl;
	cout << "3" << c.x() << endl;
	Func( TT(1) );
	FIELD_NAME<int, MY_FIELD> N;
	return 0;
}


