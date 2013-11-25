/**
 * @file 20110530_event.cpp
 * @brief 
 * @author bonly
 * @date 2013年11月23日 bonly Created
 */
extern "C"
{
    #include "lua.h"
    #include "lauxlib.h"
    #include "lualib.h"
}
#include<luabind/luabind.hpp>
#include <string>
#include <cstdio>

#define EVENT_SAMPLE 1000
std::string g_strEventHandler = "";

bool dostring(lua_State* L, const char* str)
{
        if (luaL_loadbuffer(L, str, std::strlen(str), str) || lua_pcall(L, 0, 0, 0))
        {
                const char* a = lua_tostring(L, -1);
                std::cout << a << "\n";
                lua_pop(L, 1);
                return true;
        }
        return false;
}


lua_State *g_pLua = 0;
extern "C" int _RegisterEvent(lua_State *L){
    g_strEventHanler = g_pLua->GetStringArgument(1, "");
}

void FireEvent(int id){
    if(g_strEventHandler != ""){
        char buf[254];
        sprintf(buf, "%s(%d)", g_strEventHandler.c_str(), id);
        dostring(g_pLua, buf);
    }
}

int main(int argc, char* argv[]){
    lua_State* g_pLua = luaL_newstate();
    return 0;
}
