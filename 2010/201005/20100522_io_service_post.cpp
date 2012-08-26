/**
 *  @file 20100522_io_service.cpp
 *
 *  @date 2012-7-10
 *  @author Bonly
 */
#include <windows.h>
#include <boost/asio.hpp>
#include <boost/thread.hpp>
#include <boost/shared_ptr.hpp>
#include <boost/bind.hpp>
#include <iostream>
#include <boost/enable_shared_from_this.hpp>

bool g_exit = false;

typedef boost::shared_ptr<boost::asio::io_service> io_ptr;
typedef boost::shared_ptr<boost::asio::io_service::strand> strand_ptr;
typedef boost::shared_ptr<boost::asio::io_service::work> work_ptr;
typedef boost::shared_ptr<boost::thread> thread_ptr;

class Worker: public boost::enable_shared_from_this<Worker> ///此处没用到也是可以的
{
    public:
        Worker() :
                    io(new boost::asio::io_service), str(
                                new boost::asio::io_service::strand(*io)), work(
                                new boost::asio::io_service::work(*io))
        //,thread(new boost::thread(boost::bind(&boost::asio::io_service::run, io))) ///此处生成线程会导至不断复制worker类
        {
            std::clog << "Worker create " << std::endl;
        }
        ~Worker()
        {
            std::clog << "Worker finish " << std::endl;
        }

        void start()
        {
            thread.reset(
                        new boost::thread(
                                    boost::bind(&boost::asio::io_service::run,
                                                boost::ref(io))));
        }
        void tellmetimes(unsigned int &num)
        {
            /*
             static unsigned int lnum = 10;
             //std::clog << "num: " << num  << std::endl; //clog在多线程中有内存泄漏!
             //std::clog << "lnum: " << lnum << std::endl;
             //printf("num: %d\n", num);
             //printf("lnum: %d\n", lnum);
             //if (lnum >= (~(int(0)))) //老是0?
             if (lnum >= (~(size_t(0))))
             lnum = 0;
             ++lnum;
             */

        }

        void say()
        {
            std::clog << "hello" << std::endl;
        }
        void world()
        {
            std::clog << " world" << std::endl;
        }
    public:
        io_ptr io;
        strand_ptr str;
        work_ptr work;
        thread_ptr thread;
};

Worker *g_wk;
class Sender: public boost::enable_shared_from_this<Sender>
{
    public:
        Sender()
        //:thread(new boost::thread(boost::bind(&Sender::run, this, wk)))
        {
            std::clog << "Sender create " << std::endl;
        }

        void start(/*Worker &wk*/)
        {
            thread.reset(
                        new boost::thread(
                                    boost::bind(&Sender::run, this)));
                                    //,boost::ref(wk)))); ///指定用引用方式,否则会多出一次释构
        }
        void run(/*Worker *wk*/)
        {
            unsigned int num = 0;
            boost::chrono::milliseconds dura(2000);
            while (!g_exit)
            {
                //boost::this_thread::sleep_for(dura);
                //std::clog << "sender running " << std::endl;
                //wk.say(); ///可能是非线程安全的
                g_wk->io->post(boost::bind(&Worker::say, boost::ref(g_wk)));///线程安全
                //wk.io->post(boost::bind(&Worker::world, boost::ref(wk))); ///线程安全
                //wk.io->dispatch(boost::bind(&Worker::tellmetimes, boost::ref(wk), num)); ///注意是&wk,是引用,但内存暴涨,否则bind会复制一个wk实例
                //wk.io->post(boost::bind(&Worker::tellmetimes, boost::ref(wk), boost::ref(num))); ///注意是&wk,是引用,但内存暴涨,否则bind会复制一个wk实例
                //*
                 if (num >= (~std::size_t(0)))
                 num = 0;
                 ++num;
                 //*/
            }
        }
        ~Sender()
        {
            std::clog << "Sender finish " << std::endl;
        }
    public:
        boost::shared_ptr<boost::thread> thread;
};
bool signal_handle(DWORD dwCtrlType)
{
    std::clog << "Recive signal...\r\n" << std::endl;

    switch (dwCtrlType)
    {
        case CTRL_C_EVENT: // ctrl + c
        case CTRL_BREAK_EVENT: // ctrl + break
        case CTRL_CLOSE_EVENT: // 关闭控制台
            std::clog << "Prepare to exit service!" << std::endl;
            g_exit = true;
            return 0;
        default:
            return -1;
    }
    return -1;
}

int main()
{
    Worker work;
    work.start();
    g_wk = &work;

    Sender send;
    send.start();

    bool ret = SetConsoleCtrlHandler((PHANDLER_ROUTINE) signal_handle, true);
    if (ret == false)
    {
        std::clog << "Registery signal failed!" << std::endl;
        exit(-1);
    }
    boost::chrono::milliseconds dura(2000);
    while (!g_exit)
    {
        boost::this_thread::sleep_for(dura);
    }
    return 0;
}

/**
 * post/dispatch只是为了线程安全而提供的调用对象接口的方法(需strand来保证)!直接调暴露出来的方法可能是非线程安全的
 */

