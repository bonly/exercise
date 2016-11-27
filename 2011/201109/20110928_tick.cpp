/**
 * @file Tick.cpp
 * @brief 
 * @author bonly
 * @date 2013-3-1 bonly Created
 */

#include "Tick.hpp"
#include <cstdlib>
#include <cstdio>

namespace bus {

const char *const WeekAry[] = {
    "sun",
    "mon",
    "tue",
    "wed",
    "thu",
    "fri",
    "sat",

    "Sun",
    "Mon",
    "Tue",
    "Wed",
    "Thu",
    "Fri",
    "Sat",
    NULL
};

const char *const MonAry[] = {
    "jan",
    "feb",
    "mar",
    "apr",
    "may",
    "jun",
    "jul",
    "aug",
    "sep",
    "oct",
    "nov",
    "dec",

    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Nov",
    "Dec",
    NULL
};

/*
数组, 数组大小, 修正, 别名,  被解释串
*/
char* Tick::parseField(char *ary, int modvalue, int off,
                        const char *const *names, char *ptr){
    char *base = ptr;
    int n1 = -1; //范围值的首
    int n2 = -1; //范围值的尾

    if (base == NULL) {
        return (NULL);
    }

    while (*ptr != ' ' && *ptr != '\t' && *ptr != '\n') {
        int skip = 0;  /// 初始化步长为0

        /// 处理数值或'*'
        if (*ptr == '*') {
            n1 = 0;     /// 此值为0时，所有数组成员值都设置为1
            n2 = modvalue - 1; /// 结束值为最大值-1
            skip = 1;
            ++ptr;
        } else if (*ptr >= '0' && *ptr <= '9') {  /// 数值类分析
            if (n1 < 0) {
                n1 = strtol(ptr, &ptr, 10) + off;  /// 起始点+修正值
            } else {
                n2 = strtol(ptr, &ptr, 10) + off;  /// 结束点+修正值
            }
            skip = 1;  /// 获取到值，设置步长
        } else if (names) { /// 别名数存在
            int i;

            for (i = 0; names[i]; ++i) {
                if (strncmp(ptr, names[i], strlen(names[i])) == 0) {
                    break;  /// 匹配别名
                }
            }
            if (names[i]) {
                ptr += strlen(names[i]);  /// 跳过字符串中的别名
                if (n1 < 0) {
                    n1 = i;  /// 根据别名设置开始点
                } else {
                    n2 = i;  /// 根据别名设置结束点 '-'操作
                }
                skip = 1;  /// 设置步长
            }
        }

        /// 处理范围 '-'
        if (skip == 0) { ///检查上一步是否成功，失败则结束
            printf("%s: %s\n","failed parsing ", base);
            return (NULL);
        }
        if (*ptr == '-' && n2 < 0) { ///范围性的，读取下一个字符作为结束点
            ++ptr;
            continue;
        }

        /// 取值结束，处理数组数据
        if (n2 < 0) { /// n2仍为原始值（-1），即起始结束点相同
            n2 = n1;
        }
        if (*ptr == '/') { /// “每”的处理方式
            skip = strtol(ptr + 1, &ptr, 10); /// 跳过n个数组点
        }

        {
            int s0 = 1;
            int failsafe = 1024;  /// 循环保险,避免无尽的死循环

            --n1;
            do {
                n1 = (n1 + 1) % modvalue; /// 使用下标自增

                if (--s0 == 0) { /// 用作'/'每操作时，跳过次数点 @note ==比较 级别 高于 运算--
                    ary[n1 % modvalue] = 1;  /// 给数组赋值 1：包括
                    s0 = skip;  /// 重置'/'操作的跳过点
                }
            }while (n1 != n2 && --failsafe); /// 起始结束值重合

            if (failsafe == 0) { /// 不应用尽保险循环次数
                printf("%s: %s\n" ,"failed parsing ", base);
                return (NULL);
            }
        }
        if (*ptr != ',') {  /// 不是','分隔多值，则结束
            break;
        }
        ++ptr;
        n1 = -1;
        n2 = -1;
    }

    if (*ptr != ' ' && *ptr != '\t' && *ptr != '\n') { /// 下一字符必须为间隔符
        printf("%s: %s\n", "failed parsing ", base);
        return (NULL);
    }

    while (*ptr == ' ' || *ptr == '\t' || *ptr == '\n') { /// 跳过间隔符
        ++ptr;
    }
#ifdef FEATURE_DEBUG_OPT
    if (true) {
        int i;

        for (i = 0; i < modvalue; ++i) {
            printf("\005%d", ary[i]);
        }
        printf("\005\n");
    }
#endif

    return (ptr);
}

bool Tick::Test(struct tm *tp){
    if (Mins[tp->tm_min] && Hrs[tp->tm_hour] &&
            (Days[tp->tm_mday] || Week[tp->tm_wday]) &&
            Mons[tp->tm_mon]){
        return true;
    }
    else return false;
}

bool Tick::TestNow(){
    time_t t = time(NULL);
    struct tm *tp = localtime(&t);
    return Test(tp);
}

int Tick::Parse(char *ary){
    char str[strlen(ary)+3];
    memcpy(str, ary, sizeof(str));
    strcat(str, "\te");

    char *ptr = str;
    ptr = parseField(Mins, 60, 0, NULL, str);
    ptr = parseField(Hrs, 24, 0, NULL, ptr);
    ptr = parseField(Days, 32, 0, NULL, ptr);
    ptr = parseField(Mons, 12, -1, MonAry, ptr);
    ptr = parseField(Week, 7, 0, WeekAry, ptr);

    if (ptr == NULL) {
        return -1;
    }

    /**
    * @note 修正日期和周的使用：
    * 如果其中一个不是*，另一个是*
    * 则*的那一个被所有置0
    */
    fixDayWeek();
    return 0;
}

void Tick::fixDayWeek(){
#define arysize(ary)    (sizeof(ary)/sizeof((ary)[0]))
    short i;
    short weekUsed = 0;
    short daysUsed = 0;

    for (i = 0; i < (short)arysize(Week); ++i) {
        if (Week[i] == 0) {
            weekUsed = 1;
            break;
        }
    }
    for (i = 0; i < (short)arysize(Days); ++i) {
        if (Days[i] == 0) {
            daysUsed = 1;
            break;
        }
    }
    if (weekUsed && !daysUsed) {
        memset(Days, 0, sizeof(Days));
    }
    if (daysUsed && !weekUsed) {
        memset(Week, 0, sizeof(Week));
    }
#undef arysize
}

} /* namespace bus */

/* test case
int main(){
    bus::Tick tk;
    char test[]="* 12 2,4 3-10/2 * abc";  /// 没有3,10/2的格式
    if (0!=tk.Parse(test))
        printf("parse failed 0\n");

    char test1[]="6 * * * *";
    tk.Reset();
    if (0!=tk.Parse(test1))
        printf("parse failed 1\n");

    if (tk.TestNow())
        printf("need to run");
    return 0;
}
*/
// g++ -static -c 20110928_tick.cpp -o libtick.a
// CREATE FUNCTION cron RETURNS STRING SONAME 'libbonly.so';
// select cron();
