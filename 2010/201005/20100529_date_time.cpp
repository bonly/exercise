/**
 * @file 20100529_date_time.cpp
 * @brief
 *
 * @author bonly
 * @date 2012-7-14 bonly created
 */
/*
 * 定义一个文件包括源代码
 */
#define BOOST_DATE_TIME_SOURCE
#include <libs/date_time/src/gregorian/greg_names.hpp>
#include <libs/date_time/src/gregorian/date_generators.cpp>
#include <libs/date_time/src/gregorian/greg_month.cpp>
#include <libs/date_time/src/gregorian/greg_weekday.cpp>
#include <libs/date_time/src/gregorian/gregorian_types.cpp>

/*
 * 其它文件定义 BOOST_DATE_TIME_SOURCE
 * #define BOOST_DATE_TIME_NO_LIB OR BOOST_ALL_NO_LIB
 * 并include 头文件即可
 */
#include <boost/date_time/gregorian/gregorian.hpp>
#include <iostream>
using namespace std;
using namespace boost::gregorian;

int main()
{
    date d1(2012,7,14);
    date d2(2012, Jan, 1);

    date d3 = from_string("1999-12-31");
    date d4 ( from_string("2005/1/1"));
    date d5 = from_undelimited_string("20011219");

    clog << day_clock::local_day() << endl;
    clog << day_clock::universal_day() << endl;
    clog << d5 << endl;
    clog << d3 << endl;

    date d6(neg_infin);
    date d7(pos_infin);
    date d8(not_a_date_time);
    date d9(max_date_time);
    date d10(min_date_time);

    clog << d9 << endl;

    date::ymd_type ymd = d10.year_month_day();
    clog << ymd.year << endl;

    clog << d10.day_of_week() << endl;
    clog << d10.day_of_year() << endl;
    clog << d10.end_of_month() << endl;

    clog << to_simple_string(d4) << endl;
    clog << to_iso_extended_string(d3) << endl;
    clog << to_iso_string(d1) << endl;

    clog << d1 + days(-10) + months(1) + years(1) - weeks(1) << endl;

    date d(2101,3,30);
    d -= months(1);
    d -= months(1);
    d += months(2);
    ///如不想进位可使用 #undef BOOST_TIME_OPTIONAL_GREGORIAN_TYPES

    ///区间是左闭右开,左边必须小于右边
    date_period p1(date(2010,1,1), days(10));
    clog << p1.contains(date(2010,1,12)) << endl;

    for (day_iterator it=date(2010,3,1); it != date(2010,4,1); ++it)
    {
        clog << *it << "\t" << it->day_of_week() << endl;
    }

    clog << boost::gregorian::gregorian_calendar::is_leap_year(date(2012,12,1).year()) << endl;
    return 0;
}

