/**
 * @file 20100527_test_time.cpp
 * @brief
 *
 * @author bonly
 * @date 2012-7-11 bonly created
 */
#include <deque>
#include <boost/asio.hpp>
#include <boost/thread.hpp>
#include <boost/timer/timer.hpp>

const int TEST_COUNT = 2000000;

void some_work(int w)
{
    w += 2;
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

class deque_test
{
    public:
        deque_test()
        {
        }

        void run_test()
        {
            boost::thread t(boost::bind(&deque_test::work, this));

            printf("test_deque\n");
            boost::timer::auto_cpu_timer auto_timer;

            for (int i = 0; i < TEST_COUNT; ++i)
            {
                boost::lock_guard<boost::mutex> g(m_);
                q_.push_back(i);
            }

            t.join();
        }

        void work()
        {
            while (true)
            {
                boost::lock_guard<boost::mutex> g(m_);

                if (q_.empty())
                {
                    boost::this_thread::sleep(boost::posix_time::millisec(1));
                    continue;
                }

                int i = q_.front();
                some_work(i);
                q_.pop_front();

                if (i == TEST_COUNT - 1)
                {
                    break;
                }
            }
        }

    private:
        boost::mutex m_;
        std::deque<int> q_;
};

#ifdef WIN32
class iocp_test
{
    public:
        iocp_test()
        {
            iocp_ = ::CreateIoCompletionPort(INVALID_HANDLE_VALUE, NULL, NULL,
                        0);
        }

        ~iocp_test()
        {
            ::CloseHandle(iocp_);
        }

        void run_test()
        {
            boost::thread t(boost::bind(&iocp_test::work, this));

            printf("iocp_test\n");
            boost::timer::auto_cpu_timer auto_timer;

            for (int i = 0; i < TEST_COUNT; ++i)
            {
                ::PostQueuedCompletionStatus(iocp_, i, NULL, NULL);
            }

            t.join();
        }

        void work()
        {
            int counter = 0;
            DWORD value;
            ULONG completionKey;
            LPOVERLAPPED overlapped;

            while (true)
            {
                if (!::GetQueuedCompletionStatus(iocp_, (LPDWORD) & value,
                            &completionKey, &overlapped, INFINITE))
                {
                    printf("iocp get failed\n");
                    break;
                }

                counter++;
                if (counter == TEST_COUNT)
                {
                    break;
                }
            }
        }

    private:
        HANDLE iocp_;
};
#endif

int main(int argc, char** argv)
{
#ifdef WIN32
    iocp_test iocp;
    iocp.run_test();
#endif

    test_io_service();

    deque_test td;
    td.run_test();

    /* get this result

     test_io_service
     2.921010s wall, 1.484375s user + 3.046875s system = 4.531250s CPU (155.1%)
     test_deque
     0.343004s wall, 0.187500s user + 0.031250s system = 0.218750s CPU (63.8%)

     */

    /* when add iocp, get this result

     iocp_test
     2.525279s wall, 0.375000s user + 3.062500s system = 3.437500s CPU (136.1%)
     test_io_service
     3.098606s wall, 1.734375s user + 2.812500s system = 4.546875s CPU (146.7%)
     test_deque
     0.733904s wall, 0.171875s user + 0.015625s system = 0.187500s CPU (25.5%)

     */

    getchar();
}
