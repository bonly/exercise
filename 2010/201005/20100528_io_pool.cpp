/**
 * @file 20100528_io_pool.cpp
 * @brief
 *
 * @author bonly
 * @date 2012-7-11 bonly created
 */
#include <boost/asio.hpp>
#include <boost/bind.hpp>
#include <boost/thread.hpp>
#include <iostream>

class thread_pool_checker: private boost::noncopyable
{
    public:

        thread_pool_checker(boost::asio::io_service& io_service,
                    boost::thread_group& threads, unsigned int max_threads,
                    long threshold_seconds, long periodic_seconds) :
                    io_service_(io_service), timer_(io_service), threads_(
                                threads), max_threads_(max_threads), threshold_seconds_(
                                threshold_seconds), periodic_seconds_(
                                periodic_seconds)
        {
            schedule_check();
        }

    private:

        void schedule_check();
        void on_check(const boost::system::error_code& error);

    private:

        boost::asio::io_service& io_service_;
        boost::asio::deadline_timer timer_;
        boost::thread_group& threads_;
        unsigned int max_threads_;
        long threshold_seconds_;
        long periodic_seconds_;
};

void thread_pool_checker::schedule_check()
{
    // Thread pool is already at max size.
    if (max_threads_ <= threads_.size())
    {
        std::cout << "Thread pool has reached its max.  Example will shutdown."
                    << std::endl;
        io_service_.stop();
        return;
    }

    // Schedule check to see if pool needs to increase.
    std::cout << "Will check if pool needs to increase in " << periodic_seconds_
                << " seconds." << std::endl;
    timer_.expires_from_now(boost::posix_time::seconds(periodic_seconds_));
    timer_.async_wait(
                boost::bind(&thread_pool_checker::on_check, this,
                            boost::asio::placeholders::error));
}

void thread_pool_checker::on_check(const boost::system::error_code& error)
{
    // On error, return early.
    if (error)
        return;

    // Check how long this job was waiting in the service queue.  This
    // returns the expiration time relative to now.  Thus, if it expired
    // 7 seconds ago, then the delta time is -7 seconds.
    boost::posix_time::time_duration delta = timer_.expires_from_now();
    long wait_in_seconds = -delta.seconds();

    // If the time delta is greater than the threshold, then the job
    // remained in the service queue for too long, so increase the
    // thread pool.
    std::cout << "Job job sat in queue for " << wait_in_seconds << " seconds."
                << std::endl;
    if (threshold_seconds_ < wait_in_seconds)
    {
        std::cout << "Increasing thread pool." << std::endl;
        threads_.create_thread(
                    boost::bind(&boost::asio::io_service::run, &io_service_));
    }

    // Otherwise, schedule another pool check.
    run();
}

// Busy work functions.
void busy_work(boost::asio::io_service&, unsigned int);

void add_busy_work(boost::asio::io_service& io_service, unsigned int count)
{
    io_service.post(boost::bind(busy_work, boost::ref(io_service), count));
}

void busy_work(boost::asio::io_service& io_service, unsigned int count)
{
    boost::this_thread::sleep(boost::posix_time::seconds(5));

    count += 1;

    // When the count is 3, spawn additional busy work.
    if (3 == count)
    {
        add_busy_work(io_service, 0);
    }
    add_busy_work(io_service, count);
}

int main()
{
    using boost::asio::ip::tcp;

    // Create io service.
    boost::asio::io_service io_service;

    // Add some busy work to the service.
    add_busy_work(io_service, 0);

    // Create thread group and thread_pool_checker.
    boost::thread_group threads;
    thread_pool_checker checker(io_service, threads, 3, // Max pool size.
                2, // Create thread if job waits for 2 sec.
                3); // Check if pool needs to grow every 3 sec.

    // Start running the io service.
    io_service.run();

    threads.join_all();

    return 0;
}

