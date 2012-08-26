/*
 * 20100615_callback.cpp
 *
 *  Created on: 2012-7-24
 *      Author: bonly
 */
#include <cstdio>
#include <boost/bind.hpp>
template<typename Handler>
class handle_object
{
    public:
        handle_object(Handler handler) :
                    Handler_(handler)
        {
        }
        Handler Handler_;
        static void handle_objectCall(void *hobject)
        {
            handle_object<Handler> *p = (handle_object<Handler> *) hobject;
            p->Handler_();
        }
};
typedef void (*handle_objectCallT)(void *hobject);
class Test
{
    public:
        Test()
        {
            m_Call = NULL;
        }
        template<class Handler>
        void Send(void *data, int len, Handler handler)
        {
            handle_object<Handler> *pobjet = new handle_object<Handler>(handler);
            m_Call = handle_object<Handler>::handle_objectCall;
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
    printf("%d\n", i);
}
void test2(int i, int j)
{
    printf("%d,%d\n", i, j);
}
int main(int argc, char* argv[])
{
    Test t;
    t.Send(NULL, 5, boost::bind(test, 2));
    t.Call();
    t.Send(NULL, 5, boost::bind(test2, 2, 3));
    t.Call();
    return 0;
}

