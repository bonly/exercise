/*
 * ���в���--report_level=detailed --show_progress=yes
 */
#define BOOST_TEST_MODULE ���� ����
#include <boost/test/included/unit_test.hpp>


BOOST_AUTO_TEST_SUITE( my_suite1 )   //�����1

BOOST_AUTO_TEST_CASE( my_test1 )    //���ڰ�1
{
    BOOST_CHECK( 2 == 1 );
}

BOOST_AUTO_TEST_CASE( my_test2 )   //���ڰ�1
{
    int i = 0;

    BOOST_CHECK_EQUAL( i, 2 );

    BOOST_CHECK_EQUAL( i, 0 );
}

BOOST_AUTO_TEST_SUITE_END()   //��1�������

//____________________________________________________________________________//
BOOST_AUTO_TEST_CASE( my_test3 )  //����master��
{
    int i = 0;

    BOOST_CHECK_EQUAL( i, 0 );
}

//____________________________________________________________________________//

BOOST_AUTO_TEST_SUITE( my_suite2 )   //�����2

BOOST_AUTO_TEST_CASE( my_test4 )  //���ڰ�2
{
    int i = 0;

    BOOST_CHECK_EQUAL( i, 1 );
}

BOOST_AUTO_TEST_SUITE( internal_suite )  //�����2���ڲ���3

BOOST_AUTO_TEST_CASE( my_test5 )  //���ڰ�2��İ�3
{
    int i = 0;
    BOOST_CHECK_EQUAL( i, 1 );
}

BOOST_AUTO_TEST_SUITE_END()  //��3�������


BOOST_AUTO_TEST_SUITE_END()  //��2�������

