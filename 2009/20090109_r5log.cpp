
// create by derek 20070420

#include "r5log.hpp"
#include "sysinfo.hpp"

#define IO_BUFFER_SIZE 64*1024

const char *LOG_HEAD[] = 
{
    "TRACE.",
    "DEBUG.",
    "INFO..",
    "WARN..",
    "ERROR.",
    "FATAL."
};

void R5_Log::setLogNameHead(const char *nm)
{
    // 日志的文件名产生规则：r5log_时间_进程ID_日志计数.log
    int len = strlen(nm);
    if(len != 0 && memcmp(prefix, nm, len)!=0)
    {
        strcpy(prefix, nm);
        log_count = 0;
        if (os_log != NULL )
        {
            fclose(os_log);
            os_log = NULL;
        }
    }
}

int R5_Log::setLogSubDir(const char *dir)
{
    int len = strlen(dir);
    if(len != 0&&memcmp(subdir, dir, len)!=0)
    {
	    char dirpath[PATH_MAX];
	    memset(dirpath, 0 ,PATH_MAX);
	    sprintf(dirpath, "%s%s",  LOG_PATH(), dir);
		if (!checkPath(dirpath))
			return -1;
			
        strcpy(subdir, dir); 
	    log_count = 0;
	    if (os_log != NULL )
	    {
	        fclose(os_log);
	        os_log = NULL;
	    }
        return 0;
	}
		
    return -1;
}

int R5_Log::setLogDir(const char *dir)
{
    int len = strlen(dir);
    if(len != 0 && memcmp(fulldir, dir, len)!=0)
    {
    	if (!checkPath(dir))
    		return -1;
    		
        strcpy(fulldir, dir);        
	    log_count = 0;
	    if (os_log != NULL )
	    {
	        fclose(os_log);
	        os_log = NULL;
	    }
    }
    return 0;
}


FILE *R5_Log::logOpen()
{
    if (os_log != NULL )
    {
        fclose(os_log);
        os_log = NULL;
        
        char name2[PATH_MAX] = { 0 };
        snprintf(name2, PATH_MAX, "%s.tmp", name);
                    
        rename(name2, name);
    }

    time_t  t;
    struct  tm  tt;
    time(&t);

#ifdef SunOS
    memcpy(&tt, localtime(&t), sizeof(struct tm));
#else
    localtime_r(&t, &tt);
#endif

    if (date.tm_mday!=tt.tm_mday||date.tm_mon!=tt.tm_mon||date.tm_year!=tt.tm_year)
    {
        // 日期改变，重新计数
        log_count = 0;
        memcpy(&date, &tt, sizeof(date));
    }
    
    char dirpath[PATH_MAX];
    memset(dirpath, 0 ,PATH_MAX);
    if( fulldir[0] == 0)
        sprintf(dirpath, "%s%s",  LOG_PATH(), subdir);
    else
        sprintf(dirpath, "%s",  fulldir);

    snprintf(name, PATH_MAX, "%s%s%s_%d_%04d%02d%02d_%d.log",
       dirpath, "/", prefix, getpid(), date.tm_year+1900, date.tm_mon+1, date.tm_mday, log_count);
    
    
    char name2[PATH_MAX] = { 0 };
    snprintf(name2, PATH_MAX, "%s.tmp", name);
    
    logOpen(name2);
    log_count++;
    
    return os_log;
}

FILE *R5_Log::logOpen(const char *name)
{  
  // to do
  // 如果原文间存在则将原文件备份
    os_log = fopen(name, "w");
    
    static char szBuf[IO_BUFFER_SIZE+1];
    if (os_log == NULL)
    	fprintf(stderr, "ERROR init [%s] log file\n", name);
    else 
        setvbuf(os_log, szBuf, _IOFBF, IO_BUFFER_SIZE);

    return os_log;
}

void R5_Log::logOnlyClose()
{
    if(os_log != NULL)
        fclose(os_log);
    os_log = NULL;
}

