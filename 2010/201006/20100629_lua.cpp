/*
 * 20100625_lua.cpp
 *
 *  Created on: 2012-8-20
 *      Author: bonly
 */
#include<iostream>
extern "C"
{
    #include "lua.h"
    #include "lauxlib.h"
    #include "lualib.h"
}

#include<string>
using namespace std;
//#include<luabind/lua_include.hpp>
//#include<luabind/function.hpp>
#include<luabind/luabind.hpp>
void testFunc()
{
    cout<<"helo there, i am a cpp fun"<<endl;
}
int main(int argc, char* argv[])
{
    //首先声明luaState环境
    using namespace luabind;
    //lua_State* L = lua_open();  //也可以用luaL_newState()函数
    lua_State* L = luaL_newstate();
    luaL_openlibs(L);   //注意将lua默认库打开，要不会出现N多错误的，比如print函数都没有
    //将c++中的函数暴露给lua
    module(L, "cppapi")
    [
        def("testFunc", (void(*)(void))testFunc)
    ];

    //加载lua脚本并执行，我们临时起名test.lua
    luaL_dofile(L, "test.lua");
    /*
    int ret = luaL_loadfile(L, "test.lua");
    clog << "ret: " << ret << endl;
    */
    try
    {
        //调用lua中的整形全局变量
        int nLuaGlobal =     luabind::object_cast<int>(luabind::globals(L)["nGlobal"]) ;
        //调用lua中的字符串变量
        string strLuaGlobal = luabind::object_cast<string>(luabind::globals(L)["strGlobal"]);
        //获取table,方法一，通过luabind::object 固有方法
        luabind::object luaTable = luabind::globals(L)["t"] ;
        string name=luabind::object_cast<string>(luaTable["name"]) ;
        int age = luabind::object_cast<int>(luaTable["age"]) ;
        //获取table，方法二，通过gettable
        string desc = luabind::object_cast<string>(luabind::gettable(luaTable,"desc"));
        //下面是调用lua中函数
        int nAddRes = luabind::call_function<int>(L, "add", 3, 4) ;
        string strEchoRes = luabind::call_function<string>(L, "strEcho", "c++参数") ;

        luaTable["myval"] = "a test";
        clog << "c get lua myval:  " << luabind::object_cast<string>(luaTable["myval"]) << endl;

        luabind::call_function<void>(L, "pf");
    }
    catch(...)
    {
        cout<<"error"<<endl;
    }

    lua_close(L);
    return 0;
}

/**
lua脚本:
nGlobal = 10 --一个全局的整形变量
strGlobal = "hello i am in lua" --一个全局的字符串变量
--一个返回值为int类型的函数
function add(a, b)
    return a+b
end
--一个返回值为string类型的函数
function strEcho(a)
    print(a .. 10)
    return 'haha i have print your input param'
end
cppapi.testFunc() --调用c++暴露的一个测试函数
t={name='ettan', age=23, desc='正值花季年龄'}
 */


