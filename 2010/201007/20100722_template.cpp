//静态选择结构，名称就可以很明白的表达意思了
template<bool, class T, class F>
struct IF
{
        typedef T result;
};
template<class T, class F>
struct IF<false, T, F>
{
        typedef F result;
};

//静态FOR循环结构的实现，FOR循环的结束语句
struct STOP
{
        template<class E> 
        static void execute(E &e){}
};

//FOR循环的循环条件结构
struct LESS
{ 
        template<int x, int y> 
        struct code
        {
             enum { value = (x <  y) };
        }; 
};

//FOR循环结构实现
template<int from, class ConditionType, int to, int step, class StatementType>
struct FOR
{
        //真正的循环过程在这个函数里面，理解下面的过程将是理解整篇文章的关键
        //实际上这里采用的是模板递归实现的循环过程。
        template<class EnvironmentType>
        static void execute(EnvironmentType&e)
        {
                //执行本条语句
                IF<ConditionType::template code<from,to>::value
                        ,typename StatementType::template code<from>
                        ,STOP>::result::execute(e);
                //执行本条语句的下一条语句
                IF<ConditionType::template code<from,to>::value
                        ,FOR<from+step,ConditionType,to,step,StatementType>
                        ,STOP>::result::execute(e);
        }
};

//////////////////////////////////////////////////////////////////
//为了能够在其它的环境下也能构方便的编译运行该程序，我将Loki里面的
//typelist概念重新表述为cons,这样一方面可以出去Loki中的Typelist的神秘面纱
//另一方面也是因为这段代码比较简单
struct null_type;//类型列表的尾部元素表示类型已经结束了
template<class H,class T>
struct cons
{
        typedef H head_type;
        typedef T tail_type;
};
//计算类型列表的长度的元函数
template<class T>struct length;
template<>struct length<null_type>
{
        enum{value=0};
};
template<class H,class T>struct length<cons<H,T> >
{
        enum{value=1+length<T>::value};
};
//根据索引得到相应索引值类型的元函数
template<class Cons,unsigned int index>struct type;
template<class H,class T>struct type<cons<H,T>,0>
{
        typedef H result;
};
template<class H,class T,unsigned int i>struct type<cons<H,T>,i>
{
        typedef typename type<T,i-1>::result result;
};
////////////////////////////////////////////////////////////////////////////////
//下面是产生代码的元函数，相当于Loki里面的GenScatterHierarchy结构
template<class T,template<class>class Unit>
struct scatter : public Unit<T>
{
};
template<class H,class T,template<class>class Unit>
struct scatter<cons<H,T>,Unit>
        : public scatter<H,Unit>
        , public scatter<T,Unit>
{
        typedef cons<H,T> cons_type;
};
template<template<class>class Unit>
struct scatter<null_type,Unit>
{
};
////////////////////////////////////////////////////////////////////////////////
//下面是测试代码
#include <iostream>
#include <iterator>
#include <vector>
//仅仅为了兼容scatter产生代码的接口而追加一层代码
template<class T> struct VectorUnit:public std::vector<T>{};
template <class Cons>
class Class : public scatter<Cons,VectorUnit>
{
public:
        typedef Cons cons_type;//类型列表类型
        struct PRINT{//用于输出用的静态循环环境
                struct Environment//环境变量
                {
                        Environment(Class&ref):_ref(ref){}
                        Class&_ref;
                };
                struct Statement//静态循环语句
                {
                        template<int i>
                        struct code//语句代码
                        {
                                template<class E>
                                static void execute(E&e)
                                {
                                        //Your code here
                                        typedef typename type<cons_type,i>::result CT;
                                        VectorUnit<CT>&v = static_cast<VectorUnit<CT>&>(e._ref);
                                        std::copy(v.begin(),v.end(),std::ostream_iterator<CT>(std::cout," "));
                                        std::cout << std::endl;
                                }
                        };
                };
        };
        //下面就是使用静态循环代码实现的类型循环语句
        void print()
        {
                typename PRINT::Environment e(*this);//每一个循环都需要使用的环境变量
                //下面的这行代码就相当于很多的宏替换了，但是这行代码就可以使这个类以库的形式提供了
                FOR<0,LESS,length<cons_type>::value,+1,typename PRINT::Statement>::execute(e);
        }
};

int main()
{
        //给出5个类型的类型列表
        typedef cons<int,
                cons<short,
                cons<long,
                cons<float,
                cons<double,
                null_type> > > > > CONS;
        //用这个类型列表调用Class模板生成一个具体类CLASS
        typedef Class<CONS> CLASS;
        //定义一个这样的具体类的一个变量
        CLASS cls;
        //初始化类变量cls中的数据
        static_cast<VectorUnit<int>&>(cls).push_back(1);
        static_cast<VectorUnit<int>&>(cls).push_back(2);
        static_cast<VectorUnit<int>&>(cls).push_back(3);
        static_cast<VectorUnit<int>&>(cls).push_back(4);

        static_cast<VectorUnit<short>&>(cls).push_back(4);
        static_cast<VectorUnit<short>&>(cls).push_back(3);
        static_cast<VectorUnit<short>&>(cls).push_back(2);
        static_cast<VectorUnit<short>&>(cls).push_back(1);

        static_cast<VectorUnit<long>&>(cls).push_back(10);
        static_cast<VectorUnit<long>&>(cls).push_back(20);
        static_cast<VectorUnit<long>&>(cls).push_back(30);
        static_cast<VectorUnit<long>&>(cls).push_back(40);

        static_cast<VectorUnit<float>&>(cls).push_back(1.1);
        static_cast<VectorUnit<float>&>(cls).push_back(2.2);
        static_cast<VectorUnit<float>&>(cls).push_back(3.3);
        static_cast<VectorUnit<float>&>(cls).push_back(4.4);

        static_cast<VectorUnit<double>&>(cls).push_back(4.4);
        static_cast<VectorUnit<double>&>(cls).push_back(3.3);
        static_cast<VectorUnit<double>&>(cls).push_back(2.2);
        static_cast<VectorUnit<double>&>(cls).push_back(1.1);
        //现在就可以利用前面的静态循环print函数输出结果了
        cls.print();
        return 0;
}
