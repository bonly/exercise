/*
 * 运行参数--report_level=detailed --show_progress=yes
 */
#define BOOST_TEST_MODULE 测试 案例
#include <boost/test/included/unit_test.hpp>


BOOST_AUTO_TEST_SUITE( my_suite1 )   //定义包1

BOOST_AUTO_TEST_CASE( my_test1 )    //属于包1
{
    BOOST_CHECK( 2 == 1 );
}

BOOST_AUTO_TEST_CASE( my_test2 )   //属于包1
{
    int i = 0;

    BOOST_CHECK_EQUAL( i, 2 );

    BOOST_CHECK_EQUAL( i, 0 );
}

BOOST_AUTO_TEST_SUITE_END()   //包1定义结束

//____________________________________________________________________________//
BOOST_AUTO_TEST_CASE( my_test3 )  //属于master包
{
    int i = 0;

    BOOST_CHECK_EQUAL( i, 0 );
}

//____________________________________________________________________________//

BOOST_AUTO_TEST_SUITE( my_suite2 )   //定义包2

BOOST_AUTO_TEST_CASE( my_test4 )  //属于包2
{
    int i = 0;

    BOOST_CHECK_EQUAL( i, 1 );
}

BOOST_AUTO_TEST_SUITE( internal_suite )  //定义包2的内部包3

BOOST_AUTO_TEST_CASE( my_test5 )  //属于包2里的包3
{
    int i = 0;
    BOOST_CHECK_EQUAL( i, 1 );
}

BOOST_AUTO_TEST_SUITE_END()  //包3定义结束


BOOST_AUTO_TEST_SUITE_END()  //包2定义结束

