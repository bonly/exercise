#include <string>
#include <iostream>
#include <cstdio>
using namespace std;

template <typename Handler>
class handle_object
{
public:
    handle_object(Handler handler):Handler_(handler){}
    Handler Handler_;
    static void handle_objectCall(void *hobject)
    {
        handle_object<Handler> *p =(handle_object<Handler> *) hobject;
        p->Handler_();
    }
};

typedef void (*handle_objectCallT)(void *hobject);

class Test
{
public:
    Test()
    {
        m_Call =NULL;
    }
    template<class Handler>
    void Send(void *data,int len,Handler handler)
    {
        handle_object<Handler> *pobjet = new handle_object<Handler>(handler);
        m_Call= handle_object<Handler>::handle_objectCall;
        handlerobj = pobjet;
    }
    void Call()
    {
        m_Call(handlerobj);
    }
    void *handlerobj;
    handle_objectCallT m_Call;
};

void test(int i)
{
    printf("%d\n",i);
}
void test2(int i,int j)
{
    printf("%d,%d\n",i,j);
}


#define BOOST_TEST_MODULE native_socket_test
#include <boost/test/included/unit_test.hpp>

BOOST_AUTO_TEST_CASE (hostname1)
{
  Test t;
  t.Send(NULL,5,boost::bind(test,2));
  t.Call();
  t.Send(NULL,5,boost::bind(test2,2,3));
  t.Call();
  BOOST_CHECK(true);
}

/*
 * mingw
 */

