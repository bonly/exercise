#ifndef __COMPONENT_H__
#define __COMPONENT_H__
#include <string>
using namespace std;

//饮料抽象类
class Beverage{
    public:
        virtual string Get_description();
        virtual double Cost() = 0;
    
    protected:
        string description;
};


//抽象装饰者调用类
class CondimentDecorator: public Beverage{
    public:
        virtual string Get_description()=0;
};

//具体饮料，被装饰的对象
class Espresso: public Beverage{
    public:
        Espresso();
        ~Espresso();
        double Cost();
};
class HouseBlend: public Beverage{
    public:
        HouseBlend();
        ~HouseBlend();
        double Cost();
};

//具体调料类 
class Mocha: public CondimentDecorator{
    public:
        Mocha(Beverage *beverage);
        string Get_description();
        double Cost();
    private:
        Beverage *beverage; //记录被装饰者，即饮料
};

#endif

#include <iostream>
#include <string>
using namespace std;

string Beverage::Get_description(){
    return description;
}

//被装饰者类
HouseBlend::HouseBlend(){
    description = "House Blend Coffee";
}
HouseBlend::~HouseBlend(){
    cout << "~HouseBlend()"<< endl;
}
double HouseBlend::Cost(){
    return 0.89;
}

Espresso::Espresso(){
    description = "Espresso";
}
Espresso::~Espresso(){
    cout << "~Espresso()" << endl;
}
double Espresso::Cost(){
    return 1.99;
}

//装饰者类
Mocha::Mocha(Beverage *beverage){
    this->beverage=beverage;
}
string Mocha::Get_description(){
    return beverage->Get_description() + ",Mocha";
}
double Mocha::Cost(){
    return 0.20 + beverage->Cost();
}


using namespace std;
int main(int argc, char* argv[]){
    //订一杯Espresso, 不需要调料，并打印出描述及其价格
    Beverage *beverage = new Espresso();
    cout << beverage->Get_description() << " $" << beverage->Cost() << endl;
    delete beverage;
    
    //订一杯HouseBlend并加入Mocha调料，打印出描述及其价格
    Beverage *beverage2 = new HouseBlend();
    beverage2 = new Mocha(beverage2);
    cout << beverage2->Get_description() << " $"<< beverage2->Cost() << endl;
    delete beverage2;
    
    return 0;
}
