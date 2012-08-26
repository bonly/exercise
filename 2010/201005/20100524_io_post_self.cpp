/**
 *  @file 20100524_io_post_self.cpp
 *
 *  @date 2012-7-11
 *  @author Bonly
 *  @brief 先post 到自己(相当于排队),再调目标操作
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
typedef boost::shared_ptr<boost::asio::io_service::work> work_ptr;
typedef boost::shared_ptr<boost::thread> thread_ptr;

class Worker
{
    public:
        Worker() :
                    io(new boost::asio::io_service), work(
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
        work_ptr work;
        thread_ptr thread;
};

Worker *g_wk;
class Sender: public boost::asio::io_service::service
{
    public:
        Sender() :  boost::asio::io_service::service(io),
                    //io(new boost::asio::io_service),
                    work(io)
        //:thread(new boost::thread(boost::bind(&Sender::run, this, wk)))
        {
            std::clog << "Sender create " << std::endl;
        }

        virtual void shutdown_service()
        {
        }
        void start(Worker &wk)
        {
            thread.reset(
                        new boost::thread(
                                    boost::bind(&boost::asio::io_service::run,
                                                boost::ref(io))));
            run_thread.reset(
                        new boost::thread(boost::bind(&Sender::run, boost::ref(this), wk)));
        }
        void run(Worker &wk)
        {
            unsigned int num = 0;
            while (!g_exit)
            {
                //boost::chrono::milliseconds dura(2000);
                //boost::this_thread::sleep_for(dura);
                //std::clog << "sender running " << std::endl;
                //printf("sender running\n");
                //wk.say(); ///可能是非线程安全的
                io.post(boost::bind(&Sender::say, this));
                //wk.io->post(boost::bind(&Worker::say, boost::ref(wk)));///线程安全
                //wk.io->post(boost::bind(&Worker::world, boost::ref(wk))); ///线程安全
                //wk.io->dispatch(boost::bind(&Worker::tellmetimes, boost::ref(wk), num)); ///注意是&wk,是引用,但内存暴涨,否则bind会复制一个wk实例
                //wk.io->post(boost::bind(&Worker::tellmetimes, boost::ref(wk), boost::ref(num))); ///注意是&wk,是引用,但内存暴涨,否则bind会复制一个wk实例

                 if (num >= (~std::size_t(0)))
                   num = 0;
                 ++num;
            }
        }
        void say()
        {
            //g_wk->io->post(boost::bind(&Worker::say, boost::ref(g_wk)));
            //g_wk->say();
            //std::clog << "say in sender" << std::endl;
            //printf("say in sender\n");
        }
        ~Sender()
        {
            //thread->join();
            //run_thread->join();
            std::clog << "Sender finish " << std::endl;
        }
    public:
        //io_ptr io;
        boost::asio::io_service io;
        //work_ptr work;
        boost::asio::io_service::work work;
        boost::shared_ptr<boost::thread> thread;
        boost::shared_ptr<boost::thread> run_thread;
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
    //work.start();
    Sender send;
    g_wk = &work;
    send.start(work);

    bool ret = SetConsoleCtrlHandler((PHANDLER_ROUTINE) signal_handle, true);
    if (ret == false)
    {
        std::clog << "Registery signal failed!" << std::endl;
        exit(-1);
    }
    while (!g_exit)
    {
        boost::chrono::milliseconds dura(2000);
        boost::this_thread::sleep_for(dura);
    }
    return 0;
}

/**
 * post/dispatch只是为了线程安全而提供的调用对象接口的方法!直接调暴露出来的方法可能是非线程安全的
 */

