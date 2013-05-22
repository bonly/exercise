#include <iostream>
#include <string>
#include <time.h>
#include <cstdio>
using namespace std;

/*
 * 基本的时间类工具
 */
class CUtcTime
{
public:
  CUtcTime(){};
  ~CUtcTime(){};

  static time_t Time(time_t tTime=0);
  static time_t MakeTime(int nYear, int nMonth, int nDay, int nHour=0, int nMinute=0, int nSecond=0);
  static char *ToString(char *szTime, const char *szFormat, time_t tUtcTime);
  static char *Format(char *szTime, const char *szFormat, const struct tm *ptm);
  static int  IsLeapYear(int nYear);
};

#define SECONDS_1900_1970 2208988800L

//======================  CUtcTime Class  ========================
/*
 * 获取系统当前时间: 从00:00:00 on January 1, 1900开始的秒数
 * time() 从00:00:00 on January 1, 1970开始的时间
 *
 */
time_t CUtcTime::Time(time_t tTime)
{
    if (tTime <= 0)
    {
        time(&tTime);
    }

    return tTime + SECONDS_1900_1970;
}

/*
 * 闰年判断
 * 能被4整除但不能被100整除，可是却能被400整除的年是闰年
 * 闰年2月份是29日
 */
int CUtcTime::IsLeapYear(int nYear)
{
    if (((nYear%4 == 0) && (nYear%100 != 0)) || (nYear%400 == 0))
    {
        return 1;
    }
    else
    {
        return 0;
    }
}

/*
 * 日期格式输出
 */
char *CUtcTime::Format(char *szTime, const char *szFormat, const struct tm *ptm)
{
    char  chFmt;
    char *pBuf = szTime;

    while ((chFmt = *szFormat++) != '\0')
    {
        if (chFmt == '%')
        {
            switch (chFmt = *szFormat++)
            {
            case '%':
                *pBuf ++ = chFmt;
                break;
            case 'y':
                pBuf += sprintf(pBuf, "%2.2d", ptm->tm_year);
                break;
            case 'Y':
                pBuf += sprintf(pBuf, "%4.4d", ptm->tm_year + 1900);
                break;
            case 'm':
                pBuf += sprintf(pBuf, "%2.2d", ptm->tm_mon + 1);
                break;
            case 'd':
                pBuf += sprintf(pBuf, "%2.2d", ptm->tm_mday);
                break;
            case 'H':
                pBuf += sprintf(pBuf, "%2.2d", ptm->tm_hour);
                break;
            case 'M':
                pBuf += sprintf(pBuf, "%2.2d", ptm->tm_min);
                break;
            case 'S':
                pBuf += sprintf(pBuf, "%2.2d", ptm->tm_sec);
                break;
            default:
                perror("Time Format Error:");
                return NULL;
            }
        }
        else
        {
            *pBuf ++ = chFmt;
        }
    } // end of while

    *pBuf = '\0';

    return szTime;
}

/*
 *  tUtcTime 是指从00:00:00 on January 1, 1900开始的秒数
 */
char *CUtcTime::ToString(char *szTime, const char *szFormat, time_t tUtcTime)
{
    tUtcTime = tUtcTime - SECONDS_1900_1970;

  struct tm *ptm = localtime(&tUtcTime);

    return CUtcTime::Format(szTime, szFormat, ptm);
}

/*
 * 由本地时间构造一个time_t时间: seconds elapsed since 00:00:00 on January 1, 1900
 */
time_t CUtcTime::MakeTime(int nYear, int nMonth, int nDay, int nHour, int nMinute, int nSecond)
{
  struct tm tm;

  tm.tm_year  = nYear - 1900;
  tm.tm_mon   = nMonth - 1;
  tm.tm_mday  = nDay;
  tm.tm_hour  = nHour;
  tm.tm_min   = nMinute;
  tm.tm_sec   = nSecond;
  tm.tm_isdst = 0;
  tm.tm_wday  = 0;
  tm.tm_yday  = 0;

  return mktime(&tm) + SECONDS_1900_1970;
}

#include <boost/lexical_cast.hpp>
#include <string>
using namespace boost;
using namespace std;


unsigned long
today(string stm)
{
  unsigned long ret = 0;
  try
  {
  	time_t now;
  	time(&now);
  	struct tm* nw = localtime(&now);

  	int fs = stm.find_first_of (':');
  	int ls = stm.find_last_of (':');
  	string hour (stm,0,fs);
  	string min  (stm,fs+1,ls-fs-1);
  	string sec  (stm,ls+1);

  	ret = CUtcTime::MakeTime (nw->tm_year+1900,nw->tm_mon+1,nw->tm_mday,
  	                    lexical_cast<int>(hour),
  	                    lexical_cast<int>(min),
  	                    lexical_cast<int>(sec)
  	                    );
  	printf ("%ld\n",ret);
  }
  catch(std::exception &e)
  {
    cerr << e.what() << endl;
    ret = 0;
  }
	return ret;
}

int
main()
{
	time_t now;
	time(&now);

	struct tm* nw = gmtime(&now);
  printf ("year: %d\n",nw->tm_year+1900);
  printf ("month: %d\n",nw->tm_mon+1);
  printf ("mday: %d\n",nw->tm_mday);
  printf ("hour: %d\n",nw->tm_hour);
  printf ("min: %d\n",nw->tm_min);
  printf ("sec: %d\n",nw->tm_sec);

	nw = localtime(&now);
  printf ("year: %d\n",nw->tm_year+1900);
  printf ("month: %d\n",nw->tm_mon+1);
  printf ("mday: %d\n",nw->tm_mday);
  printf ("hour: %d\n",nw->tm_hour);
  printf ("min: %d\n",nw->tm_min);
  printf ("sec: %d\n",nw->tm_sec);

	string begin("12:30:00");
	string end("14:30:00");
	int fs = begin.find_first_of(':');
	int ls = begin.find_last_of(':');
	string hour (begin,0,fs-0);
	string min (begin,fs+1,ls-fs-1);
	string sec (begin,ls+1);
	printf("hour is: %s\n",hour.c_str());
	printf("min is: %s\n",min.c_str());
	printf("sec is: %s\n",sec.c_str());

	unsigned long td = today("12:30:00");
	printf("today is: %ld\n",td);
  cout << "today is: " << td << endl;
//  time_t b = CUtcTime::MakeTime(
//  		nw->tm_year+1900,nw->tm_mon,nw->tm_mday,
//
//  );




	return 0;
}


