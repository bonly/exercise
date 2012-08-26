/**
 * \file 20091103_doxygen.cpp 
 * \brief 测试文档生成 \b 后面的字是粗体
 * 
 * 空一行后是详述,至少定义一次文件"@file"或"\file"
 * "@" "\"是一样的,有以下这些命令:
 * struct: 结构
 * union: 联合
 * enum: 枚举
 * fn: 函数
 * var: 变量,类型转换,枚举值其中之一
 * def: #define
 * typedef: 类型转换
 * namespace: 名字空间
 * package: java包
 * interface: IDL接口
 * 以下地址有样例:
 * http://www.stack.nl/~dimitri/doxygen/commands.html
 *
 */
#include <cstdio>
#include <cpthread>

/** \def MAX(a,b)
 * a macro that return the maximun of \a a and \a b.
 *
 * 这是详细说明
 */
#define MAX(a,b) (((a)>(b))?(a):(b))

/**
 * \namespace myspa
 */
namespace myspa{

/** \class pubA
 * \brief 一段话
 * 类pubA
 */
class pubA
{
    /**
     * pubA 类的说明
     * 
     */
    public: 
        ///内部的公有变量
        int myint; ///< @see kf 
        float kf;///<内部公有变量后的定义说明

    private:
        int _in_pit; ///<内部私有变量
        char _in_char[10]; ///<可见的内部变量
        
    public:
        pubA(){}
        ~pubA(){}
        /**
         * 测试函数
         * @param c1 参数
         */
        int fun(int c1/**<[in] 输入*/, char *buf/**<[out]*/, int &len/**<[in,out]*/){return 0;}
};

}

/** 文档样例:
 * - 线程事件
 *   -# 一行
 *   -# 二行
 * - 另一个事件
 *   -# 另一行
 */
void another_run(){}

/**
 * \fn void run()
 * 定义函数:运行线程
 * \exception 无异常
 * \return 无值返回
 */
void run()
{
    int k=0; //< \var 定义k
    pthread_exit(0);
}

/**
 * \fn int main()
 * 定义主函数
 */

int main()
{
    pthread_t thread;  
    long count = 0;
    while(1)
    {
        if (rc = pthread_create(&thread,0,run,0))
        {
            print("error, rc is %d, so far %ld threads created\n", rc, count);
            perror("fail:");
            return -1;
        }
        count ++;
    }
    return 0;
}

/**
 * 此线程创建后并没有pthread_join，因此泄漏内存
 */

/**  
 **  A list of events: 
 *   *  <ul> 
 *   *  <li> mouse events 
 *   *     <ol> 
 *   *     <li>mouse move event 
 *   *     <li>mouse click event\n 
 *   *         More info about the click event. 
 *   *     <li>mouse double click event 
 *   *     </ol> 
 *   *  <li> keyboard events 
 *   *     <ol>      
 *   *     <li>key down event 
 *   *     <li>key up event 
 *   *     </ol> 
 *   *  </ul> 
 *   *  More text here. 
 **/ 

/// \author bonly hayes
/*
 * \cond doxygen忽略开始
 * \endcond 忽略结束
 */
