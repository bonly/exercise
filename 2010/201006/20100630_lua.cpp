/*
 * 20100630_lua.cpp
 *
 *  Created on: 2012-8-21
 *      Author: bonly
 */
#include <stdexcept>
#include <mutex>
#include <LuaContext.h>
#include <LuaContext.cpp>
#include <iostream>

int main()
{
    Lua::LuaContext lua;
    lua.writeVariable("x", 5);
    lua.executeCode("x = x + 2;");
    std::cout << lua.readVariable<int>("x") << std::endl;
    return 0;
}