void R5_Log::log1(const char *format, ...) //写日志
{
    if ( os_log == NULL )
    {
        if ( logOpen() == NULL )
        	return ;
    
    }
    else if ( log_size>=R5_LOG_SIZE )
    {
        if ( logOpen() == NULL )
        	return ;
        	
        log_size = 0;
    }
    //struct tm *lt;
    //struct timezone tz;
    struct timeval now;
    //gettimeofday(&now, &tz);
    gettimeofday(&now, 0);
    //lt = localtime(&now.tv_sec);
    char dateinfo[100];
    memset(dateinfo, 0 ,sizeof(dateinfo));
    
    struct tm ts;
    memset(&ts, 0, sizeof(ts));
    localtime_r(&now.tv_sec, &ts);
    sprintf(dateinfo, "%04d%02d%02d%02d%02d%02d", ts.tm_year+1900, ts.tm_mon + 1,ts.tm_mday,ts.tm_hour,ts.tm_min,ts.tm_sec);
    
    // 此处的日志格式可能需要调整。
    if (level>=r5_log_trace&&level<=r5_log_error)
    	fprintf(os_log, LOG_HEAD[level/20]);
    else
    	fprintf(os_log, "%06d", level);
    	
    //fprintf(os_log, "[%02d:%02d:%02d:%06d][%ld] ",
    //        lt->tm_hour, lt->tm_min, lt->tm_sec, now.tv_usec, getpid());
    fprintf(os_log, "[%s:%06d][%ld] ", dateinfo, now.tv_usec, getpid());
                        
    // Start of variable args section.
    va_list argp;
    va_start (argp, format);
  
    vfprintf(os_log, format, argp); 
      
    va_end (argp);
  
    log_size++;
}

void R5_Log::log2(const char *format, ...) //写终端
{
    //struct tm *lt;
    //struct timezone tz;
    struct timeval now;
    //gettimeofday(&now, &tz);
    gettimeofday(&now, 0);
    //lt = localtime(&now.tv_sec);
    
    // 此处的日志格式可能需要调整。
        if (level>=r5_log_trace&&level<=r5_log_error) fprintf(stdout, LOG_HEAD[level/20]);
        else     fprintf(stdout, "%06d", level);
        //fprintf(stdout, "[%02d:%02d:%02d:%06d][%ld] ",
    //        lt->tm_hour, lt->tm_min, lt->tm_sec, now.tv_usec, getpid());
    
    char dateinfo[100];
    memset(dateinfo, 0 ,sizeof(dateinfo));
    
    struct tm ts;
    memset(&ts, 0, sizeof(ts));
    localtime_r(&now.tv_sec, &ts);
    sprintf(dateinfo, "%04d%02d%02d%02d%02d%02d", ts.tm_year+1900, ts.tm_mon + 1,ts.tm_mday,ts.tm_hour,ts.tm_min,ts.tm_sec);
    
    fprintf(stdout, "[%s:%06d][%ld] ",
            dateinfo, now.tv_usec, getpid());
          
    // Start of variable args section.
    va_list argp;
    va_start (argp, format);
  
    vfprintf(stdout, format, argp); 
      
    va_end (argp);
}

bool R5_Log::checkPath(const char *path)
{
	DIR *dir = NULL;
    if((dir= opendir(path)) ==NULL) //子目录不存在，则创建
    {
        //if( mkdir(dirpath, 0755) == -1)
        //{
        	fprintf(stderr,"Directory: %s is not exist\n", path);
            return false;
        //}
    }
    if(dir != NULL)
    {
        if (closedir(dir) < 0)
        {
            fprintf(stderr,"Unable to close directory %s\n", path);
            return false;
        }
        dir = NULL;
    }
    if(access(path, W_OK) < 0)
    {
        fprintf(stderr,"The dir access permissions do not allow!path=%s\n", path);
        return false;
    }
	return true;
}

R5_Log g_r5_log;

// example:

//    R5_TRACE("main");

// 注意：下面的日志宏在使用的时候一定要多加一个小括号将日志信息括起来

//     R5_INFO(("INFO\n"));
//    R5_DEBUG(("DEBUG\n"));
//    R5_WARN(("WARN\n"));
//    R5_ERROR(("ERROR\n"));
    
//    struct tm *lt;
//    struct timezone tz;
//    struct timeval now;
//    gettimeofday(&now, &tz);
//    lt = localtime(&now.tv_sec);
//    R5_ERROR(("%02d:%02d:%02d:%06d\n",
//    lt->tm_hour, lt->tm_min, lt->tm_sec, now.tv_usec));

