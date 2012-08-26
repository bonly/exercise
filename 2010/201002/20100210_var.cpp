#include <iostream>


#define TBL(x) var##x
#define VAR(x) stvar TBL(x)

typedef struct 
{
    int ik;
    char name[20];
} stvar;

template<typename T>
class CA
{
   T iTA;
};

typedef int TBL(c); /// 变量类型用宏来定义

typedef struct
{
    int lk;
    char name[12];
} TBL(d);  /// 结构体类型名可以用宏来定义!

typedef CA<int> TBL(e); /// 用户自定义类型名也可以用宏来定义!
typedef CA<stvar> TBL(stvar);

int main()
{
    stvar TBL(a);  /// 定义了变量 vara 
    VAR(b); /// 定义了变量 varb
    TBL(c) fc;  

    return 0;
}

