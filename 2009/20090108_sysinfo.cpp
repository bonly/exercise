
// create by derek 20070420

#include <stdio.h>
#include <unistd.h>
#include <string.h>
#include <stdlib.h>

#include "sysinfo.hpp"

void SysInfo::init()
{
    char *p_home = getenv("OCS_HOME");
    sprintf(path[en_home_path],"%s/", p_home);

    // init bin path
    char *p_bin = getenv("OCS_BIN");
    if(p_bin == NULL)
        snprintf(path[en_bin_path], PATH_MAX, "%sbin/", path[en_home_path]);
    else
        snprintf(path[en_bin_path], PATH_MAX, "%s/", p_bin);
    //printf("bin path:%s\n",path[en_bin_path]);
    
    // init etc path
    char *p_etc = getenv("OCS_ETC");
    if(p_etc == NULL)
        snprintf(path[en_etc_path], PATH_MAX, "%setc/", path[en_home_path]);
    else
        snprintf(path[en_etc_path], PATH_MAX, "%s/", p_etc);
    //printf("etc path: %s\n",path[en_etc_path]);
    
    // init lib path
    char *p_lib = getenv("OCS_LIB");
    if(p_lib == NULL)
        snprintf(path[en_lib_path], PATH_MAX, "%slib/", path[en_home_path]);
    else
        snprintf(path[en_lib_path], PATH_MAX, "%s/",p_lib);
    //printf("lib path: %s\n",path[en_lib_path]);
    
    // init log path
    char *p_log = getenv("OCS_LOG");
    if(p_log == NULL)
        snprintf(path[en_log_path], PATH_MAX, "%slogs/", path[en_home_path]);
    else
        snprintf(path[en_log_path], PATH_MAX, "%s/", p_log);
    //printf("log path: %s\n",path[en_log_path]);
}

SysInfo g_sys_info;

