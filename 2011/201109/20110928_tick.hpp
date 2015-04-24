/**
 * @file Tick.hpp
 * @brief 
 * @author bonly
 * @date 2013-3-1 bonly Created
 */

#ifndef TICK_HPP_
#define TICK_HPP_
#include <time.h>
#include <string.h>
namespace bus {

class Tick{
public:
    char Mins[60]; /// 0-59
    char Hrs[24];  /// 0-23
    char Days[32]; /// 1-31
    char Mons[12]; /// 0-11
    char Week[7];  /// 0-6 星天开始
    Tick(){Reset();}
    int Parse(char *ary);
    void Reset(){memset(this, 0, sizeof(Tick));}
    bool TestNow();
    bool Test(struct tm *tp);

private:
    char* parseField(char *ary, int modvalue, int off,
                            const char *const *names, char *ptr);
    void fixDayWeek();
};

} /* namespace bus */
#endif /* TICK_HPP_ */
