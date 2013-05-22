
// create by derek 20070410
// modify 20070428
// modify 20070429

#ifndef __R5_LOG__
#define __R5_LOG__

#include <stdio.h>
#include <stdarg.h>
#include <sys/time.h>
#include <time.h>
#include <unistd.h>
#include <string>
#include <assert.h>
#include <dirent.h>
#include <sys/stat.h>               //迁移到linux时增加，hxt，20080730
#include <sys/types.h>
using std::string;

// 注意：此日志类非线程安全

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------

#define R5_LOG(le, X) \
	do{\
      if ( g_r5_log.checkLevelFile(le) )\
      {\
        g_r5_log.setConditional(le, __LINE__, __FILE__);\
        g_r5_log.log1 X;\
      }\
      if ( g_r5_log.checkLevelTerm(le) )\
      {\
        g_r5_log.setConditional(le, __LINE__, __FILE__);\
        g_r5_log.log2 X;\
      }\
    }while(0)

//-----------------------------------------------------------------------------
#define R5_INFO(X) \
      R5_LOG(r5_log_info, X);

#define R5_WARN(X) \
      R5_LOG(r5_log_warn, X);

#define R5_ERROR(X) \
       R5_LOG(r5_log_error, X);


//-----------------------------------------------------------------------------
#ifdef NR5_LOG

#define R5_DEBUG(X) do { }while(0)
#define R5_TRACE(X) do {} while (0)
#define _R5_TRACE_(X) do {} while (0)

#else

#define R5_TRACE(X) R5_Trace ___ (X, __LINE__, __FILE__)

#define _R5_TRACE_(X) \
  do { \
      R5_LOG(r5_log_trace, X);\
  }while(0)

#define R5_DEBUG(X) \
  do { \
      R5_LOG(r5_log_debug, X);\
  }while(0)
  
#endif
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// 日志单个文件最大记录数
#define R5_LOG_SIZE        1000000

//-----------------------------------------------------------------------------
#define SET_LOG_LEVEL(fl, tm) g_r5_log.setLevelClass(fl, tm)

//-----------------------------------------------------------------------------
#define FLUSH_LOG_FILE() g_r5_log.flush()

//-----------------------------------------------------------------------------
#define SET_LOG_NAME_HEAD(x) g_r5_log.setLogNameHead(x)

//-----------------------------------------------------------------------------

#define SET_LOG_SUB_DIR(x) g_r5_log.setLogSubDir(x)

//-----------------------------------------------------------------------------

#define SET_LOG_DIR(x) g_r5_log.setLogDir(x)

//-----------------------------------------------------------------------------
// add 20081207
#define LOG_CHECK_TRACE() (g_r5_log.checkLevelFile(r5_log_trace)||g_r5_log.checkLevelTerm(r5_log_trace))
#define LOG_CHECK_DEBUG() (g_r5_log.checkLevelFile(r5_log_debug)||g_r5_log.checkLevelTerm(r5_log_debug))
#define LOG_CHECK_INFO() (g_r5_log.checkLevelFile(r5_log_info)||g_r5_log.checkLevelTerm(r5_log_info))
#define LOG_CHECK_WARN() (g_r5_log.checkLevelFile(r5_log_warn)||g_r5_log.checkLevelTerm(r5_log_warn))
#define LOG_CHECK_ERROR() (g_r5_log.checkLevelFile(r5_log_error)||g_r5_log.checkLevelTerm(r5_log_error))
#define LOG_CHECK_FATAL() (g_r5_log.checkLevelFile(r5_log_fatal)||g_r5_log.checkLevelTerm(r5_log_fatal))
//-----------------------------------------------------------------------------


#define BASE_R5_LOG(level, ...)\
        R5_LOG(level, ("[base] " __VA_ARGS__)) 
#define BASE_R5_DEBUG(...)//\ R5_DEBUG(("[base] " __VA_ARGS__)) 
#define BASE_R5_INFO(...)\
        R5_INFO(("[base] " __VA_ARGS__)) 
#define BASE_R5_ERROR(...)\
        R5_ERROR(("[base] " __VA_ARGS__)) 
#define BASE_R5_WARN(...)\
        R5_WARN(("[base] " __VA_ARGS__)) 
#define BASE_R5_TRACE(...)\
        R5_TRACE("[base] " __VA_ARGS__) 

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
enum LOG_LEVEL
{
    r5_log_trace = 0,
    r5_log_debug = 20,
    r5_log_info = 40,
    r5_log_warn = 60,
    r5_log_error = 80,
    r5_log_fatal = 100,
    r5_log_end = 120
};

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
class R5_Log
{
public:
    R5_Log():line(0), file(0), level(r5_log_trace), level_file(-1), level_term(-1),
               os_log(NULL), log_size(0), log_count(0)//, isdefalutpath(0),os_custom(0)
    {
        memset(name, 0, PATH_MAX);
        memset(&date, 0, sizeof(date));
        memset(prefix, 0, sizeof(prefix));
        memset(subdir, 0, sizeof(subdir));        
        memset(fulldir, 0, sizeof(fulldir));
    }
            
    ~R5_Log()
    { 
        if (os_log != NULL )
        {
            fclose(os_log);
            
            char name2[PATH_MAX] = { 0 };
            snprintf(name2, PATH_MAX, "%s.tmp", name);
                        
            rename(name2, name);
        }
    }
    
public:
    int setLogDir(const char *dir);
    int setLogSubDir(const char *dir);
    void setLogNameHead(const char *nm);
    FILE *logOpen();
    FILE *logOpen(const char *name);
    void logOnlyClose(); 

    //void log(const char *format, ...); // 写日志和终端
    void log1(const char *format, ...); //写日志
    void log2(const char *format, ...); //写终端
    
    void setConditional(int le, int li, const char * f)//, FILE* os=0)
    {
        assert(le>=r5_log_trace&&le<=r5_log_end);
        level = le;
        //os_custom = os;
        line = li;
        file = f;
    }
    
    void setLevelClass(int le_fl, int le_tm)
    {
        level_file = le_fl;
        level_term = le_tm;
    }
    
    bool checkLevelFile(int le)
    {
        //printf("%d, %d\n", le, level_file);
        return le>=level_file;
    }
    
    bool checkLevelTerm(int le)
    {
        //printf("%d, %d\n", le, level_term);
        return le>=level_term;
    }
    
    void flush()
    {
        if (os_log != NULL ) { fflush(os_log); }
    }
    
    bool checkPath(const char *path);
private:
    int line;
    const char* file;
    int level;
    
    int level_file;
    int level_term;
    
    FILE *os_log;
    //FILE *os_custom;
    
    unsigned int log_size; // log size
    unsigned int log_count; // log count
    char name[PATH_MAX];
    char prefix[128];    
    char subdir[128];
    char fulldir[PATH_MAX];
    //int  isdefalutpath;
    struct  tm  date;
};

extern R5_Log g_r5_log;

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
class R5_Trace
{
public:
    R5_Trace(const char *msg, int line, const char *file)
    {
        name = msg;
        _R5_TRACE_(("calling %s in file `%s' on line %d\n", msg, file, line));
    }
    
    ~R5_Trace()
    {
        _R5_TRACE_(("leaving %s\n", name));
    }
    
private:
    const char *name;
};

#endif //__R5_LOG__


