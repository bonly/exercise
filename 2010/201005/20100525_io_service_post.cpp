/**
 *  @file 20100522_io_service.cpp
 *
 *  @date 2012-7-10
 *  @author Bonly
 */
//#include <windows.h>
#include <boost/asio.hpp>
#include <boost/thread.hpp>
#include <boost/shared_ptr.hpp>
#include <boost/bind.hpp>
#include <iostream>
#include <boost/enable_shared_from_this.hpp>
#include <signal.h>
#include <boost/timer/timer.hpp>

const int TEST_COUNT = 2000000;
void some_work(int w)
{
    w += 2;
}

bool g_exit = false;
boost::asio::io_service g_io;
boost::asio::io_service::work g_work(g_io);

typedef boost::shared_ptr<boost::asio::io_service> io_ptr;
typedef boost::shared_ptr<boost::asio::io_service::strand> strand_ptr;
typedef boost::shared_ptr<boost::asio::io_service::work> work_ptr;
typedef boost::shared_ptr<boost::thread> thread_ptr;

boost::thread_group thr;

void fun()
{
    int i = 0;
    for (int j = 0; j < 100; ++j)
    {
        i += j;
        std::clog << i << std::endl;
    }
}
class Worker: public boost::enable_shared_from_this<Worker> ///此处没用到也是可以的
{
    public:
        Worker() :
                    // io(new boost::asio::io_service),
                    str(io), work(io)
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
            /*
             thread = boost::thread(
             boost::bind(&boost::asio::io_service::run, boost::ref(io)));
             */
            thr.add_thread(
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
            //std::clog << "hello" << std::endl;
        }
        void world()
        {
            std::clog << " world" << std::endl;
        }
        void stop()
        {
            io.stop();
            //thread.join();
        }
    public:
        boost::asio::io_service io;
        boost::asio::io_service::strand str;
        boost::asio::io_service::work work;
        //boost::thread thread;
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
            thr.add_thread(new boost::thread(boost::bind(&Sender::run, this)));
            //thread =boost::thread(boost::bind(&Sender::run, this));
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
                //g_wk->io.post(fun);
                g_io.post(fun);

                //g_wk->io.post(g_wk->str.wrap(boost::bind(&Worker::say, boost::ref(g_wk)))); ///线程安全
                //g_wk->io.post(boost::bind(&Worker::world, g_wk)); ///线程安全
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
        void stop()
        {
            //thread.join();
        }
        ~Sender()
        {
            std::clog << "Sender finish " << std::endl;
        }
    public:
        boost::thread thread;
};

#ifdef WIN32
bool signal_handle(DWORD dwCtrlType)
{
    std::clog << "Recive signal...\r\n" << std::endl;

    switch (dwCtrlType)
    {
        case CTRL_C_EVENT: // ctrl + c
        case CTRL_BREAK_EVENT:// ctrl + break
        case CTRL_CLOSE_EVENT:// 关闭控制台
        std::clog << "Prepare to exit service!" << std::endl;
        g_exit = true;
        return 0;
        default:
        return -1;
    }
    return -1;
}
#else

bool signal_handle()
{
    /// 阻塞所有信号,不传递给新创建的线程
    sigset_t new_mask;
    sigfillset(&new_mask);
    sigset_t old_mask;
    pthread_sigmask(SIG_BLOCK, &new_mask, &old_mask);

    /// 创建线程后台运行
    //boost::thread_group tg;
    //tg.create_thread(boost::bind(&server::run, &s));

    /// 恢复原来的信号
    pthread_sigmask(SIG_SETMASK, &old_mask, 0);

    /// 等待特定的信号
    sigset_t wait_mask;
    sigemptyset(&wait_mask);
    sigaddset(&wait_mask, SIGINT);
    sigaddset(&wait_mask, SIGQUIT);
    sigaddset(&wait_mask, SIGTERM);
    sigaddset(&wait_mask, SIGUSR2);
    sigaddset(&wait_mask, SIGHUP);
    pthread_sigmask(SIG_BLOCK, &wait_mask, 0);
    int sig = 0;
    int ret = -1;
    while (-1 != (ret = sigwait(&wait_mask, &sig)))
    {
        std::clog << "Receive signal. " << sig;
        if (sig == SIGUSR2)
        {
            //flush_log();
            continue;
        }
        if (sig == SIGHUP)
        {
            std::clog << "Receive reload config signal.";
            //load_conf(true);
            continue;
        }
        if (sig == SIGTERM || sig == SIGQUIT || sig == SIGINT)
        {
            std::clog << "Receive stop signal, Exit.";
            g_exit = true;
            break;
        }
    }
    if (ret == -1)
    {
        std::clog << "sigwaitinfo() returned err: " << errno << "\t"
                    << strerror(errno);
    }
    return 0;
}
#endif

void test_io_service_loop()
{
    boost::asio::io_service ios;
    std::unique_ptr<boost::asio::io_service::work> w(
                new boost::asio::io_service::work(ios));
    boost::thread t(boost::bind(&boost::asio::io_service::run, &ios));

    while (!g_exit)
    {
        printf("test_io_service\n");
        boost::timer::auto_cpu_timer auto_timer;

        for (int i = 0; i < TEST_COUNT; ++i)
        {
            ios.post(boost::bind(&some_work, i));
        }
    }

    w.reset();
    t.join();
}

void test_io_service()
{
    boost::asio::io_service ios;
    std::unique_ptr<boost::asio::io_service::work> w(
                new boost::asio::io_service::work(ios));
    boost::thread t(boost::bind(&boost::asio::io_service::run, &ios));

    printf("test_io_service\n");
    boost::timer::auto_cpu_timer auto_timer;

    for (int i = 0; i < TEST_COUNT; ++i)
    {
        ios.post(boost::bind(&some_work, i));
    }

    w.reset();
    t.join();
}

int main()
{
    //test_io_service(); 单次运行没问题,但内存好像也会涨
    //thr.add_thread(new boost::thread(boost::bind(&test_io_service))); 线程运行也没问题,内存也涨
    test_io_service_loop(); //明显是会涨的,证明queue没有限制大小

    /*
     //Worker work;
     //work.start();
     //g_wk = &work;

     Sender send;
     send.start();

     thr.add_thread(new boost::thread(boost::bind(&boost::asio::io_service::run, boost::ref(g_io))));
     */
#ifdef WIN32
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
#else
    signal_handle();
#endif

    //work.stop();
    // send.stop();
    thr.join_all();
    return 0;
}

/**
 * post/dispatch只是为了线程安全而提供的调用对象接口的方法(需strand来保证)!直接调暴露出来的方法可能是非线程安全的
 */

