
// create by derek 20070420

#ifndef __R5_SYSTEM_INFORMATION_
#define __R5_SYSTEM_INFORMATION_

#include <limits.h> 
#include <assert.h>
#include <string.h>

enum 
{
    en_lib_path=0,
    en_etc_path,
    en_bin_path,
    en_home_path,
    en_log_path,
    
    en_end_path
};

class SysInfo
{
public:
    SysInfo(){ memset(path, 0, sizeof(path)); init(); }
    ~SysInfo() { }
    
public:
    const char *getPath(int type)
    {
        assert(type>=en_lib_path);
        assert(type<en_end_path);
        return path[type]; 
    }
    
public:
    void init();
    
private:
    char path[en_end_path][PATH_MAX];
};

extern SysInfo g_sys_info;

#define BIN_PATH() g_sys_info.getPath(en_bin_path)
#define LIB_PATH() g_sys_info.getPath(en_lib_path)
#define ETC_PATH() g_sys_info.getPath(en_etc_path)
#define HOME_PATH() g_sys_info.getPath(en_home_path)
#define LOG_PATH() g_sys_info.getPath(en_log_path)

// #define INIT_SYS_INFO() //g_sys_info.init()

#endif //__R5_SYSTEM_INFORMATION_

