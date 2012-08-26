/*
 * ���в���
 --report_level=detailed --show_progress=yes --run_test=my_suite1/my_test1
 */
#define BOOST_TEST_DYN_LINK
#define BOOST_TEST_MODULE ���� ����
#include <boost/test/included/unit_test.hpp>
#include <iostream>
#include <boost/tokenizer.hpp>
#include <string>
using namespace std;
using namespace boost;

BOOST_AUTO_TEST_SUITE( my_suite1 )   //�����1

BOOST_AUTO_TEST_CASE( my_test1 )    //���ڰ�1
{   //ʹ��Ĭ�ϵķָ���
    string s = "This is,  a test";
    tokenizer<> tok(s);
    tokenizer<>::iterator beg=tok.begin();
    BOOST_CHECK( strcmp((*beg).c_str(),"This") == 0 );
    ++beg;
    BOOST_CHECK( strcmp((*beg).c_str(),"is") == 0 );
    ++beg;
    BOOST_CHECK( strcmp((*beg).c_str(),"a") == 0 );
    ++beg;
    BOOST_CHECK( strcmp((*beg).c_str(),"test") == 0 );
}

BOOST_AUTO_TEST_CASE( my_test2 )   //���ڰ�1
{ //escaped_list_separator<>������,�ָ���CSV��ʽ
	string s = "Field 1,\"putting quotes around fields, allows commas\",Field 3";
  tokenizer<escaped_list_separator<char> > tok(s);
  tokenizer<escaped_list_separator<char> >::iterator beg=tok.begin();
  BOOST_CHECK( strcmp((*beg).c_str(),"Field 1") == 0 );
  ++beg;
  BOOST_CHECK( strcmp((*beg).c_str(),"putting quotes around fields, allows commas") == 0 );
}

BOOST_AUTO_TEST_CASE( my_test3 )   //���ڰ�1
{ //ƫ�Ʒֲ�
	string s = "12252001";
	int offsets[] = {2,2,4};
	offset_separator f(offsets, offsets+3);
  tokenizer<offset_separator> tok(s);

  tokenizer<offset_separator>::iterator beg=tok.begin();
  BOOST_CHECK( strcmp((*beg).c_str(),"12") == 0 );
  ++beg;
  BOOST_CHECK( strcmp((*beg).c_str(),"25") == 0 );
  ++beg;
  BOOST_CHECK( strcmp((*beg).c_str(),"2001") == 0 );
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

